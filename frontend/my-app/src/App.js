import React from 'react';
import { BrowserRouter as Router, Route, Routes, useNavigate, useLocation } from 'react-router-dom';
import './App.css';
import FilmographieAct from './FilmographieAct';
import FilmographieRea from './FilmographieRea';
import Admin from './Admin';

function AppContent() {
  const navigate = useNavigate();

  return (
    <div className="App">
      <header className="App-header">
        <form>
          <div className="menu-container">
            <ul className="menu">
              <li>
                <button onClick={() => navigate('/FilmographieAct')} className="App-link">
                  Voir la Filmographie (Acteur)
                </button>
              </li>
              <li>
                <button onClick={() => navigate('/FilmographieRea')} className="App-link">
                  Voir la Filmographie (Réalisateur)
                </button>
              </li>
              <li>
                <button onClick={() => navigate('/Admin')} className="App-link">
                  Admin
                </button>
              </li>
            </ul>
          </div>
        </form>
      </header>
      <div className="body">
        <img src="/omarSY.png" className="App-logo" alt="omar si" />
        <br />
        <h1>OMAR SY</h1>
      </div>
    </div>
  );
}

function App() {
  const location = useLocation();

  return (
    <div>
      {location.pathname !== '/FilmographieAct' && location.pathname !== '/FilmographieRea' && location.pathname !== '/Admin' && <AppContent />}
      
      <Routes>
        <Route path="/FilmographieAct" element={<FilmographieAct />} />
        <Route path="/FilmographieRea" element={<FilmographieRea />} />
        <Route path="/Admin" element={<Admin />} />
      </Routes>
    </div>
  );
}

export default function Main() {
  return (
    <Router>
      <App />
    </Router>
  );
}
