import './FilmographieAct.css';
import OmarSy from './OmarSy.png';
import React, { useEffect, useState } from 'react';
import { getMadeMovies } from './Axios';

function Filmographie() {
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
            <header className="header">
                <div className="profile">
                </div>
                <div className="intro">
                    <h2>Omar</h2>
                    <h2>Sy</h2>
                </div>
                <img src={OmarSy} className="Omar-img" alt="logo" />
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

export default Filmographie;
