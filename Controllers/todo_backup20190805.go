package todo

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"

	// "bytes"
	// "encoding/json"
	// "io/ioutil"
	// "log"

	// "strconv"

	// "github.com/jinzhu/gorm"
	// MTD "simpleapigogin_sqlserver/Model"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// var db *gorm.DB
// var condb *sql.DB

func connect1() (*sql.DB, error) {
	db, err := sql.Open("mssql", "server=172.16.1.186;user id=Appdev;password=1234;database=BAFLiteDB;")
	if err != nil {
		return nil, err
	}

	return db, nil
}

// func init() {
// 	//open a db connection
// 	var err error
// 	// db, err = gorm.Open("mysql", "root@/golang?charset=utf8&parseTime=True&loc=Local")
// 	// if err != nil {
// 	// 	panic("failed to connect database")
// 	// }
// 	condb, err := sql.Open("mssql", "server=172.16.1.186;user id=Appdev;password=1234;database=BAFLiteDB;")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	fmt.Println(condb)
// 	//Migrate the schema
// 	// db.AutoMigrate(MTD.TodoModel{})

// }

// // fetchAllTodo fetch all todos
// func FetchAllTodo(c *gin.Context) {

// 	var todos []MTD.TodoModel
// 	var _todos []MTD.TransformedTodo

// 	db.Find(&todos)

// 	if len(todos) <= 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}

// 	//transforms the todos for building a good response
// 	for _, item := range todos {
// 		completed := false
// 		if item.Completed == 1 {
// 			completed = true
// 		} else {
// 			completed = false
// 		}
// 		_todos = append(_todos, MTD.TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
// 	}
// 	// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
// 	c.JSON(http.StatusOK, _todos)
// }

// // createTodo add a new todo
// func CreateTodo(c *gin.Context) {
// 	var titletdm MTD.TodoModel
// 	completed, _ := strconv.Atoi(c.PostForm("completed"))
// 	// var title := c.PostForm("title")
// 	c.BindJSON(&titletdm)
// 	todo := MTD.TodoModel{Title: titletdm.Title, Completed: completed}
// 	// db.Save(&todo)
// 	db.Save(&todo)
// 	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID, "fieldtitle": titletdm.Title})
// }

// // fetchSingleTodo fetch a single todo
// func FetchSingleTodo(c *gin.Context) {
// 	var todo MTD.TodoModel
// 	todoID := c.Param("id")

// 	db.First(&todo, todoID)

// 	if todo.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}

// 	completed := false
// 	if todo.Completed == 1 {
// 		completed = true
// 	} else {
// 		completed = false
// 	}

// 	_todo := MTD.TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
// }

// // updateTodo update a todo
// func UpdateTodo(c *gin.Context) {
// 	var todo MTD.TodoModel
// 	// var titletdm MTD.TodoModel
// 	todoID := c.Param("id")
// 	// firstname := c.Param("firstname")

// 	db.First(&todo, todoID)

// 	if todo.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}
// 	c.BindJSON(&todo)
// 	db.Model(&todo).Update("title", todo.Title)
// 	completed, _ := strconv.Atoi(c.PostForm("completed"))
// 	db.Model(&todo).Update("completed", completed)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!", "tes": todo.Title})
// }

// // deleteTodo remove a todo
// func DeleteTodo(c *gin.Context) {
// 	var todo MTD.TodoModel
// 	todoID := c.Param("id")

// 	db.First(&todo, todoID)

// 	if todo.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}

// 	db.Delete(&todo)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
// }

// Login todo
func LoginTodo1(c *gin.Context) {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// var id = "BAF Lite"
	// err = db.QueryRow("SELECT REF_APP_ID,REF_APP_NAME FROM REF_APP WHERE REF_APP_NAME= ?", id).Scan(&result.Ref_app_id, &result.Ref_app_name)

	// for get all data sql server
	// rows, err := db.Query("SELECT TOP 10 * FROM REF_APP")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// defer rows.Close()
	// var resultawal []MTD.AppRes
	// var _resultawal []MTD.AppResJSON
	// for rows.Next() {
	// 	var result = MTD.AppRes{}
	// 	err := rows.Scan(&result.Ref_app_id, &result.Ref_app_name,
	// &result.Descr, &result.Url_apps, &result.Img_apps,
	// &result.Isactive, &result.Deleted, &result.Usr_crt,
	// &result.Dtm_crt, &result.Usr_upd, &result.Dtm_upd)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		return
	// 	}
	// 	// log.Println(result)

	// 	resultawal = append(resultawal, result)
	// }

	// if err = rows.Err(); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// for _, result := range resultawal {
	// 	// fmt.Println(result.Ref_app_name)
	// 	// c.JSON(http.StatusOK, gin.H{
	// 	// 	"status": http.StatusOK,
	// 	// 	"result": gin.H{
	// 	// 		"REF_APP_ID": result.Ref_app_id, "REF_APP_NAME": result.Ref_app_name}})
	// 	_resultawal = append(_resultawal, MTD.AppResJSON{Ref_app_id: result.Ref_app_id,
	// 		Ref_app_name: result.Ref_app_name, Descr: result.Descr,
	// 		Url_apps: result.Url_apps, Img_apps: result.Img_apps,
	// 		Isactive: result.Isactive, Deleted: result.Deleted,
	// 		Usr_crt: result.Usr_crt, Dtm_crt: result.Dtm_crt,
	// 		Usr_upd: result.Usr_upd, Dtm_upd: result.Dtm_upd})
	// }
	// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": _resultawal})

	// for get some data sql server
	// var result MTD.AppRes
	// var refapp MTD.AppRes
	// c.BindJSON(&refapp)
	// var id = refapp.Ref_app_name
	// err = db.QueryRow("SELECT * FROM REF_APP WHERE REF_APP_NAME= ?", id).
	// 	Scan(&result.Ref_app_id, &result.Ref_app_name,
	// 		&result.Descr, &result.Url_apps, &result.Img_apps,
	// 		&result.Isactive, &result.Deleted, &result.Usr_crt,
	// 		&result.Dtm_crt, &result.Usr_upd, &result.Dtm_upd)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": result})

	// // for post sql server
	// var result MTD.AppRes
	// c.BindJSON(&result)
	// var RefAppName = result.Ref_app_name
	// stmt, err := db.Prepare("INSERT INTO REF_APP (REF_APP_NAME) VALUES(?)")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// insert, err := stmt.Exec(RefAppName)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// if http.StatusOK == 200 {
	// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": "berhasil input", "hasil": insert})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": "gagal input"})
	// }

	// for put sql server
	// var id = c.Param("id")
	// var result MTD.AppRes
	// var resultinsert MTD.AppRes
	// c.BindJSON(&resultinsert)
	// var RefAppName = resultinsert.Ref_app_name
	// stmt, err := db.Prepare("UPDATE REF_APP SET REF_APP_NAME = ? WHERE REF_APP_ID= ?")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// insert, err := stmt.Exec(RefAppName, id)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// err = db.QueryRow("SELECT * FROM REF_APP WHERE REF_APP_ID= ?", id).
	// 	Scan(&result.Ref_app_id, &result.Ref_app_name,
	// 		&result.Descr, &result.Url_apps, &result.Img_apps,
	// 		&result.Isactive, &result.Deleted, &result.Usr_crt,
	// 		&result.Dtm_crt, &result.Usr_upd, &result.Dtm_upd)

	// if http.StatusOK == 200 {
	// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": result, "ket": "berhasil update", "insres": insert})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": "gagal update"})
	// }

	// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "name": RefAppName, "id": RefAppID})
}
