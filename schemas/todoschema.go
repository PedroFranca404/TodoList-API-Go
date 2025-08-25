package schemas

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	text string
}
