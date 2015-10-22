package seeds

import (
	"log"
	"os"
	"strconv"

	"github.com/FoxComm/FoxComm/endpoints"
	"github.com/FoxComm/FoxComm/social_shopping/models"
)

func (sr SeedRunner) GenerateSocialShoppingPreferences() {
	ssPreferences := []models.Preference{
		{PreferenceableType: "SocialShopping", Key: "Network", Value: "Facebook", Enable: true},
		{PreferenceableType: "SocialShopping", Key: "Network", Value: "Twitter", Enable: true},
		{PreferenceableType: "Facebook", Key: "FB_APP_ID", Value: "898499940178813", Enable: false},
		{PreferenceableType: "Facebook", Key: "FB_SECRET", Value: "0328d75ede9f6f47598c19ae883a7948", Enable: false},
		{PreferenceableType: "Facebook", Key: "URL_REDIRECT_1", Value: "/account/general-settings", Enable: false},
		{PreferenceableType: "Facebook", Key: "URL_REDIRECT_2", Value: "/account/address-book", Enable: false},
	}

	ssFbPermissions := []models.FBPermission{
		{Name: "user_interests"}, {Name: "email"}, {Name: "user_friends"}, {Name: "publish_actions"},
	}
	fbPermissions := []models.FBPermission{}

	if storeIDStr := os.Getenv("StoreID"); storeIDStr != "" {
		storeID, _ := strconv.Atoi(storeIDStr)

		log.Println("Deleting SocialShopping Preferences...")
		ss := &models.SocialShopping{}
		ss.InitializeWithStoreID(storeID, (*endpoints.Endpoint)(endpoints.SocialShoppingAPI))
		ss.Delete(models.Preference{})
		log.Println("SocialShopping Preferences Deleted...")

		log.Println("Inserting SocialShopping Preferences...")
		for _, preference := range ssPreferences {
			ss.Save(&preference)
		}
		log.Println("SocialShopping Preferences Inserted...")

		log.Println("Verifying if FB Permissions exist already...")
		ss.Find(&fbPermissions)
		if len(fbPermissions) == 0 {
			log.Println("Inserting FB Permissions...")
			for _, fbPermission := range ssFbPermissions {
				ss.Save(&fbPermission)
			}
			log.Println("FB Permissions Inserted...")
		}
	} else {
		panic("Forgot to set env['StoreID']?")
	}
}
