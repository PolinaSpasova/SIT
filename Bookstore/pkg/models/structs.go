package models

type Book struct {
	Id         int64   `gorm:"primaryKey"  json:"id,omitempty"`
	Title      string  `json:"title,omitempty"`
	AuthorId   int64   `gorm:"foringKey" json:"author,omitempty"`
	CategoryId int64   `gorm:"foringKey" json:"category,omitempty"`
	Price      float64 `json:"price,omitempty"`
}

type Category struct {
	Id    int64  `gorm:"primaryKey"  json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Books []Book `json:"books,omitempty"`
}

type Author struct {
	Id        int64  `gorm:"primaryKey" json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Biography string `json:"biography,omitempty"`
	Books     []Book `json:"books,omitempty"`
}
