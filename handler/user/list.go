package user

import (
	. "github.com/geekr-dev/go-rest-api/handler"
	"github.com/geekr-dev/go-rest-api/pkg/errno"
	"github.com/geekr-dev/go-rest-api/service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUsers(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
