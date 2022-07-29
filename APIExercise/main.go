package main

import (
	"apiexercise/obj"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	DBname := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "polina", "password", "bookstore")
	db, err := sql.Open("postgres", DBname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := echo.New()

	router.GET("/categories", func(ctx echo.Context) error {
		rows, err := db.Query("SELECT * FROM categories")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		cat := []obj.Category{}
		for rows.Next() {
			var c obj.Category
			rows.Scan(&c.Id, &c.Name)
			rowsB, err := db.Query("SELECT * FROM books WHERE c_id=$1", c.Id)
			if err != nil {
				log.Fatal(err)
			}
			books := []obj.Book{}
			for rowsB.Next() {
				var b obj.Book
				rowsB.Scan(&b.Id, &b.Title, &b.AuthorId, &b.CategoryId, &b.Price)
				books = append(books, b)
			}
			c.Books = books
			cat = append(cat, c)
		}
		var result obj.Categories
		result.AllCat = cat

		return ctx.JSON(http.StatusOK, result)
	})

	router.POST("/categories", func(ctx echo.Context) error {
		c := new(obj.Category)
		if err := ctx.Bind(c); err != nil {
			return err
		}
		err = db.QueryRow("INSERT INTO categories (name) VALUES ($1) RETURNING id", c.Name).Scan(&c.Id)
		if err != nil {
			log.Fatal(err)
		}
		return ctx.JSON(http.StatusOK, c)
	})

	router.GET("/categories/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		res := new(obj.Category)
		err := db.QueryRow("SELECT * FROM categories WHERE id=($1)", id).Scan(&res.Id, &res.Name)
		if err != nil {
			log.Fatal(err)
		}
		return ctx.JSON(http.StatusOK, res)
	})

	router.PUT("/categories/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		cat := new(obj.Category)
		if err := ctx.Bind(cat); err != nil {
			return err
		}
		_, err := db.Exec("UPDATE categories SET name=$1 WHERE id=$2", cat.Name, id)
		if err != nil {
			log.Fatal(err)
		}
		err = db.QueryRow("SELECT * FROM categories WHERE id=($1)", id).Scan(&cat.Id, &cat.Name)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(http.StatusOK, cat)
	})

	router.DELETE("/categories/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		row := db.QueryRow("SELECT * FROM categories WHERE id=($1)", id)
		var c obj.Category
		row.Scan(&c.Id, &c.Name)

		_, err := db.Exec(" DELETE FROM books WHERE c_id=$1", id)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(" DELETE FROM categories WHERE id=$1", id)
		if err != nil {
			log.Fatal(err)
		}
		return ctx.JSON(http.StatusOK, c)
	})

	router.GET("/authors", func(ctx echo.Context) error {
		rows, err := db.Query("SELECT * FROM authors")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		auth := []obj.Author{}
		for rows.Next() {
			var a obj.Author
			rows.Scan(&a.Id, &a.Name, &a.Biography)
			rowsB, err := db.Query("SELECT * FROM books WHERE c_id=$1", a.Id)
			if err != nil {
				log.Fatal(err)
			}
			books := []obj.Book{}
			for rowsB.Next() {
				var b obj.Book
				rowsB.Scan(&b.Id, &b.Title, &b.AuthorId, &b.CategoryId, &b.Price)
				books = append(books, b)
			}
			a.Books = books
			auth = append(auth, a)
		}
		var result obj.Authors
		result.AllAuthors = auth

		return ctx.JSON(http.StatusOK, result)
	})

	router.POST("/authors", func(ctx echo.Context) error {
		a := new(obj.Author)
		if err := ctx.Bind(a); err != nil {
			return err
		}
		err = db.QueryRow("INSERT INTO authors (name,biography) VALUES ($1, $2) RETURNING id", a.Name, a.Biography).Scan(&a.Id)
		if err != nil {
			log.Fatal(err)
		}
		return ctx.JSON(http.StatusOK, a)
	})

	router.GET("/authors/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		rows, err := db.Query("SELECT * FROM authors WHERE id=$1", id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var auth obj.Author
		for rows.Next() {
			rows.Scan(&auth.Id, &auth.Name, &auth.Biography)
			rowsB, err := db.Query("SELECT * FROM books WHERE c_id=$1", auth.Id)
			if err != nil {
				log.Fatal(err)
			}
			books := []obj.Book{}
			for rowsB.Next() {
				var b obj.Book
				rowsB.Scan(&b.Id, &b.Title, &b.AuthorId, &b.CategoryId, &b.Price)
				books = append(books, b)
			}
			auth.Books = books
		}

		return ctx.JSON(http.StatusOK, auth)
	})

	router.PUT("/authors/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		a := new(obj.Author)
		if err := ctx.Bind(a); err != nil {
			return err
		}
		_, err := db.Exec("UPDATE authors SET name=$1,biography=$2 WHERE id=$3", a.Name, a.Biography, id)
		if err != nil {
			log.Fatal(err)
		}
		rows, err := db.Query("SELECT * FROM authors WHERE id=($1)", id)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&a.Id, &a.Name, &a.Biography)
		}
		return ctx.JSON(http.StatusOK, a)
	})

	router.DELETE("/authors/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		row := db.QueryRow("SELECT * FROM authors WHERE id=($1)", id)
		var a obj.Author
		row.Scan(&a.Id, &a.Name, &a.Biography)

		_, err := db.Exec(" DELETE FROM books WHERE a_id=$1", id)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(" DELETE FROM authors WHERE id=$1", id)
		if err != nil {
			log.Fatal(err)
		}
		return ctx.JSON(http.StatusOK, a)
	})

	router.GET("/books", func(ctx echo.Context) error {
		rows, err := db.Query("SELECT * FROM books")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		book := []obj.BookAC{}
		for rows.Next() {
			var b obj.BookAC
			aid, cid := 0, 0
			rows.Scan(&b.Id, &b.Title, &aid, &cid, &b.Price)
			err := db.QueryRow("SELECT * FROM authors WHERE id=$1", aid).Scan(&b.Author.Id, &b.Author.Name, &b.Author.Biography)
			if err != nil {
				log.Fatal(err)
			}
			err = db.QueryRow("SELECT * FROM categories WHERE id=$1", cid).Scan(&b.Category.Id, &b.Category.Name)
			if err != nil {
				log.Fatal(err)
			}
			book = append(book, b)
		}
		var result obj.Books
		result.AllBooks = book

		return ctx.JSON(http.StatusOK, result)
	})

	router.POST("/books", func(ctx echo.Context) error {
		b := new(obj.Book)
		if err := ctx.Bind(b); err != nil {
			return err
		}
		err = db.QueryRow("INSERT INTO books (title,a_id,c_id,price) VALUES ($1 , $2,$3,$4) RETURNING id", b.Title, b.AuthorId, b.CategoryId, b.Price).Scan(&b.Id)
		if err != nil {
			log.Fatal(err)
		}
		result := new(obj.BookAC)
		err := db.QueryRow("SELECT id,title,price FROM books WHERE id=$1", b.Id).Scan(&result.Id, &result.Title, &result.Price)
		if err != nil {
			log.Fatal(err)
		}
		err = db.QueryRow("Select * FROM authors WHERE id=$1", b.AuthorId).Scan(&result.Author.Id, &result.Author.Name, &result.Author.Biography)
		if err != nil {
			log.Fatal(err)
		}
		err = db.QueryRow("Select * FROM categories WHERE id=$1", b.CategoryId).Scan(&result.Category.Id, &result.Category.Name)
		if err != nil {
			log.Fatal(err)
		}
		return ctx.JSON(http.StatusOK, result)
	})

	router.GET("/books/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		var book obj.BookAC
		aid, cid := 0, 0
		err := db.QueryRow("SELECT * FROM books WHERE id=$1", id).Scan(&book.Id, &book.Title, &aid, &cid, &book.Price)
		if err != nil {
			log.Fatal(err)
		}
		err = db.QueryRow("SELECT * FROM authors WHERE id=$1", aid).Scan(&book.Author.Id, &book.Author.Name, &book.Author.Biography)
		if err != nil {
			log.Fatal(err)
		}
		err = db.QueryRow("SELECT * FROM categories WHERE id=$1", cid).Scan(&book.Category.Id, &book.Category.Name)
		if err != nil {
			log.Fatal(err)
		}
		return ctx.JSON(http.StatusOK, book)
	})

	router.PUT("/books/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		aid, cid := 0, 0
		book := new(obj.BookAC)
		if err := ctx.Bind(book); err != nil {
			return err
		}
		_, err := db.Exec("UPDATE books SET price=$1 WHERE id=$2", book.Price, id)
		if err != nil {
			log.Fatal(err)
		}
		err = db.QueryRow("SELECT * FROM books WHERE id=$1", id).Scan(&book.Id, &book.Title, &aid, &cid, &book.Price)
		if err != nil {
			log.Fatal(err)
		}
		err = db.QueryRow("SELECT * FROM authors WHERE id=$1", aid).Scan(&book.Author.Id, &book.Author.Name, &book.Author.Biography)
		if err != nil {
			log.Fatal(err)
		}
		err = db.QueryRow("SELECT * FROM categories WHERE id=$1", cid).Scan(&book.Category.Id, &book.Category.Name)
		if err != nil {
			log.Fatal(err)
		}
		return ctx.JSON(http.StatusOK, book)
	})

	router.DELETE("/books/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		var book obj.BookAC
		aid, cid := 0, 0
		err := db.QueryRow("SELECT * FROM books WHERE id=$1", id).Scan(&book.Id, &book.Title, &aid, &cid, &book.Price)
		if err != nil {
			log.Fatal(err)
		}
		err = db.QueryRow("SELECT * FROM authors WHERE id=$1", aid).Scan(&book.Author.Id, &book.Author.Name, &book.Author.Biography)
		if err != nil {
			log.Fatal(err)
		}
		err = db.QueryRow("SELECT * FROM categories WHERE id=$1", cid).Scan(&book.Category.Id, &book.Category.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(" DELETE FROM books WHERE id=$1", id)
		if err != nil {
			log.Fatal(err)
		}

		return ctx.JSON(http.StatusOK, book)
	})

	router.Logger.Fatal(router.Start(":3000"))

}
