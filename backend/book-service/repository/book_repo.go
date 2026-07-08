package repository

import (
	"book-service/domain"
	"gorm.io/gorm"
)

type bookRepo struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) domain.BookRepository {
	return &bookRepo{db: db}
}

func (r *bookRepo) GetAll(level string) ([]domain.Book, error) {
	var books []domain.Book
	query := r.db.Preload("Chapters")
	if level != "" {
		query = query.Where("level = ?", level)
	}
	result := query.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *bookRepo) GetByID(id uint) (*domain.Book, error) {
	var book domain.Book
	result := r.db.Preload("Chapters").First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (r *bookRepo) Create(book *domain.Book) error {
	result := r.db.Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *bookRepo) Update(book *domain.Book) error {
	result := r.db.Save(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *bookRepo) Delete(id uint) error {
	result := r.db.Delete(&domain.Book{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *bookRepo) AddChapter(chapter *domain.Chapter) error {
	result := r.db.Create(chapter)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *bookRepo) UpdateChapter(chapter *domain.Chapter) error {
	result := r.db.Save(chapter)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *bookRepo) DeleteChapter(id uint) error {
	result := r.db.Delete(&domain.Chapter{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}