import React, { Component, useState,useEffect } from 'react';
import './LoginPage.css';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { useHistory } from 'react-router-dom';

function LoginPage(props) {
    const history = useHistory();

    const routeChange = () => {
        props.Login();
        let path = `/home`;
        history.push(path);
    }



    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [buttonDisabled, setDisabled] = useState(true);

    useEffect(() => {
        SpellCheck();
        }, [username,password])

    const handleChangeUsername = (event) => {
        setUsername(event.target.value);

    };
    const handleChangePassword = (event) => {
        setPassword(event.target.value);
    };

    function SpellCheck()
    {
        if(username.length>2 && password.length>2) setDisabled(false);
        else setDisabled(true);
    }


    return (
        <div>
            <div className="nav-inputs">
                <ul>
                    <li className="li">
                        <TextField required id="outlined-basic" label="username" variant="outlined" onChange={handleChangeUsername}/>
                    </li>
                    <li className="li">
                        <TextField required id="outlined-basic" label="password" variant="outlined" onChange={handleChangePassword}/>
                    </li>
                    <li className="li-2">
                        <Button variant="contained" color="primary" size="large" onClick={routeChange} disabled={buttonDisabled}>Login</Button>
                    </li>

                </ul>
            </div>
        </div>
    );
}

export default LoginPage;