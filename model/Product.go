package model

import (
	"time"
)

//Product ...
type Product struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Stock    int64
	Price    float64
	Image    string
	CreateAt time.Time
}
