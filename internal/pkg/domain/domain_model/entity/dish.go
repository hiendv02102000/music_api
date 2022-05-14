package entity

// Dish struct
type Dish struct {
	ID   int    `gorm:"column:id;primary_key;auto_increment;not null"`
	Name string `gorm:"column:name;"`
	BaseModelWithDeleteAt
}

func (u *Dish) TableName() string {
	return "dish"
}
