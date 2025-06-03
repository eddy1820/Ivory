package database

import (
	"fmt"
	"gate/internal/config"
	"gate/internal/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewDBEngine(cfg config.DatabaseConfig, serverCfg config.ServerConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local",
		cfg.UserName, cfg.Password, cfg.Host, cfg.DBName, cfg.Charset, cfg.ParseTime)

	logger.Logger.Debug().Msg("dsn: " + dsn)

	logLevel := gormLogger.Silent
	if serverCfg.IsDebug() {
		logLevel = gormLogger.Info
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger:         gormLogger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		logger.Logger.Error().Err(err).Msg("failed to connect database")
		return nil, err
	}
	logger.Logger.Info().Msg("MySQL connected successfully")

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)

	return db, nil
}
