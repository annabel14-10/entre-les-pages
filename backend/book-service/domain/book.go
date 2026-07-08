package domain

import "time"

type Chapter struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	BookID    uint      `json:"book_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Book struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Level     string    `json:"level"`
	ImageURL  string    `json:"image_url"`
	Chapters  []Chapter `json:"chapters" gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookRepository interface {
	GetAll(level string) ([]Book, error)
	GetByID(id uint) (*Book, error)
	Create(book *Book) error
	Update(book *Book) error
	Delete(id uint) error

	AddChapter(chapter *Chapter) error
	UpdateChapter(chapter *Chapter) error
	DeleteChapter(id uint) error
}

type BookUsecase interface {
	FetchAllBooks(level string) ([]Book, error)
	FetchBookByID(id uint) (*Book, error)
	AddBook(book *Book) error
	ModifyBook(book *Book) error
	RemoveBook(id uint) error

	AddChapter(chapter *Chapter) error
	ModifyChapter(chapter *Chapter) error
	RemoveChapter(id uint) error
}