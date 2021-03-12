package postgres

import "gorm.io/gorm"

//go:generate go run $PWD/tools/codegen/codegen.go model-stubs Company

type Company struct {
	gorm.Model
	ID    string `gorm:"primaryKey" map:"id"`
	Title string `map:"title"`
}
