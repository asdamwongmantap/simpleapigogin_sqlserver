package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	TD "simpleapigogin/Controllers"
)

var db *gorm.DB

// func init() {
// 	//open a db connection
// 	var err error
// 	db, err = gorm.Open("mysql", "root@/golang?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	//Migrate the schema
// 	// db.AutoMigrate(&todoModel{})
// }

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", TD.CreateTodo)
		v1.GET("/", TD.FetchAllTodo)
		v1.GET("/:id", TD.FetchSingleTodo)
		v1.PUT("/:id", TD.UpdateTodo)
		v1.DELETE("/:id", TD.DeleteTodo)
	}
	router.Run(":5001")

}

// type (
// 	// todoModel describes a todoModel type
// 	todoModel struct {
// 		gorm.Model
// 		Title     string `json:"title"`
// 		Completed int    `json:"completed"`
// 	}

// 	// transformedTodo represents a formatted todo
// 	transformedTodo struct {
// 		ID        uint   `json:"id"`
// 		Title     string `json:"title"`
// 		Completed bool   `json:"completed"`
// 	}
// )
