package models

type Activity struct {
	Title       string `json:"title" gorm:"column:Title;primaryKey"`
	Description string `json:"description" gorm:"column:Description"`
	Location    string `json:"location" gorm:"column:Location"`
	ImageName   string `json:"image_name" gorm:"column:ImageName"`
	Time        string `json:"time" gorm:"column:Time"`
}
