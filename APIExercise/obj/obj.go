package obj

type Category struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Books []Book `json:"books,omitempty"`
}

type Book struct {
	Id         int64   `json:"id,omitempty"`
	Title      string  `json:"title,omitempty"`
	AuthorId   int64   `json:"author,omitempty"`
	CategoryId int64   `json:"category,omitempty"`
	Price      float64 `json:"price,omitempty"`
}

type Author struct {
	Id        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Biography string `json:"biography,omitempty"`
	Books     []Book `json:"books,omitempty"`
}

type Categories struct {
	AllCat []Category `json:"categories,omitempty"`
}

type Authors struct {
	AllAuthors []Author `json:"authors,omitempty"`
}

type BookAC struct {
	Id       int64    `json:"id,omitempty"`
	Title    string   `json:"title,omitempty"`
	Author   Author   `json:"author,omitempty"`
	Category Category `json:"category,omitempty"`
	Price    float64  `json:"price,omitempty"`
}

type Books struct {
	AllBooks []BookAC `json:"books,omitempty"`
}
