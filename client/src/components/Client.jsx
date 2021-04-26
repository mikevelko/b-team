import React, { Component, useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import './Client.css';
import TextField from '@material-ui/core/TextField';
import axios from 'axios';

function Client() {


    const [name, setName] = useState("");
    const [surname, setSurname] = useState("");
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");


    //first entry to this page by useffect
    useEffect(() => {
        fetchItems();
    }, []);

    const fetchItems = () => {
        axios.get('/api-client/client', { headers: { 'accept': 'application/json', 'x-session-token': window.localStorage.getItem("token") } })
            .then(response => {
                setName(response.data.name);
                setSurname(response.data.surname);
                setUsername(response.data.username);
                setEmail(response.data.email);
                console.log(response.data.name);
            })
            .catch(error => {
                console.error('There was an error!', error);
            });


    }

    const [IsEditing, setIsEditing] = useState(false);

    //valid information, updated after correct data in form 
    function isValidEmailAddress(address) {
        return !!address.match(/^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/);
    }
    function isValidUsername(username) {
        return username.length > 2;
    }

    const handleChangeUsername = (event) => {
        setUsername(event.target.value);
    };

    //const [formEmail, setFormEmail] = useState("");
    const handleChangeEmail = (event) => {
        setEmail(event.target.value);
    };


    const UpdateInformation = () => {
        const body = {
            username: username,
            email: email
        };
        const headers = {
            'accept': 'application/json',
            'x-session-token': window.localStorage.getItem("token"),
            'Content-Type': 'application/json'
        };
        axios.patch('/api-client/client', body, { headers })
            .then(response => {
                console.log(response);
                setIsEditing(false);
                fetchItems();
            })
            .catch(error => {
                console.error('There was an error!', error.response);
            });
        
    }

    function EditInformation() {
        setIsEditing(true);
    }

    return (
        <div>
            <div className="nav-buttons">
                {IsEditing ? <Button variant="contained" color="primary" size="large" onClick={UpdateInformation} disabled={!isValidEmailAddress(email) || !isValidUsername(username)}>Save profile</Button> :
                    <Button variant="contained" color="primary" size="large" onClick={EditInformation}>Edit profile</Button>}
            </div>
            <div className="nav-inputs">
                <ul>
                    <li className="li">
                        <TextField required id="outlined-basic" label="Name" variant="outlined" disabled={true} value={name} />
                    </li>
                    <li className="li">
                        <TextField required id="outlined-basic" label="Surname" variant="outlined" disabled={true} value={surname} />
                    </li>
                    <li className="li">
                        <TextField required id="outlined-basic" label="username" variant="outlined" disabled={!IsEditing} value={username} onChange={handleChangeUsername}
                            error={!isValidUsername(username)} helperText={!isValidUsername(username) ? 'username is too short' : ' '} />
                    </li>
                    <li className="li">
                        <TextField required id="outlined-basic" label="email" variant="outlined" disabled={!IsEditing} value={email} onChange={handleChangeEmail}
                            error={!isValidEmailAddress(email)} helperText={!isValidEmailAddress(email) ? 'email is not valid' : ' '} />
                    </li>
                </ul>
            </div>
        </div>
    );

}

export default Client;