package entity

// tag struct
type Tag struct {
	ID         int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Name       string `gorm:"column:name;"`
	FoodPostID int    `gorm:"column:food_post_id;"`
	BaseModel
}

func (u *Tag) TableName() string {
	return "tag"
}
