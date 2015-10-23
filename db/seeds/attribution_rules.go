package seeds

import (
	"log"

	"github.com/FoxComm/libs/endpoints"
	"github.com/FoxComm/FoxComm/social_analytics/repositories"
	"github.com/FoxComm/FoxComm/social_analytics/attributions"
	"github.com/FoxComm/libs/db/fixtures"
)

func (sr SeedRunner) DeleteAllRules() {
	log.Printf("Deleting All Rules...")
	storeID := StoreID()

	ruleSet := repositories.ActivityAttributionRuleSetRepo{}
	ruleSet.InitializeWithStoreID(storeID, repositories.ActivityAttributionRuleSetCollection, endpoints.SocialAnalyticsAPI)
	ruleSet.DestroyAll()

	attributeRule := repositories.AttributionRuleRepo{}
	attributeRule.InitializeWithStoreID(storeID, repositories.AttributionRuleCollection, endpoints.SocialAnalyticsAPI)
	attributeRule.DestroyAll()

	promotionPref := repositories.PromotionPreferencesRepo{}
	promotionPref.InitializeWithStoreID(storeID, repositories.PromotionPreferencesCollection, endpoints.SocialAnalyticsAPI)
	promotionPref.DestroyAll()

	balanceAccum := repositories.BalanceAccumulationRuleSetRepo{}
	balanceAccum.InitializeWithStoreID(storeID, repositories.BalanceAccumulationRuleSetCollection, endpoints.SocialAnalyticsAPI)
	balanceAccum.DestroyAll()

	balanceRules := repositories.BalanceRuleSetRepo{}
	balanceRules.InitializeWithStoreID(storeID, repositories.BalanceRuleSetCollection, endpoints.SocialAnalyticsAPI)
	balanceRules.DestroyAll()
}

func (sr SeedRunner) GenerateAttributionRules() {
	storeID := StoreID()

	ruleSets := map[string]attributions.RuleSet{}
	err := fixtures.Load("attribution_rule_sets", &ruleSets)
	if err == nil {
		ruleSetRepo := repositories.ActivityAttributionRuleSetRepo{}
		err := ruleSetRepo.InitializeWithStoreID(storeID, repositories.ActivityAttributionRuleSetCollection, endpoints.SocialAnalyticsAPI)
		if err != nil {
			log.Fatalf("Can't connect to ruleSetRepo: %s", err)
		}
		log.Println("Generating attribution rule sets")

		for _, ruleSet := range ruleSets {
			ruleSetRepo.Create(&ruleSet)
		}
		log.Println("Completed attribution rule sets")
	}
}
