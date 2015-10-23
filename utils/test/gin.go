package test

import (
	"net/http"

	"github.com/FoxComm/libs/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

// NewContext creates a dummy Gin Context object. This can be used to pass down
// the stack and initialize repositories.
func NewContext(dataSource string) (*gin.Context, error) {
	req, err := http.NewRequest("GET", "#", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("FC-Data-Source", dataSource)

	return &gin.Context{Request: req}, nil
}
