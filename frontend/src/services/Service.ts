import { BASE_URL } from './../Variables';
import type { AxiosInstance } from 'axios';
import axios from 'axios';

export class Service {
  protected instance: AxiosInstance;
  constructor() {
    this.instance = axios.create({
      baseURL: `${BASE_URL}/auth`,
    });

    this.instance.interceptors.request.use(
      (config) => {
        if (window) {
          const token = window.localStorage.getItem('token');

          if (token) {
            config.headers = {
              Authorization: `Bearer ${token}`,
            };
          }
        }

        // Do something before request is sent
        return config;
      },
      (error) => {
        // Do something with request error
        return Promise.reject(error);
      }
    );
  }
}
