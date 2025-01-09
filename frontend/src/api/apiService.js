import axios from "axios";

const API_BASE_URL = "http://localhost:8080";

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

export const analyzeURL = async (url) => {
  try {
    const response = await api.post("/analyze", { url });
    return response.data;
  } catch (error) {
    throw error.response
      ? error.response.data.error
      : "Something went wrong. Please try again.";
  }
};
