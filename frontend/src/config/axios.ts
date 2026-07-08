import axios from 'axios'; 

const API_URL = import.meta.env.VITE_API_URL;

export const apiClient = axios.create({
    baseURL : import.meta.env.VITE_API_URL,
    withCredentials: true,
});
apiClient.interceptors.response.use(
    (response) => response, 
    async (error) => {
        const originalRequest = error.config;

        if (error.response?.status == 401 && !originalRequest._retry) {
            originalRequest._retry = true; 

            try {
                await axios.post(
                    `${API_URL}/refresh`, 
                    {}, 
                    {withCredentials : true}
                ); 
                return apiClient(originalRequest); 
            } catch (refreshError) {
                console.warn("Session has fully expired. Logging out ... "); 
                localStorage.removeItem('isAuthenticated'); 
                window.location.href = '/login'; 

                return Promise.reject(refreshError); 
            }
        }
        return Promise.reject(error); 
    }
)

