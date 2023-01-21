package remote

import (
	"time"

	"gitlab.com/chicho69-cesar/proxy/books"
)

type Data struct {
	books books.Books
	server string
	port uint16
	user string
	password string
}

func New(server string, port uint16, user, password string) *Data {
	d := &Data {
		server: server,
		port: port,
		user: user,
		password: password,
	}

	d.load()
	return d
}

func (d *Data) ByID(ID uint) books.Book {
	time.Sleep(2 * time.Second)
	for _, v := range d.books {
		if v.ID == ID {
			return v
		}
	}

	return books.Book {}
}

func (d *Data) All() books.Books {
	time.Sleep(4 * time.Second)
	return d.books
}

func (d *Data) load() {
	d.books = make(books.Books, 0, 5)
	d.books = append(
		d.books, 
		books.Book {ID: 1, Name: "El arte de la guerra", Author: "Sun Tzu"},
		books.Book {ID: 2, Name: "La pelota no entra por azar", Author: "Ferran Soriano"},
		books.Book {ID: 3, Name: "Jesus, CEO", Author: "Laurie Beth"},
		books.Book {ID: 4, Name: "La biografia de Steve Jobs", Author: "Walter Isaacson"},
		books.Book {ID: 5, Name: "Pequeño cerdo capitalista", Author: "Sofia Macias"},
	)
}
