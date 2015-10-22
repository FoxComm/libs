package seeds

import (
	"log"
	"os"

	"gopkg.in/mgo.v2/bson"

	"github.com/FoxComm/libs/db/migrations"
)

type SeedRunner struct {
	SeedsRun bool
}

var ClickAttribRuleId bson.ObjectId
var LinkToParentAttribRuleId bson.ObjectId

func (sr *SeedRunner) AttachToCommandLine() {
	// If you didn't run main.go with an optional 'seeds' param, let's just skip.
	if len(os.Args) < 2 || os.Args[1] != "seeds" && os.Args[1] != "migrations" {
		return
	}

	log.Printf("The SeedRunner has been attached to the commandLine %+v", os.Args)

	if len(os.Args) == 3 {
		if os.Args[1] == "seeds" {
			//You added the seeds command, let's run them.
			switch os.Args[2] {
			case "all":
				sr.GenerateAllRules()
			case "attribution":
				sr.GenerateAttributionRules()
			case "accumulation":
				sr.GenerateAccumulationRules()
			case "reset":
				sr.DeleteAllRules()
			case "social_shopping":
				sr.GenerateSocialShoppingPreferences()
			case "stores":
				sr.GenerateStores()
			case "promotions":
				sr.GeneratePromotionPreferences()
			}
		}
		if os.Args[1] == "migrations" {
			migrations.RunAllMigrations()
		}
	}
}

func (sr *SeedRunner) GenerateAllRules() {
	log.Printf("Generating All Rules...")
	sr.GenerateAttributionRules()
	sr.GenerateAccumulationRules()
	sr.GeneratePromotionPreferences()
	sr.GenerateBalanceRules()
	sr.GenerateSocialShoppingPreferences()
}
