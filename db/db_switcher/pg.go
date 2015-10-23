package db_switcher

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/jpfuentes2/go-env/autoload"

	"github.com/FoxComm/libs/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PG struct {
	*gorm.DB
}

func (pg *PG) InitializeWithContext(context *gin.Context) error {
	db, err := utils.GetPostgres(context)
	pg.DB = db
	return err
}

func (pg *PG) InitializeForFeature(featureName string) error {
	dataSourceKey := fmt.Sprintf("%s_DATASOURCE", strings.ToUpper(featureName))
	dataSource := os.Getenv(dataSourceKey)

	if dataSource == "" {
		return fmt.Errorf("Data source for %s is not found in the ENV", dataSourceKey)
	}

	db, err := utils.GetPostgresWithDataSource(dataSource)
	if err != nil {
		return err
	}

	pg.DB = db
	return nil
}
