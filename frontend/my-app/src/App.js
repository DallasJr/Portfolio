import React from 'react';
import {BrowserRouter as Router, Route, Routes, useLocation, useNavigate} from 'react-router-dom';
import './App.css';
import FilmographieAct from './FilmographieAct';
import FilmographieRea from './FilmographieRea';
import Admin from './Admin';
import {Button} from "@mui/material";

function AppContent() {
    const navigate = useNavigate();

    return (
        <div className="App">
            <form>
                <div className="menu-container">
                    <ul className="menu">
                        <li>
                            <Button variant="outlined" onClick={() => navigate('/FilmographieAct')}
                                    className="App-link">
                                Voir la Filmographie (Acteur)
                            </Button>
                        </li>
                        <li>
                            <Button variant="outlined" onClick={() => navigate('/FilmographieRea')}
                                    className="App-link">
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
            <div className="Description">
                <div className="">
                    <img src="/omarSY.png" className="App-logo" alt="omar si"/>
                    <br/>
                    <h1>OMAR SY</h1>
                </div>
                <p>
                    Who is Omar Sy?
                    <br></br>
                    OMAR SY is a distinguished French actor, film producer, director, screenwriter, and comedian.
                    He was born on January 20, 1988, in Trappes, France. Omar Sy is widely recognized in France for
                    his comedic sketches with Fred Testot on the television show Service après-vente des émissions on
                    Canal+
                    (2005–2012).
                    <br></br>
                    In 2012, he was awarded the César for Best Actor for his performance in the film Intouchables.
                    This success notably opened the doors to Hollywood, where he went on to star in several American
                    blockbusters.
                    Despite his remarkable career, Omar Sy has been subjected to harassment due to his name. He has been
                    the
                    target of wordplay
                    and photo montages, which are in poor taste and far from humorous.
                </p>
            </div>
        </div>
    );
}

function App() {
    const location = useLocation();

    return (
        <div>
            {location.pathname !== '/FilmographieAct' && location.pathname !== '/FilmographieRea' && location.pathname !== '/Admin' &&
                <AppContent/>}

            <Routes>
                <Route path="/FilmographieAct" element={<FilmographieAct/>}/>
                <Route path="/FilmographieRea" element={<FilmographieRea/>}/>
                <Route path="/Admin" element={<Admin/>}/>
            </Routes>
        </div>
    );
}

export default function Main() {
    return (
        <Router>
            <App/>
        </Router>
    );
}
