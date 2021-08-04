import { BASE_URL } from './../Variables';
import type { AxiosInstance, AxiosResponse } from 'axios';
import axios from 'axios';
import { navigate } from 'svelte-routing';

export class Service {
  protected instance: AxiosInstance;
  constructor(uri: string) {
    this.instance = axios.create({
      baseURL: `${BASE_URL}/${uri}`,
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

    this.instance.interceptors.response.use(
      (response) => {
        //

        return response;
      },
      (error) => {
        if (error.response) {
          const response = error.response as AxiosResponse<any>;
          if (response.status === 401) {
            console.log('Redirecting!');

            navigate('/login', { replace: true });
          }
        }
        // Do something with request error
        return Promise.reject(error);
      }
    );
  }
}
