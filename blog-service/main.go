package main

import (
	"database/sql"
	"github.com/go-programming-tour-book/blog-service/global"
	router "github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

var (
	db *sql.DB
)

func init(){
	err := setupSetting()
	if err!=nil{
		log.Fatalf("init.setupSetting err:%v/n",err)
	}
	initConn()
}

func setupSetting() error{
	setting, err := setting.NewSetting()
	if err !=nil{
		return err
	}
	setting.ReadSection("Database",&global.DatabaseSetting)
	return nil
}

func initConn()  {
	db, _ = sql.Open(global.DatabaseSetting.DBType,
		global.DatabaseSetting.Username+":"+global.DatabaseSetting.Pwd+"@tcp("+global.DatabaseSetting.Host+")/"+global.DatabaseSetting.Database)
	db.SetMaxOpenConns( 10 )
	db.SetMaxIdleConns( 5 )
	db.Ping()
}


func main(){

	// 1.go get -u github.com/gin-gonic/gin@v1.5.0
	router := router.NewRouter(db)
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

