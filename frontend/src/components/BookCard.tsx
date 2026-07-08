import { type Book } from '../api/book';

const API_ORIGIN = import.meta.env.VITE_API_URL.replace(/\/api$/, '');

interface BookCardProps {
  book: Book;
  onDelete: (id: number) => void;
  onUpdate: (book: Book) => void;
}

export default function BookCard({ book, onDelete, onUpdate }: BookCardProps) {
  return (
    <div className="border border-gray-200 rounded-lg p-4 shadow-sm bg-white hover:shadow-md transition flex flex-col">
      <div className="h-32 bg-gray-200 rounded mb-4 flex items-center justify-center text-gray-500 overflow-hidden">
        {book.image_url ? (
          <img src={`${API_ORIGIN}${book.image_url}`} alt={book.title} className="w-full h-full object-cover" />
        ) : (
          "No Image"
        )}
      </div>
      
      <h2 className="text-lg font-bold truncate text-gray-800">{book.title}</h2>
      <p className="text-gray-600 text-sm mb-4">{book.author}</p>
      
      <div className="mt-auto flex justify-between items-center border-t border-gray-100 pt-3">
        <span className="inline-block px-2 py-1 bg-blue-100 text-blue-800 text-xs font-semibold rounded">
          {book.level}
        </span>

        <div className="flex gap-2">
          <button 
            onClick={() => onUpdate(book)}
            className="p-1.5 text-gray-400 hover:text-blue-600 hover:bg-blue-50 rounded transition"
            title="Edit Book"
          >
            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
            </svg>
          </button>

          <button 
            onClick={() => onDelete(book.id)}
            className="p-1.5 text-gray-400 hover:text-red-600 hover:bg-red-50 rounded transition"
            title="Delete Book"
          >
            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2}>
              <path strokeLinecap="round" strokeLinejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
            </svg>
          </button>
        </div>
      </div>
    </div>
  );
}