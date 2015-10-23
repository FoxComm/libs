package db_switcher

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/FoxComm/libs/endpoints"
	"github.com/FoxComm/core_services/feature_manager/core"
	"github.com/FoxComm/libs/db/masterdb"
	"github.com/FoxComm/libs/utils"
)

type Mongo struct {
	Collection *mgo.Collection
	Database   *mgo.Database
	Context    *gin.Context
}

func (mgo *Mongo) GetContext() *gin.Context {
	return mgo.Context
}

func (mgo *Mongo) InitializeWithContext(context *gin.Context, collection string) error {
	db, err := utils.GetMongo(context)

	if err == nil {
		mgo.Context = context
		mgo.Collection = db.C(collection)
		mgo.Database = db
	}

	return err
}

func (mgo *Mongo) InitializeWithRequest(request *http.Request, collection string) error {
	dataSource := request.Header.Get("FC-Data-Source")

	db, err := utils.GetMongoWithDataSource(dataSource)

	if err == nil {
		mgo.Collection = db.C(collection)
		mgo.Database = db
	}

	return err
}

func (mgo *Mongo) InitializeWithStoreID(storeID int, collection string, feature *endpoints.Endpoint) error {
	var storeFeature core.StoreFeature

	masterdb.Db().
		Joins("INNER JOIN features ON features.id = store_features.feature_id").
		Where("store_id = ? AND features.name = ?", storeID, feature.Name).
		First(&storeFeature)

	db, err := utils.GetMongoWithDataSource(storeFeature.Datasource)

	if err == nil {
		mgo.Collection = db.C(collection)
		mgo.Database = db
	}

	return err
}

func (repo *Mongo) Create(value interface{}) error {
	runCallbacks(value, "BeforeSave", "BeforeCreate")
	err := repo.Collection.Insert(value)
	if err == nil {
		runCallbacks(value, "AfterSave", "AfterCreate")
	}
	return err
}

func (repo *Mongo) CreateWithId(value interface{}) (string, error) {
	runCallbacks(value, "BeforeSave", "BeforeCreate")

	changes, err := repo.Collection.UpsertId(bson.NewObjectId(), value)

	if err == nil {
		runCallbacks(value, "AfterSave", "AfterCreate")
	}
	return changes.UpsertedId.(bson.ObjectId).Hex(), err
}

func (repo *Mongo) Find(query interface{}, result interface{}) error {
	err := repo.Collection.Find(query).One(result)
	if err == nil {
		runCallbacks(result, "AfterFind")
	}
	return err
}

func (repo *Mongo) Update(id string, value interface{}) error {
	runCallbacks(value, "BeforeSave", "BeforeUpdate")
	err := repo.Collection.UpdateId(bson.ObjectIdHex(id), value)
	if err == nil {
		runCallbacks(value, "AfterSave", "AfterUpdate")
	}
	return err
}

func (repo *Mongo) Upsert(query interface{}, result interface{}) error {
	_, err := repo.Collection.Upsert(query, result)
	return err
}

func (repo *Mongo) FindAll(query interface{}, result interface{}) error {
	return repo.Collection.Find(query).All(result)
}

func (repo *Mongo) FindPage(query interface{}, result interface{}, page, per int, sort ...string) error {
	skip := (per * page) - per
	err := repo.Collection.Find(query).Skip(skip).Limit(per).Sort(sort...).All(result)
	return err
}

func (repo *Mongo) FindAllSorted(query interface{}, result interface{}, sortBy string) error {
	err := repo.Collection.Find(query).Sort(sortBy).All(result)
	return err
}

func (repo *Mongo) FindAllMonthly(grouped interface{}, result interface{}) error {
	pipe := repo.Collection.Pipe(grouped)
	err := pipe.All(result)

	return err
}

func (repo *Mongo) FindById(id string, result interface{}) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	return repo.Find(query, result)
}

func (repo *Mongo) DestroyAll() error {
	err := repo.Collection.DropCollection()
	return err
}

func (repo *Mongo) DestroyById(id string) error {
	err := repo.Collection.RemoveId(bson.ObjectIdHex(id))
	return err
}

func runCallbacks(value interface{}, callbacks ...string) {
	for _, callback := range callbacks {
		reflectedValue := getValue(reflect.ValueOf(value)).Addr()
		if method := reflectedValue.MethodByName(callback); method.IsValid() {
			methodInterface := method.Interface()
			if m, ok := methodInterface.(func()); ok {
				m()
			}
		}
	}
}

func getValue(value reflect.Value) reflect.Value {
	if value.Kind() != reflect.Ptr {
		return value
	} else {
		value = value.Elem()
		return getValue(value)
	}
}
