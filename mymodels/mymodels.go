package mymodels

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

//
// gorm audit model which:
// - add user ID of CreatedBy and UpdatedBy in corresponding model structure
// - upon gorm create/update, save db context value of "userID" to CreatedBy and UpdatedBy
// - if context value of "userID" is not found, leave it as is
//

type UserIDType string

const UserIdKey UserIDType = "userID"

type GormModel struct {
	// json need to encode the Deleted when the frontend syn the delta changes and need to know which item was deleted
	Deleted   soft_delete.DeletedAt `gorm:"softDelete:flag;" json:"deleted"`
	CreatedAt time.Time             `json:"createdAt"`
	UpdatedAt time.Time             `json:"updatedAt"`
}

type AuditModel struct {
	CreatedBy uint32 `json:"createdBy"`
	UpdatedBy uint32 `json:"updatedBy"`
}

func (a *AuditModel) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	userIdAny := ctx.Value(UserIdKey)
	userId, ok := userIdAny.(uint32)
	if !ok {
		return nil
	}
	a.CreatedBy = userId
	a.UpdatedBy = userId
	return nil
}

func (a *AuditModel) BeforeUpdate(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	userIdAny := ctx.Value(UserIdKey)
	userId, ok := userIdAny.(uint32)
	if !ok {
		return nil
	}
	tx.Statement.SetColumn("updated_by", userId)
	a.UpdatedBy = userId
	return nil
}
