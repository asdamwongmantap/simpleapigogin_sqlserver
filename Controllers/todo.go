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
	// QueryRow("SELECT * FROM REF_APP WHERE REF_APP_NAME= ?", id).
	rows, err := db.Query("SELECT TOP 10* FROM REF_APP ORDER BY DTM_CRT DESC")
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

	// if err = rows.Err(); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

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
	var txt MTD.AppSP
	c.BindJSON(&txt)
	var Refappid = ""
	var Refappname = txt.Ref_app_name
	var Descr = txt.Descr
	var Urlapps = txt.Url_apps
	var Imgapps = txt.Img_apps
	var Isactive = txt.Isactive
	var Usrcrt = txt.Usr_crt
	// hasher := md5.New()
	// hasher.Write([]byte(password))
	// var passwordmd5 = hex.EncodeToString(hasher.Sum(nil))
	stmt, err := db.Prepare("EXEC SPINSERTUPDATEREFAPP ?,?,?,?,?,?,?")
	if err != nil {
		panic(err.Error())
	}

	insert, err := stmt.Exec(Refappid, Refappname, Descr, Urlapps, Imgapps, Isactive, Usrcrt)
	if err != nil {
		panic(err.Error())
	}

	if http.StatusOK == 200 {
		c.JSON(http.StatusOK, gin.H{"status": "sesuai", "result": "berhasil input", "hasil": insert})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "tidak sesuai", "result": "gagal input"})
	}
}

// fetchSingleTodo fetch a single todo
func FetchSingleTodo(c *gin.Context) {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	// for get some data sql server
	var result MTD.AppRes
	// var refapp MTD.AppRes
	// c.BindJSON(&refapp)
	// var id = refapp.Ref_app_name
	var id = c.Param("id")
	err = db.QueryRow("SELECT * FROM REF_APP WHERE REF_APP_ID= ?", id).
		Scan(&result.Ref_app_id, &result.Ref_app_name,
			&result.Descr, &result.Url_apps, &result.Img_apps,
			&result.Isactive, &result.Deleted, &result.Usr_crt,
			&result.Dtm_crt, &result.Usr_upd, &result.Dtm_upd)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
}

// updateTodo update a todo
func UpdateTodo(c *gin.Context) {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	// for put sql server
	var id = c.Param("id")
	var result MTD.AppRes
	// var resultinsert MTD.AppRes
	// c.BindJSON(&resultinsert)
	// var RefAppName = resultinsert.Ref_app_name
	var txt MTD.AppSP
	c.BindJSON(&txt)
	var Refappid = id
	var Refappname = txt.Ref_app_name
	var Descr = txt.Descr
	var Urlapps = txt.Url_apps
	var Imgapps = txt.Img_apps
	var Isactive = txt.Isactive
	var Usrcrt = txt.Usr_crt
	stmt, err := db.Prepare("EXEC SPINSERTUPDATEREFAPP ?,?,?,?,?,?,?")
	if err != nil {
		panic(err.Error())
	}

	insert, err := stmt.Exec(Refappid, Refappname, Descr, Urlapps, Imgapps, Isactive, Usrcrt)
	if err != nil {
		panic(err.Error())
	}
	err = db.QueryRow("SELECT * FROM REF_APP WHERE REF_APP_ID= ?", id).
		Scan(&result.Ref_app_id, &result.Ref_app_name,
			&result.Descr, &result.Url_apps, &result.Img_apps,
			&result.Isactive, &result.Deleted, &result.Usr_crt,
			&result.Dtm_crt, &result.Usr_upd, &result.Dtm_upd)

	if http.StatusOK == 200 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": result, "ket": "berhasil update", "insres": insert})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "ket": "gagal update"})
	}
}

// deleteTodo remove a todo
func DeleteTodo(c *gin.Context) {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	// for put sql server
	var id = c.Param("id")
	var result MTD.AppRes
	// var resultinsert MTD.AppRes
	// c.BindJSON(&resultinsert)
	// var RefAppName = resultinsert.Ref_app_name
	stmt, err := db.Prepare("DELETE FROM REF_APP WHERE REF_APP_ID= ?")
	if err != nil {
		panic(err.Error())
	}

	insert, err := stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}
	err = db.QueryRow("SELECT * FROM REF_APP WHERE REF_APP_ID= ?", id).
		Scan(&result.Ref_app_id, &result.Ref_app_name,
			&result.Descr, &result.Url_apps, &result.Img_apps,
			&result.Isactive, &result.Deleted, &result.Usr_crt,
			&result.Dtm_crt, &result.Usr_upd, &result.Dtm_upd)

	if http.StatusOK == 200 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": result, "ket": "berhasil update", "insres": insert})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": "gagal update"})
	}
}

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
