// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLaptap = "laptap"

// Laptap mapped from table <laptap>
type Laptap struct {
	LaptapID    string          `gorm:"column:laptap_id;primaryKey" json:"laptap_id"`
	Brand       string          `gorm:"column:brand" json:"brand"`
	Name        string          `gorm:"column:name" json:"name"`
	CPUID       int32           `gorm:"column:cpu_id" json:"cpu_id"`
	RAMID       int32           `gorm:"column:ram_id" json:"ram_id"`
	ScreenID    int32           `gorm:"column:screen_id" json:"screen_id"`
	KeyboardID  int32           `gorm:"column:keyboard_id" json:"keyboard_id"`
	WeightID    int32           `gorm:"column:weight_id" json:"weight_id"`
	PriceUsd    float32         `gorm:"column:price_usd" json:"price_usd"`
	ReleaseYear int32           `gorm:"column:release_year" json:"release_year"`
	UpdateAt    time.Time       `gorm:"column:update_at" json:"update_at"`
	CPU         CPU             `gorm:"foreignKey:CPUID" json:"cpu"`
	RAM         Memory          `gorm:"foreignKey:MemoryID" json:"ram"`
	Screen      Screen          `gorm:"foreignKey:ScreenID" json:"screen"`
	Keyboard    Keyboard        `gorm:"foreignKey:KeyboardID" json:"keyboard"`
	Weight      Weight          `gorm:"foreignKey:WeightID" json:"weight"`
	GPUS        []LaptapGpu     `gorm:"foreignKey:LaptapID" json:"gpus"`
	Storages    []LaptapStorage `gorm:"foreignKey:LaptapID" json:"storages"`
}

// TableName Laptap's table name
func (*Laptap) TableName() string {
	return TableNameLaptap
}
