// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameScreenResolution = "screen_resolution"

// ScreenResolution mapped from table <screen_resolution>
type ScreenResolution struct {
	ScreenResolutionID int32 `gorm:"column:screen_resolution_id;primaryKey" json:"screen_resolution_id"`
	Width              int32 `gorm:"column:width" json:"width"`
	Height             int32 `gorm:"column:height" json:"height"`
}

// TableName ScreenResolution's table name
func (*ScreenResolution) TableName() string {
	return TableNameScreenResolution
}
