import axios from 'axios';

const service = axios.create({
  baseURL: 'http://localhost:8080',
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json'
  }
});

// 请求拦截器
service.interceptors.request.use(
  config => {
    // 可以在这里添加 token
    const token = localStorage.getItem('token');
    if (token) {
      config.headers['Authorization'] = token;
    }
    return config;
  },
  error => {
    console.error('请求错误:', error);
    return Promise.reject(error);
  }
);

// 响应拦截器
service.interceptors.response.use(
  response => {
    const res = response.data;
    if (res.code !== 200) {
      console.error('响应错误:', res.message);
      return Promise.reject(new Error(res.message || 'Error'));
    } else {
      return res;
    }
  },
  error => {
    console.error('响应错误:', error);
    return Promise.reject(error);
  }
);

// 用户 API
export const userApi = {
  register: (data) => service.post('/user/register', data),
  login: (data) => service.post('/user/login', data),
  getInfo: (uuid) => service.get(`/user/${uuid}`),
  search: (name) => service.get(`/user/search?name=${name}`),
  getFriends: (uuid) => service.get(`/user/friends/${uuid}`),
  addFriend: (data) => service.post('/user/addFriend', data)
};