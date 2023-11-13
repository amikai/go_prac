package other

type Config struct {
	ISBN  string
	Name  string
	Pages int
}

type Book struct {
	ISBN  string
	Name  string
	Pages int
}

func NewBook(conf *Config) *Book {
	return &Book{
		ISBN:  conf.ISBN,
		Name:  conf.Name,
		Pages: conf.Pages,
	}
}
