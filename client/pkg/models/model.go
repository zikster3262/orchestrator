package models

type Worker struct {
	Id       string `json:"id" gorm:"primaryKey"`
	Workerid string `json:"workerid"`
	Token    string `json:"token"`
}
