package domain

type Tasks struct {
	ID     int    `json:"id" gorm:"column:id"`
	UserID int    `json:"user_id" gorm:"column:user_id"`
	Title  string `json:"title" gorm:"column:title"`
	Desc   string `json:"description" gorm:"column:description"`
	Status string `json:"status" gorm:"column:status"`
}
