package models

import (
	"airun/brain-tracker/server/src/types"

	"gorm.io/gorm"
)

type SignalData struct {
	gorm.Model
	Values []interface{}    `gorm:"column:values;serializer:json"`
	Sid    string           `gorm:"column:sid"`
	Time   float64          `gorm:"column:time"`
	Type   types.SignalType `gorm:"column:type"`
}
