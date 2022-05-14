package entity

// Dish struct
type FoodPost struct {
	ID      int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Title   string `gorm:"column:title;"`
	Content string `gorm:"column:content"`
	UserID  int    `gorm:"column:user_id;"`
	DishID  int    `gorm:"column:dish_id;"`
}

func (u *FoodPost) TableName() string {
	return "food_post"
}
