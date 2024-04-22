package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"helloword/app/stock/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDB)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

// NewDB 初始化db
func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "social-service/data/gorm"))

	var config gorm.Config
	config.SkipDefaultTransaction = false
	config.DisableForeignKeyConstraintWhenMigrating = true

	dsn := "root:123456@tcp(127.0.0.1:3306)/kratos-im?parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 171,
	}), &config)

	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	db.AutoMigrate()

	return db
}
