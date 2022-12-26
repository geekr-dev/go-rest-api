package user

import (
	"fmt"

	"github.com/geekr-dev/go-rest-api/handler"
	"github.com/geekr-dev/go-rest-api/pkg/errno"
	"github.com/geekr-dev/go-rest-api/pkg/log"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	username := c.Param("username")
	log.Info("URL username: %s", username)

	desc := c.Query("desc")
	log.Info("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Info("Header Content-Type: %s", contentType)

	log.Debug("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		handler.SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}

	if r.Password == "" {
		handler.SendResponse(c, fmt.Errorf("password is empty"), nil)
	}

	resp := CreateResponse{
		Username: username,
	}
	handler.SendResponse(c, nil, resp)
}
