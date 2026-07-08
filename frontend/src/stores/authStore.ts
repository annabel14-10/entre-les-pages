import {create} from 'zustand'; 
import { AuthApi, type User } from '../api/auth';

interface AuthState {
    isAuthenticated: boolean; 
    user: User | null;

    login: () => void; 
    logout: () => void; 
    fetchUser: () => Promise<void>;  
}

export const useAuthStore = create<AuthState>((set) => ({
    isAuthenticated: localStorage.getItem('isAuthenticated') === 'true', 
    user : null, 

    login: () => {
        localStorage.setItem('isAuthenticated', 'true'); 
        set({isAuthenticated:true});  
    }, 

    logout: () => {
        localStorage.setItem('isAuthenticated', 'false'); 
        set({isAuthenticated:false, user : null}); 
    },

    fetchUser: async () => {
        try {
            const userData = await AuthApi.getMe(); 
            set({user : userData}); 
        } catch (error) {
            console.error("Failed to fetch user profile", error); 
            set({isAuthenticated : false, user : null}); 
            localStorage.removeItem('isAuthenticated'); 
        }
    }
}));

