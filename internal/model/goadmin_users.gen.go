// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGoadminUser = "goadmin_users"

// GoadminUser mapped from table <goadmin_users>
type GoadminUser struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Username      string    `gorm:"column:username;not null" json:"username"`
	Password      string    `gorm:"column:password;not null" json:"password"`
	Name          string    `gorm:"column:name;not null" json:"name"`
	Avatar        string    `gorm:"column:avatar" json:"avatar"`
	RememberToken string    `gorm:"column:remember_token" json:"remember_token"`
	CreatedAt     time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName GoadminUser's table name
func (*GoadminUser) TableName() string {
	return TableNameGoadminUser
}
