// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameStorage = "storage"

// Storage mapped from table <storage>
type Storage struct {
	StorageID int32         `gorm:"column:storage_id;primaryKey" json:"storage_id"`
	DriverID  int32         `gorm:"column:driver_id" json:"driver_id"`
	MemoryID  int32         `gorm:"column:memory_id" json:"memory_id"`
	Driver    StorageDriver `gorm:"foreignKey:StorageDriverID" json:"driver"`
	Memory    Memory        `gorm:"foreignKey:MemoryUnitID" json:"memory"`
}

// TableName Storage's table name
func (*Storage) TableName() string {
	return TableNameStorage
}
