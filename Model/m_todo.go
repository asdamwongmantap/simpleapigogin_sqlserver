package m_todo

import (
	// "encoding/json"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// type TodoModel struct {
// 	gorm.Model
// 	Title     string `form:"firstname"`
// 	Completed int    `json:"completed"`
// }
type (
	// todoModel describes a todoModel type
	TodoModel struct {
		gorm.Model
		Title string `json:"title"`
		// Title     string `form:"title" binding:"required"`
		Completed int `json:"completed"`
	}

	// transformedTodo represents a formatted todo
	TransformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	AppRes struct {
		// Refappname string
		Ref_app_id   string
		Ref_app_name string
		Descr        string
		Url_apps     string
		Img_apps     string
		Isactive     string
		Deleted      string
		Usr_crt      string
		Dtm_crt      string
		Usr_upd      string
		Dtm_upd      string
	}
	AppResJSON struct {
		// Refappname string
		Ref_app_id   string `json:"REF_APP_ID"`
		Ref_app_name string `json:"REF_APP_NAME"`
		Descr        string `json:"DESCR"`
		Url_apps     string `json:"URL_APPS"`
		Img_apps     string `json:"IMG_APPS"`
		Isactive     string `json:"ISACTIVE"`
		Deleted      string `json:"DELETED"`
		Usr_crt      string `json:"USR_CRT"`
		Dtm_crt      string `json:"DTM_CRT"`
		Usr_upd      string `json:"USR_UPD"`
		Dtm_upd      string `json:"DTM_UPD"`
	}
)

func UserIsValid(uName, pwd string) bool {
	// DB simulation
	_uName, _pwd, _isValid := "asdammantap2", "1234", false

	if uName == _uName && pwd == _pwd {
		_isValid = true
	} else {
		_isValid = false
	}

	return _isValid
}
