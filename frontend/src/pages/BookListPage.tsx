import { useState } from 'react';
import Navbar from '../components/Navbar';
import BookCard from '../components/BookCard';
import AddBookModal from '../components/AddBookModal';
import EditBookModal from '../components/EditBookModal';
import { useBooks } from '../hooks/useBooks';
import { type Book } from '../api/book'; 

export default function BookListPage() {
  const { 
    books, 
    isLoading, 
    error, 
    levelFilter, 
    applyLevelFilter, 
    addBook, 
    isCreating,
    removeBook, 
    editBook,   
    isUpdating 
  } = useBooks();

  const [isModalOpen, setIsModalOpen] = useState(false);

  const [selectedBook, setSelectedBook] = useState<Book | null>(null);

  const levels = ["Beginner", "Intermediate", "Advanced"];

  const handleUpdateClick = (book: Book) => {
    setSelectedBook(book);
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar />
      
      <main className="p-8 max-w-6xl mx-auto">
        <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center mb-6 gap-4">
          <h1 className="text-3xl font-bold text-gray-800">Trending Books</h1>
          
          <button 
            onClick={() => setIsModalOpen(true)}
            className="bg-blue-600 text-white px-4 py-2 rounded-lg font-semibold hover:bg-blue-700 transition shadow-sm"
          >
            + Add New Book
          </button>
        </div>

        <div className="flex gap-2 mb-8 overflow-x-auto pb-2">
          <button
            onClick={() => applyLevelFilter(null)}
            className={`px-4 py-1.5 rounded-full text-sm font-semibold transition whitespace-nowrap ${
              levelFilter === null 
                ? "bg-slate-800 text-white" 
                : "bg-gray-200 text-gray-700 hover:bg-gray-300"
            }`}
          >
            All Levels
          </button>
          
          {levels.map(level => (
            <button
              key={level}
              onClick={() => applyLevelFilter(level)}
              className={`px-4 py-1.5 rounded-full text-sm font-semibold transition whitespace-nowrap ${
                levelFilter === level 
                  ? "bg-slate-800 text-white" 
                  : "bg-gray-200 text-gray-700 hover:bg-gray-300"
              }`}
            >
              {level}
            </button>
          ))}
        </div>
        
        {isLoading && (
          <div className="text-center text-gray-500 py-10 font-medium animate-pulse">
            Loading books...
          </div>
        )}

        {error && (
          <div className="bg-red-100 text-red-700 p-4 rounded mb-6 font-medium shadow-sm">
            {error}
          </div>
        )}

        {!isLoading && !error && books.length === 0 && (
          <div className="text-center py-20 bg-white border border-gray-200 rounded-lg shadow-sm">
            <h2 className="text-xl font-semibold text-gray-700">No books found</h2>
            <p className="text-gray-500 mt-2">
              {levelFilter ? `No ${levelFilter} books available.` : "Time to add a new book!"}
            </p>
          </div>
        )}

        {!isLoading && !error && books.length > 0 && (
          <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
            {books.map((book) => (
              <BookCard 
                key={book.id} 
                book={book} 
                onDelete={removeBook}        
                onUpdate={handleUpdateClick}  
              />
            ))}
          </div>
        )}
      </main>

      <AddBookModal 
        isOpen={isModalOpen} 
        onClose={() => setIsModalOpen(false)} 
        onAdd={addBook} 
        isCreating={isCreating} 
      />

      <EditBookModal 
        key={selectedBook ? selectedBook.id : 'empty-modal'} 
        
        isOpen={selectedBook !== null} 
        onClose={() => setSelectedBook(null)} 
        onEdit={editBook} 
        isUpdating={isUpdating}
        bookToEdit={selectedBook}
      />
    </div>
  );
}