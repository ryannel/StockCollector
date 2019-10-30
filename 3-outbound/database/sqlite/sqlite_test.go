package sqlite

import (
	"github.com/stretchr/testify/assert"
	"stockCollector/3-outbound/database/models"
	"testing"
)

func Test_ConnectToDb_ShouldNotError(t *testing.T) {
	_, err := New()
	assert.Nil(t, err)
}

func Test_MigrateDb_ShouldNotError(t *testing.T) {
	db, err := New()
	assert.Nil(t, err)

	err = db.Migrate()
	assert.Nil(t, err)
}

func Test_InsertCustomer_ShouldNotError(t *testing.T) {
	db, err := New()
	assert.Nil(t, err)

	err = db.Migrate()
	assert.Nil(t, err)

	company := models.Company{
		CompanyName:  "TestCompany",
		Industry:     "Industry",
		Sector:       "Sector",
		Symbol:       "Symbol",
		Exchange:     "Exchange",
		Cusip:        "Cusip",
	}

	isNew := db.ctx.NewRecord(company)
	assert.True(t, isNew)

	err = db.ctx.Create(&company).Error
	assert.Nil(t, err)

	var companies []models.Company
	err = db.ctx.Find(&companies).Error
	assert.Nil(t, err)

	assert.NotEmpty(t, companies)
}