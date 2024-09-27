import axios from 'axios';

const BASE_URL = "http://localhost:8080";

export const getMadeMovies = async () => {
    try {
        const response = await axios.get(`${BASE_URL}/made-movies`);
        return response.data;
    } catch (error) {
        console.error("Error fetching made movies", error);
        throw error;
    }
};

export const getActedMovies = async () => {
    try {
        const response = await axios.get(`${BASE_URL}/acted-movies`);
        return response.data;
    } catch (error) {
        console.error("Error fetching acted movies", error);
        throw error;
    }
};

export const getRecordedMovies = async () => {
    try {
        const response = await axios.get(`${BASE_URL}/recorded-movies`);
        return response.data;
    } catch (error) {
        console.error("Error fetching recorded movies", error);
        throw error;
    }
};

export const addMovie = async (movie) => {
    try {
        await axios.post(`${BASE_URL}/add-movie`, movie);
        console.log("Movie added successfully");
    } catch (error) {
        console.error("Error adding movie", error);
        throw error;
    }
};

export const updateMovie = async (movie) => {
    try {
        await axios.put(`${BASE_URL}/update-movie/${movie._id}`, movie);
        console.log("Movie updated successfully");
    } catch (error) {
        console.error("Error updating movie", error);
        throw error;
    }
};

export const deleteMovie = async (id) => {
    try {
        await axios.delete(`${BASE_URL}/delete-movie/${id}`);
        console.log("Movie deleted successfully");
    } catch (error) {
        console.error("Error deleting movie", error);
        throw error;
    }
};

export const login = async (user, pwd) => {
    try {
        await axios.post(`${BASE_URL}/login`, user, pwd);
        console.log("Login Successfully");
    } catch (error) {
        console.error("Error login", error);
        throw error;
    }
}