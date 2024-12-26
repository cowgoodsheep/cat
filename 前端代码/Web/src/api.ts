// api.ts
// import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';
import axios from 'axios'
import type { InternalAxiosRequestConfig, AxiosResponse, AxiosError } from 'axios';


// 创建 Axios 实例
const api = axios.create({
    baseURL: '/api', // 后端 API 地址
    timeout: 10000, // 请求超时时间
});

// 请求拦截器
api.interceptors.request.use(
    (config: InternalAxiosRequestConfig) => {
        // 在请求发送之前做一些处理，比如加入请求头信息
        // 这里可以添加用户身份验证等
        return config;
    },
    (error) => {
        // 对请求错误做些什么
        return Promise.reject(error);
    }
);

// 响应拦截器
api.interceptors.response.use(
    (response: AxiosResponse) => {
        // 对响应数据做一些处理
        return response;
    },
    (error) => {
        // 对响应错误做些什么
        return Promise.reject(error);
    }
);

// 封装请求函数
export const get = async <T = any>(url: string, config?: InternalAxiosRequestConfig) => {
    try {
        const response = await api.get<T>(url, config);
        return response.data;
    } catch (error) {
        throw error;
    }
};

export const post = async <T = any>(url: string, data?: any, config?: InternalAxiosRequestConfig) => {
    try {
        const response = await api.post<T>(url, data, config);
        return response.data;
    } catch (error) {
        throw error;
    }
};

export const put = async <T = any>(url: string, data?: any, config?: InternalAxiosRequestConfig) => {
    try {
        const response = await api.put<T>(url, data, config);
        return response.data;
    } catch (error) {
        throw error;
    }
};
export const remove = async <T = any>(url: string, data?: any, config?: InternalAxiosRequestConfig) => {
    try {
        const response = await api.delete<T>(url, { ...config, data });
        return response.data;
    } catch (error) {
        throw error;
    }
};
