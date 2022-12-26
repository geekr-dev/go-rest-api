package user

import (
	"strconv"

	. "github.com/geekr-dev/go-rest-api/handler"
	"github.com/geekr-dev/go-rest-api/pkg/errno"

	"github.com/geekr-dev/go-rest-api/model"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
