package models

type Worker struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Workerid string `json:"workerid"`
	Token    string `json:"token"`
}
