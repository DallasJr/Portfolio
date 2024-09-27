import './FilmographieAct.css';
import { useLocation, useNavigate } from 'react-router-dom';
import OmarSy from './OmarSy.png';
import React, { useEffect, useState } from 'react';
import { getActedMovies } from './Axios';
import { Button } from "@mui/material";

function FilmographieContent() {
    const navigate = useNavigate();
    const [madeMovies, setMadeMovies] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchMovies = async () => {
            setLoading(true);
            try {
                const data = await getActedMovies();
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
            <form>
                <div className="menu-container">
                    <ul className="menu">
                        <li>
                            <Button variant="outlined" onClick={() => navigate('/')} className="App-link">
                                Accueil
                            </Button>
                        </li>
                        <li>
                            <Button variant="outlined" onClick={() => navigate('/FilmographieRea')} className="App-link">
                                Voir la Filmographie (Réalisateur)
                            </Button>
                        </li>
                        <li>
                            <Button variant="outlined" onClick={() => navigate('/Admin')} className="App-link">
                                Admin
                            </Button>
                        </li>
                    </ul>
                </div>
            </form>
            <header className="header">
                <div className="profile"></div>
                <div className="intro">
                    <h2>Omar</h2>
                    <h2>Sy</h2>
                </div>
                <img src={OmarSy} className="Omar-img" alt="Omar Sy" />
            </header>

            <section className="about">
                <h3>Actor Filmography</h3>
                {loading && <p>Loading movies...</p>}
                {error && <p>{error}</p>}
                {!loading && !error && madeMovies.length > 0 ? (
                    <ul>
                        {madeMovies.map(movie => (
                            <li key={movie.id}>
                                <h2>{movie.title}</h2>
                                <p>{movie.description}</p>
                                <p>Released: {movie.releaseDate}</p>
                                <img 
                                    src={movie.image || 'default_image_url.png'} 
                                    alt={movie.title} 
                                    onError={(e) => e.target.src = 'default_image_url.png'}
                                />
                            </li>
                        ))}
                    </ul>
                ) : (
                    !loading && <p>No movies found.</p>
                )}
                <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Amet vulputate tristique quam felis. Id phasellus dui orci vulputate consequat nulla prain.</p>
            </section>
        </div>
    );
}

export default function Filmographie() {
    const location = useLocation();

    return (
        <div>
            {location.pathname !== '/FilmographieRea' && location.pathname !== '/Admin' && <FilmographieContent />}
        </div>
    );
}
