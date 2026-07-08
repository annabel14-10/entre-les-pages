import { useState, useEffect } from "react";
import { isAxiosError } from "axios";
import { BookApi, type Book } from "../api/book";

export const useBooks = () => {
    const [levelFilter, setLevelFilter] = useState<string | null>(null);

    const [books, setBooks] = useState<Book[]>([]); 
    const [error, setError] = useState<string | null>(null); 

    const [isLoading, setLoading] = useState(true); 
    const [isCreating, setCreating] = useState(false); 
    const [isDeleting, setDelete] = useState(false); 

    const [isUpdating, setIsUpdating] = useState(false);

    useEffect(() => {
        let ignore = false; 

        async function fetchFilteredBooks() {
            try {
                const data = await BookApi.getAllBooks(levelFilter); 
                if (!ignore) {
                    setBooks(data || []); 
                }
            } catch (err: unknown) {
                if (!ignore) {
                    if (isAxiosError(err)) {
                        setError(err.response?.data?.error || "Failed to load books"); 
                    } else {
                        setError("An unexpected error occured while fetching books");
                    }
                }
            } finally {
                if (!ignore) {
                    setLoading(false); 
                }
            }
        }

        fetchFilteredBooks();

        return () => {
            ignore = true; 
        };
    }, [levelFilter]);

    const applyLevelFilter = (newLevel: string | null) => {
        setLoading(true);
        setError(null);
        setLevelFilter(newLevel);
    };

    const refetchBooks = async () => {
        setLoading(true); 
        setError(null); 
        try {
            const data = await BookApi.getAllBooks(levelFilter); 
            setBooks(data || []); 
        } catch (err: unknown) {
            if (isAxiosError(err)) {
                setError(err.response?.data?.error || "Failed to load books"); 
            } else {
                setError("An unexpected error occured while fetching books");
            }
        } finally {
            setLoading(false);
        }
    };

    const addBook = async (formData: FormData) => {
        setCreating(true);
        setError(null);
        try {
        const newBook = await BookApi.createBook(formData);
        setBooks((prevBooks) => [...prevBooks, newBook]);
        return { success: true };
        } catch (err: unknown) {
            if (isAxiosError(err)) {
                setError(err.response?.data?.message || "Failed to insert book"); 
            } else {
                setError("An unexpected error occured"); 
            } return {success : false};
        } finally {
            setCreating(false); 
        }
    };

    const removeBook = async (id: number) => {
        setDelete(true); 
        setError(null); 

        try {
            await BookApi.deleteBook(id);
            setBooks((prevBooks) => prevBooks.filter((book) => book.id !== id));
            return { success: true };
        } catch (err: unknown) {
            if (isAxiosError(err)) {
                setError(err.response?.data?.message || "Failed to delete book"); 
            } else {
                setError("An unexpected error occured"); 
            } 
            return { success: false }; 
        } finally {
            setDelete(false); 
        }
    }; 

  const editBook = async (id: number, formData: FormData) => {
    setIsUpdating(true);
    setError(null);
    try {
      const updatedBook = await BookApi.updateBook(id, formData);
      
      setBooks((prevBooks) => 
        prevBooks.map((book) => (book.id === id ? updatedBook : book))
      );
      
      return { success: true };
    } catch (err: unknown) {
      if (isAxiosError(err)) {
        setError(err.response?.data?.error || 'Failed to update book');
      } else {
        setError('An unexpected error occurred while updating.');
      }
      return { success: false };
    } finally {
      setIsUpdating(false);
    }
  };

    return {
        books, 
        error, 
        isLoading, 
        isCreating, 
        isDeleting, 
        levelFilter,
        applyLevelFilter,
        refetchBooks, 
        addBook, 
        removeBook, 
        isUpdating, 
        editBook
    }; 
}