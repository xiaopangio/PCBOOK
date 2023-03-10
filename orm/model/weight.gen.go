// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWeight = "weight"

// Weight mapped from table <weight>
type Weight struct {
	WeightID     int32      `gorm:"column:weight_id;primaryKey;autoIncrement:true" json:"weight_id"`
	Value        float32    `gorm:"column:value" json:"value"`
	WeightUnitID int32      `gorm:"column:weight_unit_id" json:"weight_unit_id"`
	Unit         WeightUnit `gorm:"foreignKey:WeightUnitID" json:"unit"`
}

// TableName Weight's table name
func (*Weight) TableName() string {
	return TableNameWeight
}
