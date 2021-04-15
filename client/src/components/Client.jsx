import React, { Component, useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import './Client.css';
import TextField from '@material-ui/core/TextField';
import axios from 'axios';

function Client() {

    //first entry to this page by useffect 

    useEffect(() => {
       fetchItems();
     }, []);

    const fetchItems = () => {
        const headers = {
            'accept': 'application/json',
            'x-session-token': '{id: 1, createdAt: "2021-04-15T18:02:17Z"}'
        };

        axios.get('http://localhost:8080/api-client/client', { headers: { 'accept': 'application/json', 'x-session-token':  '{id: 1, createdAt: "2021-04-15T18:02:17Z"}'} })
            .then(response => {
                console.log(response);
            })
            .catch(error => {
                console.error('There was an error!', error);
            });


    }

    const [IsEditing, setIsEditing] = useState(false);

    //valid information, updated after correct data in form 
    const [name, setName] = useState("name");
    const [surname, setSurname] = useState("");
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");


    //information from the form
    const [formName, setFormName] = useState("");
    const handleChangeName = (event) => {
        setFormName(event.target.value);
    };

    const [formSurname, setFormSurname] = useState("");
    const handleChangeSurname = (event) => {
        setFormSurname(event.target.value);
    };

    const [formUsername, setFormUsername] = useState("");
    const handleChangeUsername = (event) => {
        setFormUsername(event.target.value);
    };

    const [formEmail, setFormEmail] = useState("");
    const handleChangeEmail = (event) => {
        setFormEmail(event.target.value);
    };


    function UpdateInformation() {
        //post method here with validation of response 
        setIsEditing(false);
    }
    function EditInformation() {
        setIsEditing(true);
    }



    //for TextField there is an option to show error 
    return (
        <div>
            <div className="nav-buttons">
                {IsEditing ? <Button variant="contained" color="primary" size="large" onClick={UpdateInformation}>Save profile</Button> :
                    <Button variant="contained" color="primary" size="large" onClick={EditInformation}>Edit profile</Button>}
            </div>
            <div className="nav-inputs">
                <ul>
                    <li className="li">
                        <TextField required id="outlined-basic" label="Name" variant="outlined" disabled={!IsEditing} defaultValue={name} onChange={handleChangeName} />
                    </li>
                    <li className="li">
                        <TextField required id="outlined-basic" label="Surname" variant="outlined" disabled={!IsEditing} defaultValue={surname} onChange={handleChangeSurname} />
                    </li>
                    <li className="li">
                        <TextField required id="outlined-basic" label="username" variant="outlined" disabled={!IsEditing} defaultValue={username} onChange={handleChangeUsername} />
                    </li>
                    <li className="li">
                        <TextField required id="outlined-basic" label="email" variant="outlined" disabled={!IsEditing} defaultValue={email} onChange={handleChangeEmail} />
                    </li>
                </ul>
            </div>
        </div>
    );

}

export default Client;