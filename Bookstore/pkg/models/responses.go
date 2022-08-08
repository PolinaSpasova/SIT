package models

type Categories struct {
	AllCat []Category `json:"categories,omitempty"`
}

type Authors struct {
	AllAuthors []Author `json:"authors,omitempty"`
}

type Books struct {
	AllBooks []Book `json:"books,omitempty"`
}

type BookAC struct {
	Id       int64    `gorm:"primaryKey" json:"id,omitempty"`
	Title    string   `json:"title,omitempty"`
	Author   Author   `gorm:"embedded" json:"author,omitempty"`
	Category Category `gorm:"embedded" json:"category,omitempty"`
	Price    float64  `json:"price,omitempty"`
}

type AllBooksAC struct {
	Books []BookAC `json:"books,omitempty"`
}
