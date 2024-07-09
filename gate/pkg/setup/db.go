package setup

import (
	"fmt"
	"gate/global"
	"gate/pkg/setting"
	"gorm.io/driver/mysql"
	"time"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gorm.io/gorm"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	//root:123456@tcp(IvoryDb)/IvoryDb?charset=utf8&parseTime=true&loc=Local
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	fmt.Println("!!!!!" + dsn)

	logLevel := logger.Silent
	if global.ServerSetting.IsDebug() {
		logLevel = logger.Info
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		// default size for string fields
		DefaultStringSize: 255,
		// disable datetime precision, which not supported before MySQL 5.6
		DisableDatetimePrecision: true,
		// drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameIndex: true,
		// `change` when rename column, rename column not supported before MySQL 8, MariaDB
		DontSupportRenameColumn: true,
		// autoconfigure based on currently MySQL version
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger:         logger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Printf("!!!!!db failed: %s", err)
		return nil, err
	}

	fmt.Printf("!!!!!!!db successful: %s", err)

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConn)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConn)

	//err = db.Callback().Create().Before("gorm:create").Register("i18n_app:before_create", updateTimestampForCreateCallback)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = db.Callback().Update().Before("gorm:update").Register("i18n_app:before_update", updateTimestampForUpdateCallback)
	//if err != nil {
	//	return nil, err
	//}

	// db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	// otgorm.AddGormCallbacks(db)

	//err = tableMigration(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func tableMigration(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

func setStatementColumn(s *gorm.Statement, name string, v interface{}) {
	if field := s.Schema.LookUpField(name); field != nil {
		s.SetColumn(name, v)
	}
}

func updateTimestampForCreateCallback(db *gorm.DB) {
	setStatementColumn(db.Statement, "created_on", time.Now().Unix())
}

func updateTimestampForUpdateCallback(db *gorm.DB) {
	setStatementColumn(db.Statement, "modified_on", time.Now().Unix())
}

// func deleteCallback(scope *gorm.Scope) {
// 	if !scope.HasError() {
// 		var extraOption string
// 		if str, ok := scope.Get("gorm:delete_option"); ok {
// 			extraOption = fmt.Sprint(str)
// 		}
//
// 		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
// 		isDelField, hasIsDelField := scope.FieldByName("IsDel")
// 		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
// 			now := time.Now().Unix()
// 			scope.Raw(fmt.Sprintf(
// 				"UPDATE %v SET %v=%v,%v=%v%v%v",
// 				scope.QuotedTableName(),
// 				scope.Quote(deletedOnField.DBName),
// 				scope.AddToVars(now),
// 				scope.Quote(isDelField.DBName),
// 				scope.AddToVars(1),
// 				addExtraSpaceIfExist(scope.CombinedConditionSql()),
// 				addExtraSpaceIfExist(extraOption),
// 			)).Exec()
// 		} else {
// 			scope.Raw(fmt.Sprintf(
// 				"DELETE FROM %v%v%v",
// 				scope.QuotedTableName(),
// 				addExtraSpaceIfExist(scope.CombinedConditionSql()),
// 				addExtraSpaceIfExist(extraOption),
// 			)).Exec()
// 		}
// 	}
// }

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
