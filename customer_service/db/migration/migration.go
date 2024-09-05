package migration

import (
	"fmt"

	database "github.com/uber-demo/customer/db"
	models "github.com/uber-demo/customer/db/model"
	"gorm.io/gorm"
)

func Upgrade1() {
	database := database.GetDb()

	createTables(database)

}

func createTables(database *gorm.DB) {
	tables := []any{}

	tables = addNewTables(database, models.Customer{}, tables)
	tables = addNewTables(database, models.Trip{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		fmt.Println(err.Error(), " db error") // switch to zap logging?
	}
	fmt.Println("tables created") // switch to zap logging !
}

func addNewTables(database *gorm.DB, model any, tables []any) []any {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}