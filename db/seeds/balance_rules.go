package seeds

import (
	"log"

	"github.com/FoxComm/FoxComm/endpoints"
	"github.com/FoxComm/FoxComm/social_analytics/repositories"
	"github.com/FoxComm/FoxComm/social_analytics/balance"
	"github.com/FoxComm/libs/db/fixtures"
)

func (sr SeedRunner) GenerateBalanceRules() {
	storeID := StoreID()
	ruleSets := map[string]balance.RuleSet{}

	err := fixtures.Load("balance_rule_sets", &ruleSets)

	if err == nil {
		log.Println("Generating balance rule sets")

		for _, ruleSet := range ruleSets {
			balanceRuleSet := repositories.BalanceRuleSetRepo{}
			balanceRuleSet.InitializeWithStoreID(storeID, repositories.BalanceRuleSetCollection, endpoints.SocialAnalyticsAPI)
			err = balanceRuleSet.Create(&ruleSet)
		}

		log.Println("Completed balance rule sets")
	}
}
