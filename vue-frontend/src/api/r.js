import axios from "axios";
import { ElMessage } from "element-plus";

export const r = axios.create({
  baseURL: "/api",
});
let token = localStorage.getItem("token");
r.interceptors.request.use((config) => {
  if (token) {
    config.headers["Authorization"] = "Bearer " + token;
  }
  return config;
});

r.interceptors.response.use(
  (response) => {
    if (response.data.code === 1) {
      if (!response.data.data) {
        ElMessage.success({
          message: response.data.message,
        });
      }
      return Promise.resolve(response.data.data);
    } else {
      ElMessage.error({
        message: response.data.message,
      });
    }
  },
  (error) => {
    console.log(error);
    if (error.response.status === 403) {
      window.location.replace(error.response.data);
    }
    return Promise.reject(error);
  }
);
