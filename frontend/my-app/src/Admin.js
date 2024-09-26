import './FilmographieAct.css';
import OmarSy from './OmarSy.png';

function Filmographie() {
    return (
        <div class="container">
        <header class="header">
            <div class="profile">
            </div>
            <div class="intro">
                <h2>Name</h2>
                <h2>Name</h2>
            </div>
            <img src={OmarSy} className="Omar-img" alt="logo" />
        </header>

        <section class="about">
            <h3>Actor Filmography</h3>
            <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Amet vulputate tristique quam felis. Id phasellus dui orci vulputate consequat nulla prain.</p>
        </section>
    </div>
    );
}

export default Filmographie;
