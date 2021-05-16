import React, { Component, useState, useEffect } from 'react';
import './LoginPage.css';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { Redirect, useHistory } from 'react-router-dom';
import axios from 'axios';

function LoginPage(props) {
    const history = useHistory();

    const routeChange = () => {

        const body = {
            login: username,
            password: password
        };
        const headers = {
            'accept': 'application/json',
            'Content-Type': 'application/json'
        };
        //const response = await axios.post('http://localhost:8080/api-client/client/login', body, {headers});
        //console.log(response==null);
        //props.Login(response.data);


        axios.post('/api-client/client/login', body, { headers })
            .then(response => {
                props.Login(response.data);
                let path = `/home`;
                history.push(path);
            })
            .catch(error => {
                console.error('There was an error!', error.response);
                setErrorMessage("Bad credentials");
            });


    }



    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [errorMessage, setErrorMessage] = useState("");
    const [buttonDisabled, setDisabled] = useState(true);

    useEffect(() => {
        SpellCheck();
        if (props.token) {
            let path = `/home`;
            history.push(path);
        }
    }, [username, password, props.isUserAuthenticated])

    const handleChangeUsername = (event) => {
        setUsername(event.target.value);

    };
    const handleChangePassword = (event) => {
        setPassword(event.target.value);
    };

    function SpellCheck() {
        if (username.length > 2 && password.length > 2) setDisabled(false);
        else setDisabled(true);
    }

    return (

        <div>
            <div className="nav-inputs">
                <ul>
                    <li className="li">
                        <TextField required id="outlined-basic" label="username" variant="outlined" onChange={handleChangeUsername} error={errorMessage !== ""}
                            helperText={errorMessage !== "" ? 'wrong credentials' : ' '} />
                    </li>
                    <li className="li">
                        <TextField required id="outlined-basic" label="password" variant="outlined" onChange={handleChangePassword} error={errorMessage !== ""}
                            helperText={errorMessage !== "" ? 'wrong credentials' : ' '} />
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