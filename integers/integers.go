package integers

import "fmt"

var Three = 3

type books struct {
	name   string
	author string
	price  float64
	pages  int
}

func newBooks(name string) (books, error) {
	b := books{
		name:   name,
		author: "",
		price:  0,
		pages:  0,
	}
	return b, nil
}

func (b *books) String() string {
	return fmt.Sprintf("%s by %s, $%.2f, %d pages", b.name, b.author, b.price, b.pages)
}
