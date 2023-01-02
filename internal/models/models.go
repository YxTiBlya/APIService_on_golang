package models

import "time"

type Contact struct {
	ID            int32  `json:"id"            gorm:"primary_key"`
	Number        string `json:"number"        gorm:"type:varchar(15);unique;not null"`
	Operator_code int32  `json:"operator_code" gorm:"not null"`
	Tag           string `json:"tag"           gorm:"type:varchar(100)"`
	Time_zone     string `json:"time_zone"     gorm:"type:varchar(10)"`
}

type Mailing struct {
	ID         int32     `json:"id"         gorm:"primary_key"`
	Start_time time.Time `json:"start_time" gorm:"type:timestamp without time zone;not null"`
	Message    string    `json:"message"    gorm:"not null"`
	Filters    string    `json:"filters"    gorm:"type:varchar(200);not null"`
	End_time   time.Time `json:"end_time"   gorm:"type:timestamp without time zone;not null"`
}

type Message struct {
	ID         int32     `json:"id"         gorm:"primary_key"`
	Datetime   time.Time `json:"datetime"   gorm:"type:timestamp without time zone;not null"`
	Status     string    `json:"status"     gorm:"type:varchar(20)"`
	Mailing_id int32     `json:"mailing_id" gorm:"not null"`
	Contact_id int32     `json:"contact_id" gorm:"not null"`
}
