import { Button, TextField } from '@material-ui/core';
import React, { useState } from 'react';
import { HOTEL_TOKEN_NAME, TryLogIn } from './FetchUtils';
import './LogIn.css'

import { useHistory } from "react-router-dom";



function LogIn() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const handleChangeUsername = (event) => {
    setUsername(event.target.value);

  };
  const handleChangePassword = (event) => {
      setPassword(event.target.value);
  };
  const history = useHistory();

  function OnClickLogInButton(){
    TryLogIn(username, password).then(function(response){
      if(response != ""){

        localStorage.setItem(HOTEL_TOKEN_NAME, response);
        history.push('/');
      }
    });
  }
  return (
    <div className="nav-inputs">
      <ul>
          <li className="li">
              <TextField required id="outlined-basic" label="username" variant="outlined" onChange={handleChangeUsername}/>
          </li>
          <li className="li">
              <TextField required id="outlined-basic2" label="password" variant="outlined" onChange={handleChangePassword}/>
          </li>
          <li className="li-2">
              <Button variant="contained" color="primary" size="large" type='submit' onClick={() =>{OnClickLogInButton()}}>Login</Button>
          </li>
      </ul>
    </div>
  );
}

export default LogIn;