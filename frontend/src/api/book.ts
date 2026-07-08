import { apiClient } from "../config/axios";

export interface Book {
    id: number; 
    title: string; 
    author: string; 
    level: string; 
    image_url: string; 
}

export const BookApi = {
    getAllBooks : async(level?: string | null) => {
        const url = level ? `/books?level=${level}` : '/books'; 
        const response = await apiClient.get<Book[]>(url); 
        return response.data; 
    },

    getBookById : async(id : number) => {
        const response = await apiClient.get<Book>(`/book/${id}`); 
        return response.data; 
    }, 

    createBook: async (data: FormData) => {
        const response = await apiClient.post<Book>('/books', data);
        return response.data;
    },

    deleteBook : async (id: number) => {
        await apiClient.delete(`/books/${id}`); 
    },

    updateBook: async (id: number, data: FormData) => {
        const response = await apiClient.put<Book>(`/books/${id}`, data);
        return response.data;
    },
};