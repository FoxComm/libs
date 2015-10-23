package migrations

import (
	"fmt"
	"os"
	"strconv"

	"github.com/FoxComm/libs/endpoints"
	social_shopping_migration "github.com/FoxComm/FoxComm/social_shopping/migration"
	feature_manager_migration "github.com/FoxComm/core_services/feature_manager/migration"
	user_migration "github.com/FoxComm/core_services/user/migration"
	"github.com/FoxComm/core_services/user/service"
)

func RunAllMigrations() {
	feature_manager_migration.Run()

	storeIDStr := os.Getenv("StoreID")
	if storeIDStr == "" {
		fmt.Println("Forgot to set env['StoreID']?")
		os.Exit(0)
	}
	storeID, _ := strconv.Atoi(storeIDStr)

	ssMigration := social_shopping_migration.Migration{}
	if err := ssMigration.InitializeWithStoreID(storeID, (*endpoints.Endpoint)(endpoints.SocialShoppingAPI)); err == nil {
		ssMigration.Run()
	} else {
		fmt.Printf("Failed to initalize migration with StoreID: %v\n", storeID)
	}

	userMigration := user_migration.Migration{}
	if err := userMigration.InitializeWithStoreID(storeID, endpoints.UserAPI); err == nil {
		userMigration.Run()

		u := &user.User{}
		u.InitializeWithStoreID(storeID, endpoints.UserAPI)
		u.Where(user.User{Email: "admin@wearebeautykind.com"}).Assign("Role", "admin").FirstOrCreate(u)
		u.UpdatePassword("123qwe123")
	} else {
		fmt.Printf("Failed to initalize migration with StoreID: %v\n", storeID)
	}
}
