// Package build  @Author xiaobaiio 2023/3/4 8:52:00
package main

import (
	"fmt"
	"github.com/xiaopangio/pcbook/orm"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func generate(g *gen.Generator) {
	//user
	role := g.GenerateModel("role")
	user := g.GenerateModel("user", gen.FieldRelate(field.HasOne, "Role", role,
		&field.RelateConfig{
			GORMTag: "foreignKey:RoleID",
		}))
	g.ApplyBasic(role, user)
	//memory
	memoryUnit := g.GenerateModel("memory_unit")
	memory := g.GenerateModel("memory", gen.FieldRelate(field.HasOne, "Unit", memoryUnit,
		&field.RelateConfig{
			GORMTag: "foreignKey:MemoryUnitID",
		}))
	g.ApplyBasic(memoryUnit, memory)
	//storage
	storageDriver := g.GenerateModel("storage_driver")
	storage := g.GenerateModel(
		"storage",
		gen.FieldRelate(field.HasOne,
			"Driver",
			storageDriver,
			&field.RelateConfig{
				GORMTag: "foreignKey:StorageDriverID",
			}),
		gen.FieldRelate(
			field.HasOne,
			"Memory",
			memory,
			&field.RelateConfig{
				GORMTag: "foreignKey:MemoryUnitID",
			}))
	g.ApplyBasic(storageDriver, storage)
	//keyboard
	keyboardLayout := g.GenerateModel("keyboard_layout")
	keyboard := g.GenerateModel("keyboard", gen.FieldRelate(field.HasOne, "Layout", keyboardLayout,
		&field.RelateConfig{
			GORMTag: "foreignKey:LayoutID",
		}))
	g.ApplyBasic(keyboardLayout, keyboard)
	//screen
	screenPanel := g.GenerateModel("screen_panel")
	screenResolution := g.GenerateModel("screen_resolution")
	screen := g.GenerateModel(
		"screen",
		gen.FieldRelate(field.HasOne,
			"Resolution",
			screenResolution,
			&field.RelateConfig{
				GORMTag: "foreignKey:ScreenResolutionID",
			}),
		gen.FieldRelate(
			field.HasOne,
			"Panel",
			screenPanel,
			&field.RelateConfig{
				GORMTag: "foreignKey:PanelID",
			}))
	g.ApplyBasic(screenPanel, screenResolution, screen)
	//weight
	weightUnit := g.GenerateModel("weight_unit")
	weight := g.GenerateModel("weight", gen.FieldRelate(field.HasOne, "Unit", weightUnit,
		&field.RelateConfig{
			GORMTag: "foreignKey:WeightUnitID",
		}))
	g.ApplyBasic(weightUnit, weight)
	//cpu
	cpu := g.GenerateModel("cpu")
	//gpu
	gpu := g.GenerateModel("gpu")
	g.ApplyBasic(cpu, gpu)
	//laptap_gpu
	laptapGpu := g.GenerateModel("laptap_gpu", gen.FieldRelate(field.HasOne, "Gpu", gpu,
		&field.RelateConfig{
			GORMTag: "foreignKey:GpuID",
		}))
	//laptap_storage
	laptapStorage := g.GenerateModel("laptap_storage", gen.FieldRelate(field.HasOne, "Storage", storage,
		&field.RelateConfig{
			GORMTag: "foreignKey:StorageID",
		}))
	//laptap
	laptap := g.GenerateModel(
		"laptap",
		gen.FieldRelate(field.HasOne,
			"CPU",
			cpu,
			&field.RelateConfig{
				GORMTag: "foreignKey:CPUID",
			}),
		gen.FieldRelate(field.HasOne,
			"RAM",
			memory,
			&field.RelateConfig{
				GORMTag: "foreignKey:MemoryID",
			}),
		gen.FieldRelate(field.HasOne,
			"Screen",
			screen,
			&field.RelateConfig{
				GORMTag: "foreignKey:ScreenID",
			}),
		gen.FieldRelate(field.HasOne,
			"Keyboard",
			keyboard,
			&field.RelateConfig{
				GORMTag: "foreignKey:KeyboardID",
			}),
		gen.FieldRelate(field.HasOne,
			"Weight",
			weight,
			&field.RelateConfig{
				GORMTag: "foreignKey:WeightID",
			}),
		gen.FieldRelate(field.HasMany,
			"GPUS",
			laptapGpu,
			&field.RelateConfig{
				GORMTag: "foreignKey:LaptapID",
			}),
		gen.FieldRelate(field.HasMany,
			"Storages",
			laptapStorage,
			&field.RelateConfig{
				GORMTag: "foreignKey:LaptapID",
			}),
	)
	g.ApplyBasic(laptapGpu, laptapStorage, laptap)
}
func main() {
	dsn := orm.Dsn()
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Errorf("cannot open db: %w", err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./orm/dal",
		ModelPkgPath: "./model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery,
	})
	g.UseDB(db)
	//g.ApplyBasic(g.GenerateAllTable()...)
	generate(g)
	g.Execute()
}
