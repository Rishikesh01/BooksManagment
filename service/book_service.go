package service

import (
	"Crud/models"
	"Crud/store"
)

type BookService interface {
	GetAllBooks() ([]models.DisplayBook, error)
	GetAuthorsBooks(books models.GetAuthorsBooks) ([]models.DisplayBook, error)
	GetBooksWithTitle(title models.GetBookWithTitle) (models.DisplayBook, error)
	GetBookWithID(book models.GetBook) (models.DisplayBook, error)
	DeleteBook(request models.DeleteRequest) error
	UpdateBook(book models.UpdateBook) error
	CreateBook(book models.SaveBook) error
}

func NewBookService(bookStore store.BookStore) BookService {
	return &bookService{bookStore: bookStore}
}

type bookService struct {
	bookStore store.BookStore
}

func (b *bookService) GetAllBooks() ([]models.DisplayBook, error) {
	val, err := b.bookStore.FindAll()
	if err != nil {
		return nil, err
	}

	displayBooks := b.convertToDisplayBook(val)

	return displayBooks, nil
}

func (b *bookService) GetAuthorsBooks(books models.GetAuthorsBooks) ([]models.DisplayBook, error) {
	val, err := b.bookStore.FindByAuthor(books.Author)
	if err != nil {
		return nil, err
	}
	displayBooks := b.convertToDisplayBook(val)
	return displayBooks, nil
}

func (b *bookService) GetBooksWithTitle(title models.GetBookWithTitle) (models.DisplayBook, error) {
	val, err := b.bookStore.FindByTitle(title.Title)
	if err != nil {
		return models.DisplayBook{}, err
	}
	displayBooks := models.DisplayBook{
		ID:          val.ID,
		Title:       val.Title,
		Author:      val.Author,
		Description: val.Description,
	}
	return displayBooks, nil
}

func (b *bookService) GetBookWithID(book models.GetBook) (models.DisplayBook, error) {
	val, err := b.bookStore.FindByID(book.ID)
	if err != nil {
		return models.DisplayBook{}, err
	}
	displayBooks := models.DisplayBook{
		ID:          val.ID,
		Title:       val.Title,
		Author:      val.Author,
		Description: val.Description,
	}
	return displayBooks, nil
}

func (b *bookService) DeleteBook(request models.DeleteRequest) error {
	if err := b.bookStore.DeleteByID(request.ID); err != nil {
		return err
	}
	return nil
}

func (b *bookService) UpdateBook(book models.UpdateBook) error {
	bk, err := b.bookStore.FindByID(book.ID)
	if err != nil {
		return err
	}
	bk.Author = book.Author
	bk.Title = book.Title
	bk.Description = book.Description
	if err := b.bookStore.UpdateBook(bk); err != nil {
		return err
	}
	return nil
}

func (b *bookService) CreateBook(book models.SaveBook) error {

	if err := b.bookStore.Save(models.Book{
		Author:      book.Author,
		Title:       book.Title,
		Description: book.Description,
	}); err != nil {
		return err
	}
	return nil
}

func (b *bookService) convertToDisplayBook(val []models.Book) []models.DisplayBook {
	displayBooks := make([]models.DisplayBook, len(val))
	for i := 0; i < len(val); i++ {
		displayBooks[i] = models.DisplayBook{
			ID:          val[i].ID,
			Title:       val[i].Title,
			Author:      val[i].Author,
			Description: val[i].Description,
		}
	}

	return displayBooks
}
