package Model

import "time"

type CustomerInfo struct {
    Id        uint      `gorm:"primaryKey" json:"id"`
    FirstName string    `json:"first_name"`
    LastName  string    `json:"last_name"`
    Email     string    `json:"email"`
    CreateDT  time.Time `json:"create_dt"`
    UpdateDT  time.Time `json:"update_dt"`
    IsDelete  bool      `json:"is_delete"`
}

func (CustomerInfo) TableName() string {
    return "CustomerInfo"
}