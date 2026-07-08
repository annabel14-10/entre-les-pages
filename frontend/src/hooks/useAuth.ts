import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { AuthApi, type LoginRequest, type RegisterRequest } from '../api/auth';
import { isAxiosError } from 'axios';
import { useAuthStore } from "../stores/authStore";

export const useAuth = () => {
    const navigate = useNavigate(); 
    const [isLoading, setLoading] = useState(false); 
    const [error, setError] = useState<string | null>(null); 

    const {login: setZustandLogin, logout: setZustandLogout} = useAuthStore(); 
    
    const login = async (data: LoginRequest) => {
        setLoading(true); 
        setError(null); 

        try {
            await AuthApi.login(data); 
            setZustandLogin(); 
            navigate('/dashboard'); 
        } catch (err: unknown) {
            if (isAxiosError(err)) {
                setError(err.response?.data?.message || "Failed to login"); 
            } else {
                setError("An unexpected error occured"); 
            }
        } finally {
            setLoading(false); 
        }
    }

    const register = async (data: RegisterRequest) => {
        setLoading(true); 
        setError(null); 
        try{
            await AuthApi.register(data); 
            setZustandLogin(); 
            navigate('/dashboard'); 
        } catch (err: unknown) {
            if (isAxiosError(err)) {
                setError(err.response?.data?.message || "Failed to Register"); 
            } else {
                setError("An unexpected error occured"); 
            } 
        } finally {
            setLoading(false); 
        }
    }

    const logout = async() => {
        try {
            await AuthApi.logout(); 
        } catch (err) {
            console.error("Failed to logout on the backend", err); 
        } finally {
            setZustandLogout(); 
            navigate('/login'); 
        }
    }

    return {login, register, logout, isLoading, error}; 

}

