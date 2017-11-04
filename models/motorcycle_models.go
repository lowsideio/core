package models

import (
  "time"

  _ "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

  "github.com/satori/go.uuid"
)

type MotorcycleModel struct {
  ID                    uuid.UUID   `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id" string:"id" form:"id" query:"id"`

  Brand                 string      `json:"brand,omitempty" string:"brand" form:"brand" query:"brand"`
  Model                 string      `json:"model,omitempty" string:"model" form:"model" query:"model"`
  Category              string      `json:"category,omitempty" string:"category" form:"category" query:"category"`

  Slug                   string      `json:"slug,omitempty" string:"category" form:"slug" query:"slug"`

  CreatedAt             time.Time   `json:"createdAt,omitempty" string:"createdAt" form:"createdAt" query:"createdAt"`
  UpdatedAt             time.Time   `json:"updatedAt,omitempty" string:"updatedAt" form:"updatedAt" query:"updatedAt"`
  DeletedAt             time.Time   `gorm:"default:null" json:"deletedAt" string:"deletedAt" form:"deletedAt" query:"deletedAt"`
}

func (m MotorcycleModel) TableName() string {
  return "motorcycles_models";
}
