package user

import (
	. "github.com/geekr-dev/go-rest-api/handler"
	"github.com/geekr-dev/go-rest-api/model"

	"github.com/geekr-dev/go-rest-api/pkg/errno"
	"github.com/gin-gonic/gin"
)

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user [post]
func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := r.checkParam(); err != nil {
		SendResponse(c, err, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	resp := CreateResponse{
		Username: r.Username,
	}
	SendResponse(c, nil, resp)
}
