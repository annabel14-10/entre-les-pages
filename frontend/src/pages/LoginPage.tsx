import { Link } from 'react-router-dom';
import React, { useState } from 'react';
import { useAuth } from '../hooks/useAuth';

export default function LoginPage() {

  const [email, setEmail] = useState(''); 
  const [password, setPassword] = useState(''); 

  const {login, isLoading, error} = useAuth(); 

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault(); 
    login({email, password}); 
  }

  return (
    <div className="flex h-screen items-center justify-center bg-gray-50">
      <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-sm border border-gray-100">
        <h1 className="text-2xl font-bold text-center mb-6 text-slate-800">Welcome Back</h1>

        {error && (
          <div className="bg-red-100 text-red-700 p-2 rounded mb-4 text-sm text-center">
            {error}
          </div>
        )}
        
        <form onSubmit={handleSubmit} className="flex flex-col gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Email</label>
            <input 
              type="email" 
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="w-full border border-gray-300 rounded p-2 focus:outline-none focus:ring-2 focus:ring-blue-500" 
              placeholder="test@test.com" 
              required
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Password</label>
            <input 
              type="password" 
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full border border-gray-300 rounded p-2 focus:outline-none focus:ring-2 focus:ring-blue-500" 
              placeholder="••••••••" 
              required
            />
          </div>
          <button 
            type="submit" 
            disabled={isLoading}
            className="bg-blue-600 text-white font-semibold py-2 rounded hover:bg-blue-700 transition mt-2 disabled:bg-blue-300"
          >
            {isLoading ? 'Logging in ...' : 'Log in'}
          </button>
        </form>

        <p className="mt-4 text-center text-sm text-gray-600">
          New to ESdemy? <Link to="/register" className="text-blue-600 hover:underline">Create an account</Link>
        </p>
      </div>
    </div>
  );
}