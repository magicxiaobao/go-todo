package main

import(
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/magicxiaobao/go-todo/config"
	"github.com/magicxiaobao/go-todo/model"
	"github.com/magicxiaobao/go-todo/route"
)

func main() {

	var err error
	println(config.GetDbUrl(config.BuildDbConfig()))
	config.DB, err = gorm.Open("mysql", config.GetDbUrl(config.BuildDbConfig()))
	if err != nil {
		fmt.Println("db open failed", err)
	}
	defer config.DB.Close()
	// migration
	config.DB.AutoMigrate(&model.Todo{})
	engine := route.SetupRoute()
	engine.Run(":8080")
}
