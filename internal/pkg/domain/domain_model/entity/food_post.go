package entity

// Dish struct
type FoodPost struct {
	ID       int     `gorm:"column:id;primary_key;auto_increment;not null"`
	Title    string  `gorm:"column:title;"`
	Content  string  `gorm:"column:content"`
	View     int     `gorm:"column:view"`
	UserID   int     `gorm:"column:user_id;"`
	Dish     string  `gorm:"column:dish;"`
	Tags     string  `gorm:"column:tags;"`
	ImageURL *string `gorm:"column:image_url"`
	BaseModelWithDeleteAt
}

func (u *FoodPost) TableName() string {
	return "food_post"
}
