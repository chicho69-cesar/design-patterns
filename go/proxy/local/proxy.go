package local

import "gitlab.com/chicho69-cesar/proxy/books"

type Proxy interface {
	GetByID(ID uint) books.Book
	GetAll() books.Books
}
