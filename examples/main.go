package main

import (
	"github.com/gin-gonic/gin"
	"github.com/niclausse/webflow"
	"github.com/niclausse/webflow/examples/controllers"
	"github.com/niclausse/webkit/resource"
	"github.com/niclausse/webkit/zlog"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	e := gin.New()

	zlog.GetSugaredLogger()
	db, err := resource.InitMySQL(resource.MysqlConf{
		Addr:            "10.117.0.4:3306",
		Database:        "lp_dev",
		User:            "homework",
		Password:        "homework",
		Charset:         "",
		MaxIdleConns:    0,
		MaxOpenConns:    0,
		ConnMaxIdlTime:  0,
		ConnMaxLifeTime: 0,
		ConnTimeOut:     0,
		WriteTimeOut:    0,
		ReadTimeOut:     0,
	}, logger.New(log.Default(), logger.Config{
		SlowThreshold:             0,
		Colorful:                  false,
		IgnoreRecordNotFoundError: false,
		LogLevel:                  0,
	}))
	if err != nil {
		panic(err)
	}

	webflow.SetDefaultDB(db)

	e.GET("/user", webflow.UseController(new(controllers.UserInfoCTL)))

	e.Run(":8090")
}
