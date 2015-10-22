package db_switcher

import (
	"github.com/FoxComm/FoxComm/endpoints"
	"github.com/FoxComm/core_services/feature_manager/core"
	"github.com/FoxComm/libs/db/masterdb"
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

func (pg *PG) InitializeWithStoreID(storeID int, feature *endpoints.Endpoint) error {
	var storeFeature core.StoreFeature
	masterdb.Db().Joins("inner join features on features.id = store_features.feature_id").
		Where("store_id = ? AND features.name = ?", storeID, feature.Name).First(&storeFeature)
	db, err := utils.GetPostgresWithDataSource(storeFeature.Datasource)
	pg.DB = db
	return err
}
