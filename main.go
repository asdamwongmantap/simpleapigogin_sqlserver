package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	TD "simpleapigogin/Controllers"

	"github.com/rs/cors"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		//return "", fmt.Errorf("$PORT not set")
		port = "5001"
	}
	return ":" + port, nil
}

func main() {

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", TD.CreateTodo)
		v1.GET("/", TD.FetchAllTodo)
		v1.GET("/:id", TD.FetchSingleTodo)
		v1.PUT("/:id", TD.UpdateTodo)
		v1.DELETE("/:id", TD.DeleteTodo)
	}
	c := cors.AllowAll()

	handler := c.Handler(router)
	// fmt.Printf(TD.FetchAllTodo())
	// router.Run(":5001", handler)
	log.Fatal(http.ListenAndServe(addr, handler))

}
