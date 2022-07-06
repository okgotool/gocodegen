// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysRole = "sys_role"

// SysRole mapped from table <sys_role>
type SysRole struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RoleName       string    `gorm:"column:role_name" json:"role_name"`
	RoleNameEn     string    `gorm:"column:role_name_en" json:"role_name_en"`
	Status         int32     `gorm:"column:status" json:"status"`
	Priority       int32     `gorm:"column:priority" json:"priority"`
	Comment        string    `gorm:"column:comment" json:"comment"`
	Deleted        int32     `gorm:"column:deleted" json:"deleted"`
	LastmodifiedBy string    `gorm:"column:lastmodified_by" json:"lastmodified_by"`
	Lastmodified   time.Time `gorm:"column:lastmodified" json:"lastmodified"`
	CreatedBy      string    `gorm:"column:created_by" json:"created_by"`
	Created        time.Time `gorm:"column:created" json:"created"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedBy      string    `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName SysRole's table name
func (*SysRole) TableName() string {
	return TableNameSysRole
}
