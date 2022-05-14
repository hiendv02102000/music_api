package entity

// Commet struct
type Comment struct {
	ID         int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Content    string `gorm:"column:content;"`
	UserID     int    `gorm:"column:user_id;"`
	FoodPostID int    `gorm:"column:food_post_id;"`
	BaseModelWithDeleteAt
}

func (u *Comment) TableName() string {
	return "comment"
}
