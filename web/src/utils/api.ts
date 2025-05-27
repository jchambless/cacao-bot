import axios from 'axios';

// Create axios instance with base configuration
const api = axios.create({
  baseURL: '/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor
api.interceptors.request.use(
  (config) => {
    // You can add auth tokens here if needed
    // config.headers.Authorization = `Bearer ${token}`;
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor
api.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    // Handle global errors here
    console.error('API Error:', error);
    return Promise.reject(error);
  }
);

// API functions
export const healthCheck = async () => {
  const response = await api.get('/health');
  return response.data;
};

// Add more API functions as needed
export const getBotInfo = async () => {
  const response = await api.get('/bot');
  return response.data;
};

export const getServerStatus = async () => {
  const response = await api.get('/server/status');
  return response.data;
};

export const getStats = async () => {
  const response = await api.get('/stats');
  return response.data;
};

export const getGuilds = async () => {
  const response = await api.get('/guilds');
  return response.data;
};

export const getGuildInfo = async (guildId: string) => {
  const response = await api.get(`/guilds/${guildId}`);
  return response.data;
};

export default api;
