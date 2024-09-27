import OmarSy from './OmarSy.png';
import {Button, Input} from "@mui/material";
import './Admin.css';
import React from 'react';



function Filmographie() {
    return (
        <section className="section">
            <div className="container">
                <header className="header">
                    <div className="profile">
                        <h5>Username</h5><Input type="text" sx={{ input: {color: 'white' }}}/>
                        <h5>Password</h5><Input type="text" sx={{ input: {color: 'white' }}}/>
                        <h5>Confirm password</h5><Input type="text" sx={{ input: {color: 'white' }}}/>
                    </div>

                    <img src={OmarSy} className="Omar-img" alt="logo"/>
                </header>
                <form>
                    <section className="about">
                        <h3>Add made movie</h3>
                        <h5>Title</h5><Input type="text" sx={{ input: {color: 'white' }}}/>
                        <h5>Release date</h5><Input type="text" sx={{ input: {color: 'white' }}}/>
                        <h5>Description</h5><textarea></textarea>
                        <h5>Image</h5><Input type="file" sx={{ input: {color: 'white' }}}/><br/>
                        <Button className='add'>Add</Button>
                    </section>
                    <section className="about">
                        <h3>Add played movie</h3>
                        <h5>Title</h5><Input type="text" sx={{ input: {color: 'white' }}}/>
                        <h5>Release date</h5><Input type="text" sx={{ input: {color: 'white' }}}/>
                        <h5>Description</h5><textarea></textarea>
                        <h5>Image</h5><Input type="file" sx={{ input: {color: 'white' }}}/><br/>
                        <Button className='add'>Add</Button>
                    </section>
                </form>
            </div>
        </section>
    );
}

export default Filmographie;
