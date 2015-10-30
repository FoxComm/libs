package utils

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/FoxComm/libs/Godeps/_workspace/src/github.com/FoxComm/goauth2/oauth"
	"github.com/FoxComm/libs/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/FoxComm/libs/Godeps/_workspace/src/github.com/jinzhu/gorm"
	"github.com/FoxComm/libs/Godeps/_workspace/src/github.com/tuvistavie/securerandom"
	"github.com/FoxComm/libs/Godeps/_workspace/src/gopkg.in/mgo.v2"
)

var dataSourceDatabaseMap = map[string]*gorm.DB{}
var dataSourceMgoMap = map[string]*mgo.Database{}

func GetOriginUserID(c *gin.Context) int {
	idstr := c.Request.Header.Get("FC-User-ID")
	id, _ := strconv.Atoi(idstr)
	return id
}

func StoreID(c *gin.Context) int {
	idstr := c.Request.Header.Get("FC-Store-ID")
	id, _ := strconv.Atoi(idstr)
	return id
}

func GetSolrHost(c *gin.Context) string {
	return c.Request.Header.Get("FC-Solr-Host")
}

func GetSpreeToken(c *gin.Context) string {
	return c.Request.Header.Get("FC-User-Spree-Token")
}

func GetStoreAdminSpreeToken(c *gin.Context) string {
	return c.Request.Header.Get("FC-Store-Admin-Spree-Token")
}

func GetDataSource(c *gin.Context) string {
	return c.Request.Header.Get("FC-Data-Source")
}

func GetStoreHost(c *gin.Context) string {
	return c.Request.Header.Get("FC-Store-Host")
}

func GetPostgresWithDataSource(dataSource string) (*gorm.DB, error) {
	if db, ok := dataSourceDatabaseMap[dataSource]; ok {
		return db, nil
	} else {
		db, err := gorm.Open("postgres", dataSource)
		if err == nil {
			dataSourceDatabaseMap[dataSource] = &db
		}
		return &db, err
	}
}

func GetPostgres(c *gin.Context) (*gorm.DB, error) {
	return GetPostgresWithDataSource(GetDataSource(c))
}

func GetMongoWithDataSource(dataSource string) (*mgo.Database, error) {

	if db, ok := dataSourceMgoMap[dataSource]; ok {
		return db, nil
	}

	sources := strings.Split(dataSource, "#")

	if len(sources) < 2 {
		return nil, fmt.Errorf("Invalid data source: %+v", sources)
	}

	session, err := mgo.Dial(sources[0])

	if err != nil {
		return nil, err
	}

	dataSourceMgoMap[dataSource] = session.DB(sources[1])

	return dataSourceMgoMap[dataSource], nil
}

func GetMongo(c *gin.Context) (*mgo.Database, error) {
	return GetMongoWithDataSource(GetDataSource(c))
}

func GetHttpSslFlexibleClient() *http.Client {
	client := new(http.Client)
	client.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true, RootCAs: oauth.Pool},
	}
	return client
}

func GenerateToken(length int) string {
	// The result may contain A-Z, a-z, 0-9, “-”, "_"
	//  and “_”. “=” is also used if padding is true.
	token, _ := securerandom.UrlSafeBase64(length, false)
	return token
}

func ReferrerUrlPattern() *regexp.Regexp {
	return regexp.MustCompile("/s/([0-9a-zA-Z_=-])+/[0-9]$")
}
