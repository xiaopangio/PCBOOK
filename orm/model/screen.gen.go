// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameScreen = "screen"

// Screen mapped from table <screen>
type Screen struct {
	ScreenID     int32            `gorm:"column:screen_id;primaryKey;autoIncrement:true" json:"screen_id"`
	SizeInch     float32          `gorm:"column:size_inch" json:"size_inch"`
	ResolutionID int32            `gorm:"column:resolution_id" json:"resolution_id"`
	PanelID      int32            `gorm:"column:panel_id" json:"panel_id"`
	Multitouch   int32            `gorm:"column:multitouch" json:"multitouch"`
	Resolution   ScreenResolution `gorm:"foreignKey:ScreenResolutionID" json:"resolution"`
	Panel        ScreenPanel      `gorm:"foreignKey:PanelID" json:"panel"`
}

// TableName Screen's table name
func (*Screen) TableName() string {
	return TableNameScreen
}
