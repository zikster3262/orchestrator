package models

type Worker struct {
	Id    int64  `json:"id" gorm:"primaryKey"`
	Token string `json:"token"`
}
