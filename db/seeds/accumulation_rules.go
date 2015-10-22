package seeds

import (
	"log"

	"github.com/FoxComm/FoxComm/endpoints"
	"github.com/FoxComm/FoxComm/repositories"
	"github.com/FoxComm/FoxComm/social_analytics/balance/accumulation"
	"github.com/FoxComm/libs/db/fixtures"
)

func (sr SeedRunner) GenerateAccumulationRules() {
	storeID := StoreID()

	ruleSets := map[string]accumulation.RuleSet{}
	err := fixtures.Load("accumulation_rule_sets", &ruleSets)
	if err == nil {
		log.Println("Generating accumulation rule sets")
		balanceAccum := repositories.BalanceAccumulationRuleSetRepo{}
		balanceAccum.InitializeWithStoreID(storeID, repositories.BalanceAccumulationRuleSetCollection, endpoints.SocialAnalyticsAPI)
		balanceAccum.DestroyAll()
		for _, ruleSet := range ruleSets {
			balanceAccum.Create(&ruleSet)
		}
		log.Println("Completed accumulation rule sets")
	}
}
