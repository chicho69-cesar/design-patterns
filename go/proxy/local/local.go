package local

import (
	"gitlab.com/chicho69-cesar/proxy/books"
	"gitlab.com/chicho69-cesar/proxy/remote"
)

type Local struct {
	Remote *remote.Data
	cache books.Books
}

func New() Proxy {
	return &Local {
		Remote: remote.New("https://some.com", 8080, "root", "1234"),
		cache: make(books.Books, 0),
	}
}

func (l *Local) GetByID(ID uint) books.Book {
	for _, v := range l.cache {
		if v.ID == ID {
			return v
		}
	}

	b := l.Remote.ByID(ID)
	l.cache = append(l.cache, b)

	return b
}

func (l *Local) GetAll() books.Books {
	return l.Remote.All()
}
