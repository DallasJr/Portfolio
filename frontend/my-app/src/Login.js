import './Login.css';
import {Button, Input} from "@mui/material";


function Login() {
    return (
        <div className="container">
            <form>
                <div className="form-group">
                    <Input type="text" className="form-control" placeholder="Username" sx={{ input: {color: 'white' }}}/><br/>
                    <Input type="password" className="form-control" placeholder="Password" sx={{ input: {color: 'white' }}}/>
                    <Button variant="contained" className="login">Login</Button>
                </div>
            </form>
        </div>
    )
}


export default Login;