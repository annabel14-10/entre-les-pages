import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import LoginPage from "./pages/LoginPage";
import RegisterPage from "./pages/RegisterPage";
import BookListPage from "./pages/BookListPage";
import ProtectedRoute from "./routes/ProtectedRoute"; 

export default function App() {
  return (
    <BrowserRouter>
      <Routes> 
        <Route path="/" element={<Navigate to="/login" replace/>} />

        <Route path="/login" element={<LoginPage/>} /> 
        <Route path="/register" element={<RegisterPage/>} />

        <Route element={<ProtectedRoute/>}>
          <Route path="/dashboard" element={<BookListPage/>} /> 
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

