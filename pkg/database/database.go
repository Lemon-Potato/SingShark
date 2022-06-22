package database

import (
	"database/sql"
	"fmt"
	"github.com/Lemon-Potato/SingShark/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBEngine struct {
	db    *gorm.DB
	sqldb *sql.DB
}

func NewDatabase(databaseSetting *setting.DatabaseSettingS) (*DBEngine, error) {
	var dbConfig gorm.Dialector
	var err error
	var DB *gorm.DB
	var SQLDB *sql.DB
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.DbHost,
		databaseSetting.DbPort,
		databaseSetting.DbDatabase,
		databaseSetting.Charset,
	)
	dbConfig = mysql.New(mysql.Config{
		DSN: dsn,
	})

	DB, err = gorm.Open(dbConfig)

	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}

	return &DBEngine{DB, SQLDB}, nil
}
