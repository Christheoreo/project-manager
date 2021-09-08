import axios, { AxiosRequestConfig } from "axios";
import { PluginObject } from "vue";

const apiClient = axios.create({
  baseURL: process.env.BASE_URL,
});

import Vue from "vue";
import { AxiosInstance } from "axios";

declare module "vue/types/vue" {
  export interface Vue {
    $http: AxiosInstance;
  }
}

const Axios: PluginObject<AxiosInstance> = {
  install(Vue, options) {
    Vue.prototype.$http = (request: AxiosRequestConfig) => {
      return apiClient(request);
    };
  },
};

export default Axios;
