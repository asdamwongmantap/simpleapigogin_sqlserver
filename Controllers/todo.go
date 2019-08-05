package todo

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	// "bytes"
	// "encoding/json"
	// "io/ioutil"
	"log"

	// "strconv"

	// "github.com/jinzhu/gorm"
	// MTD "simpleapigogin_sqlserver/Model"
	MTD "simpleapigogin_sqlserver/Model"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// var db *gorm.DB
// var condb *sql.DB

func connect() (*sql.DB, error) {
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

// fetchAllTodo fetch all todos
func FetchAllTodo(c *gin.Context) {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT TOP 10 * FROM REF_APP")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var resultawal []MTD.AppRes
	var _resultawal []MTD.AppResJSON
	for rows.Next() {
		var result = MTD.AppRes{}
		err := rows.Scan(&result.Ref_app_id, &result.Ref_app_name,
			&result.Descr, &result.Url_apps, &result.Img_apps,
			&result.Isactive, &result.Deleted, &result.Usr_crt,
			&result.Dtm_crt, &result.Usr_upd, &result.Dtm_upd)
		if err != nil {
			log.Fatal(err)
			return
		}
		// log.Println(result)

		resultawal = append(resultawal, result)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, result := range resultawal {
		_resultawal = append(_resultawal, MTD.AppResJSON{Ref_app_id: result.Ref_app_id,
			Ref_app_name: result.Ref_app_name, Descr: result.Descr,
			Url_apps: result.Url_apps, Img_apps: result.Img_apps,
			Isactive: result.Isactive, Deleted: result.Deleted,
			Usr_crt: result.Usr_crt, Dtm_crt: result.Dtm_crt,
			Usr_upd: result.Usr_upd, Dtm_upd: result.Dtm_upd})
	}
	// c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": _resultawal})
	c.JSON(http.StatusOK, _resultawal)
}

// createTodo add a new todo
func CreateTodo(c *gin.Context) {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	var result MTD.DataUser
	c.BindJSON(&result)
	var username = result.Username
	var password = result.Password
	var fullname = result.Fullname
	hasher := md5.New()
	hasher.Write([]byte(password))
	var passwordmd5 = hex.EncodeToString(hasher.Sum(nil))
	stmt, err := db.Prepare("INSERT INTO USER_ACCESS (USERNAME,PASSWORD,FULLNAME) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insert, err := stmt.Exec(username, passwordmd5, fullname)
	if err != nil {
		panic(err.Error())
	}

	if http.StatusOK == 200 {
		c.JSON(http.StatusOK, gin.H{"status": "sesuai", "result": "berhasil input", "hasil": insert})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "tidak sesuai", "result": "gagal input"})
	}
}

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
func LoginTodo(c *gin.Context) {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// for get some data sql server
	var result MTD.AppResLogin
	var refapp MTD.AppResLogin
	c.BindJSON(&refapp)
	var username = refapp.Username
	var passwordasli = refapp.Password

	hasher := md5.New()
	hasher.Write([]byte(passwordasli))
	var passwordmd5 = hex.EncodeToString(hasher.Sum(nil))
	err = db.QueryRow("SELECT USERNAME,PASSWORD,FULLNAME FROM USER_ACCESS WHERE USERNAME= ? AND PASSWORD= ?", username, passwordmd5).
		Scan(&result.Username, &result.Password, &result.Fullname)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if http.StatusOK == 200 {
		c.JSON(http.StatusOK, gin.H{"status": "sesuai", "username": result.Username, "password": result.Password})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "tidak sesuai", "username": result.Username, "password": result.Password})
	}

}
