import { Button, TextField } from '@material-ui/core';
import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import './LogIn.css'
function LogIn() {

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleChangeUsername = (event) => {
    setUsername(event.target.value);

  };
  const handleChangePassword = (event) => {
      setPassword(event.target.value);
  };
  return (
    <form className="nav-inputs" method="post">
      <ul>
          <li className="li">
              <TextField required id="outlined-basic" label="username" variant="outlined" onChange={handleChangeUsername}/>
          </li>
          <li className="li">
              <TextField required id="outlined-basic" label="password" variant="outlined" onChange={handleChangePassword}/>
          </li>
          <li className="li-2">
              <Button variant="contained" color="primary" size="large" type='submit' onClick={() =>{}}>Login</Button>
          </li>

      </ul>
    </form>
  );
}

export default LogIn;