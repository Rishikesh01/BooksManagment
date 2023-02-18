package store

import (
	"BookManagement/models"
	"gorm.io/gorm"
)

type BookStore interface {
	FindByID(id int) (*models.Book, error)
	FindAll() ([]models.Book, error)
	FindByAuthor(author string) ([]models.Book, error)
	FindByTitle(title string) (*models.Book, error)
	DeleteByID(bookID int) error
	UpdateBook(book *models.Book) error
	Save(book ...models.Book) error
}

type bookStore struct {
	db *gorm.DB
}

func NewBookStore(db *gorm.DB) BookStore {
	return &bookStore{db: db}
}

func (b *bookStore) FindByID(id int) (*models.Book, error) {
	var book models.Book
	if err := b.db.Where("id=?", id).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *bookStore) FindAll() ([]models.Book, error) {
	var books []models.Book
	if err := b.db.Find(&books).Error; err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return books, nil
}

func (b *bookStore) FindByAuthor(author string) ([]models.Book, error) {
	var books []models.Book
	if err := b.db.Where("author=?", author).Find(&books).Error; err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return books, nil
}

func (b *bookStore) FindByTitle(title string) (*models.Book, error) {
	var books models.Book
	if err := b.db.Where("title LIKE ?", title).First(&books).Error; err != nil {
		return nil, err
	}

	return &books, nil
}

func (b *bookStore) DeleteByID(bookID int) error {
	if err := b.db.Delete(&models.Book{}, bookID).Error; err != nil {
		return err
	}
	return nil
}

func (b *bookStore) UpdateBook(book *models.Book) error {
	if err := b.db.Save(book).Error; err != nil {
		return err
	}
	return nil
}

func (b *bookStore) Save(book ...models.Book) error {
	if err := b.db.Create(&book).Error; err != nil {
		return err
	}
	return nil
}
