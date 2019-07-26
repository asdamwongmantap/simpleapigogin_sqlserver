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
