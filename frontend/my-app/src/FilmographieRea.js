import './FilmographieRea.css';
import { useLocation, useNavigate } from 'react-router-dom';
import OmarSy from './OmarSy.png';
import React, { useEffect, useState } from 'react';
import { getMadeMovies } from './Axios';
import ResponsiveAppBar from './components/ResponsiveAppBar';

function FilmographieContent() {
    const navigate = useNavigate();
    const [madeMovies, setMadeMovies] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchMovies = async () => {
            setLoading(true);
            try {
                const data = await getMadeMovies();
                console.log("Movies from API:", data);
                setMadeMovies(data);
            } catch (err) {
                console.error("Error fetching made movies", err);
                setError("Could not fetch movies. Please try again later.");
            } finally {
                setLoading(false);
            }
        };        
        fetchMovies();
    }, []);

    return (
        <div className="container">
            <ResponsiveAppBar />             
            <header className="header">
                <div className="profile"></div>
                <div className="intro">
                    <h2>Omar</h2>
                    <h2>Sy</h2>
                </div>
                <img src={OmarSy} className="Omar-img" alt="Omar Sy" />
            </header>

            <section className="about">
                <h3>Director Filmography</h3>
                {loading && <p>Loading movies...</p>}
                {error && <p>{error}</p>}
                {!loading && !error && madeMovies.length > 0 ? (
                    <ul>
                        {madeMovies.map(movie => (
                            <li key={movie.id}>
                                <h2>{movie.title}</h2>
                                <p>{movie.description}</p>
                                <p>Released: {movie.releaseDate}</p>
                            </li>
                        ))}
                    </ul>
                ) : (
                    !loading && <p>No movies found.</p>
                )}
            </section>
        </div>
    );
}

export default function Filmographie() {
    const location = useLocation();

    return (
        <div>
            {location.pathname !== '/FilmographieAct' && location.pathname !== '/Admin' && <FilmographieContent />}
        </div>
    );
}
