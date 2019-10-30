package sqlite

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"stockCollector/3-outbound/database/models"
)

func New() (Database, error) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		return Database{}, fmt.Errorf("Error connecting to Sqlite3 DB instance: %w", err)
	}

	return Database{
		ctx: db,
	}, err
}

type Database struct {
	ctx *gorm.DB
}

func (db Database) Migrate() error {
	db.ctx.SingularTable(true)
	db.ctx.AutoMigrate(&models.Company{}, &models.StockPriceSnapshot{})
	db.ctx.Model(&models.StockPriceSnapshot{}).AddForeignKey("CompanyId", "Company(Id)", "RESTRICT", "RESTRICT")
	return nil
}

func (db Database) Close() error {
	return db.ctx.Close()
}