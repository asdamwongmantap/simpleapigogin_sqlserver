package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	MTD "simpleapigogin/Model"
)

var db *gorm.DB

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	// db.AutoMigrate(MTD.TodoModel{})
}

// fetchAllTodo fetch all todos
func FetchAllTodo(c *gin.Context) {

	var todos []MTD.TodoModel
	var _todos []MTD.TransformedTodo

	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	//transforms the todos for building a good response
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, MTD.TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
	c.JSON(http.StatusOK, _todos)
}

// createTodo add a new todo
func CreateTodo(c *gin.Context) {
	var titletdm MTD.TodoModel
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	// var title := c.PostForm("title")
	c.BindJSON(&titletdm)
	todo := MTD.TodoModel{Title: titletdm.Title, Completed: completed}
	// db.Save(&todo)
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID, "fieldtitle": titletdm.Title})
}

// fetchSingleTodo fetch a single todo
func FetchSingleTodo(c *gin.Context) {
	var todo MTD.TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_todo := MTD.TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}

// updateTodo update a todo
func UpdateTodo(c *gin.Context) {
	var todo MTD.TodoModel
	// var titletdm MTD.TodoModel
	todoID := c.Param("id")
	// firstname := c.Param("firstname")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	c.BindJSON(&todo)
	db.Model(&todo).Update("title", todo.Title)
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!", "tes": todo.Title})
}

// deleteTodo remove a todo
func DeleteTodo(c *gin.Context) {
	var todo MTD.TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}

// Login todo

func LoginTodo(c *gin.Context) {
	var todo MTD.TodoModel
	var titletdm MTD.TodoModel
	// todoID := c.Param("id")
	// todotitle := c.Param("title")
	c.BindJSON(&titletdm)
	name := titletdm.Title
	// name := c.PostForm("title")
	// sqlStatement := `SELECT col FROM my_table WHERE id=$1`
	// db.First(&todo, todotitle)
	db.Where("title = ?", name).Find(&todo)

	if titletdm.Title == todo.Title {
		// if c.PostForm("title") == todo.Title {
		// c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		c.JSON(http.StatusOK, gin.H{"statusaja": "sesuai"})
		// c.JSON(http.StatusOK, todo.Title)
		// c.Redirect("")
		return
	} else {
		// c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!", "tes": todo})
		c.JSON(http.StatusOK, gin.H{"statusaja": "tidak sesuai"})
		return
	}

	// completed := false
	// if todo.Completed == 1 {
	// 	completed = true
	// } else {
	// 	completed = false
	// }

	// _todo := MTD.TransformedTodo{ID: todo.ID, Title: todo.Title}
	// c.JSON(http.StatusOK, _todo)
}
