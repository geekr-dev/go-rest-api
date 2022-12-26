package model

import (
	"fmt"

	"github.com/geekr-dev/go-rest-api/pkg/auth"
	"github.com/geekr-dev/go-rest-api/pkg/constant"
	"github.com/go-playground/validator/v10"
)

// User represents a registered user.
type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (u *UserModel) TableName() string {
	return "users"
}

// Create 创建一个新的用户
func (u *UserModel) Create() error {
	return DB.Create(&u).Error
}

// Update 更新一个已存在的用户
func (u *UserModel) Update() error {
	return DB.Save(u).Error
}

// 校验用户密码是否正确
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// 对用户密码字段进行加密处理
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// 验证用户字段
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// 删除指定Id对应用户
func DeleteUser(id uint64) error {
	u := UserModel{}
	u.BaseModel.Id = id
	return DB.Delete(&u).Error
}

// 获取指定用户名对应用户信息
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Where("username = ?", username).First(&u)
	return u, d.Error
}

// 根据指定用户ID获取用户信息
func GetUserById(id int) (*UserModel, error) {
	u := &UserModel{}
	d := DB.First(&u, id)
	return u, d.Error
}

// 获取用户列表信息
func ListUsers(username string, offset, limit int) ([]*UserModel, int64, error) {
	if limit == 0 {
		limit = constant.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count int64

	whereCondition := fmt.Sprintf("username like '%%%s%%'", username)
	// 统计总量
	if err := DB.Model(&UserModel{}).Where(whereCondition).Count(&count).Error; err != nil {
		return users, count, err
	}

	// 获取分页数据
	if err := DB.Model(&UserModel{}).Where(whereCondition).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}
