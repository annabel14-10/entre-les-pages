import { useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import { AuthApi } from "../api/auth";
import { useAuthStore } from "../stores/authStore";

export default function Navbar() {
    const navigate = useNavigate(); 
    
    const { user, fetchUser, logout: setZustandLogout } = useAuthStore();
    
    useEffect(() => {
        if (!user) {
            fetchUser();
        }
    }, [user, fetchUser]);
    
    const handleLogout = async() => {
        try {
            await AuthApi.logout(); 
        } catch(err) {
            console.error("Failed to log out", err); 
        } finally {
            setZustandLogout(); 
            navigate('/login'); 
        }
    }; 

    return (
        <nav className="flex justify-between items-center bg-slate-900 p-4 text-white shadow-md">
            <h1 className="text-xl font-bold">
                <Link to="/dashboard">BookStore</Link>
            </h1>
            
            <div className="flex items-center gap-6">
                <Link to="/profile" className="text-sm font-medium hover:text-blue-300 transition">
                    Hello, {user ? user.username : 'Loading...'}
                </Link>

                <button 
                    onClick={handleLogout}
                    className="bg-slate-700 px-4 py-2 rounded font-semibold hover:bg-slate-600 transition"
                >
                    Logout
                </button>
            </div>
        </nav>
    )
}