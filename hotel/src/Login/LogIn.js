import { Button, Checkbox, FormControlLabel, TextField } from '@material-ui/core';
import React, { useState } from 'react';
import { HOTEL_TOKEN_NAME, CheckToken, TryLogIn } from '../Utils/FetchUtils';
import './LogIn.css'

import { useHistory } from "react-router-dom";



function LogIn() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [token, setToken] = useState("");

  const [loginUsingToken,setLoginUsingToken] = useState(false);
  const handleChangeUsername = (event) => {
    setUsername(event.target.value);

  };
  const handleChangePassword = (event) => {
      setPassword(event.target.value);
  };
  const history = useHistory();

  function OnClickLogInButton(){
    if(loginUsingToken){
      CheckToken(token).then(function (response) {
        console.log(response)
        if(response !== undefined){
          localStorage.setItem(HOTEL_TOKEN_NAME, token);
          history.push('/');
        }
      })
      
    }else{
      TryLogIn(username, password).then(function(response){
        if(response != ""){
  
          localStorage.setItem(HOTEL_TOKEN_NAME, response);
          history.push('/');
        }
      });
    }
  }
  return (
    <div className="nav-inputs">
      <ul>
          {!loginUsingToken ?
            <>
              <li className="li">
                <TextField required id="outlined-basic" label="username" variant="outlined" onChange={handleChangeUsername}/>
              </li>
              <li className="li">
                  <TextField required id="outlined-basic2" label="password" variant="outlined" onChange={handleChangePassword}/>
              </li>

            </>
          :
            <li className="li">
              <TextField required id="outlined-basic2" label="token" variant="outlined" value={token} onChange={(e)=>setToken(e.target.value)}/>
            </li>
          }

          <li className="li"> 
            <FormControlLabel control={<Checkbox color="primary" checked={loginUsingToken} onClick={(e)=> {setLoginUsingToken(e.target.checked)}}/>} label="Login using token"/>
          </li>
          <li className="li-2">
              <Button variant="contained" color="primary" size="large" type='submit' onClick={() =>{OnClickLogInButton()}}>Login</Button>
          </li>

      </ul>
    </div>
  );
}

export default LogIn;