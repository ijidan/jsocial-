// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGoadminUserPermissions = "goadmin_user_permissions"

// GoadminUserPermissions mapped from table <goadmin_user_permissions>
type GoadminUserPermissions struct {
	UserID       int32     `gorm:"column:user_id;primaryKey" json:"user_id"`
	PermissionID int32     `gorm:"column:permission_id;primaryKey" json:"permission_id"`
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName GoadminUserPermissions's table name
func (*GoadminUserPermissions) TableName() string {
	return TableNameGoadminUserPermissions
}
