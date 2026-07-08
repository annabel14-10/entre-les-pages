package usecase

import (
	"book-service/domain"
)

type bookUsecase struct {
	bookRepo domain.BookRepository
}

func NewBookUsecase(repo domain.BookRepository) domain.BookUsecase {
	return &bookUsecase{
		bookRepo: repo,
	}
}

func (u *bookUsecase) FetchAllBooks(level string) ([]domain.Book, error) {
	return u.bookRepo.GetAll(level)
}

func (u *bookUsecase) AddBook(book *domain.Book) error {
	return u.bookRepo.Create(book)
}

func (u *bookUsecase) FetchBookByID(id uint) (*domain.Book, error) {
	return u.bookRepo.GetByID(id)
}

func (u *bookUsecase) ModifyBook(book *domain.Book) error {
	return u.bookRepo.Update(book)
}

func (u *bookUsecase) RemoveBook(id uint) error {
	return u.bookRepo.Delete(id)
}

func (u *bookUsecase) AddChapter(chapter *domain.Chapter) error {
	return u.bookRepo.AddChapter(chapter)
}

func (u *bookUsecase) ModifyChapter(chapter *domain.Chapter) error {
	return u.bookRepo.UpdateChapter(chapter)
}

func (u *bookUsecase) RemoveChapter(id uint) error {
	return u.bookRepo.DeleteChapter(id)
}