package model

import (
	"time"
)

type Transaction struct {
	ID            uint    `gorm:"primary_key"`
	Total         float64 `json:"total"`
	Paid          float64 `json:"paid"`
	Change        float64 `json:"change"`
	PaymentType   string  `json:"payment_type"`
	PaymentDetail string  `json:"paymentdetail"`
	OrderList     string  `json:"order_list"`
	StaffID       string  `json:"staff_id"`
	CreateAt      time.Time
}
