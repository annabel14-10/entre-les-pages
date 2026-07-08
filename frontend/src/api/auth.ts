import { apiClient } from "../config/axios";

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
    username: string; 
    dob: string; 
    email: string; 
    password: string;
}

export interface AuthResponse {
  message: string;
}

export interface User {
    id: number; 
    username: string; 
    email: string; 
}
export const AuthApi = {
    login: async (data:LoginRequest) => {
        const response = await apiClient.post<AuthResponse>('/login', data); 
        return response.data 
    }, 

    register: async (data:RegisterRequest) => {
        const response = await apiClient.post<AuthResponse>('/register', data); 
        return response.data
    },

    logout: async() => {
        await apiClient.post('/logout'); 
    }, 

    getMe : async() => {
        const response = await apiClient.get<User>('/me'); 
        return response.data; 
    }
}
