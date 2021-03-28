import React, { Component } from 'react';
import './Nav.css';
import Button from '@material-ui/core/Button';
import { Link, useHistory } from 'react-router-dom';

function Nav(props) {
    const history = useHistory();

    const routeChange = () => {
        props.Logout();
        let path = `/client/login`;
        history.push(path);
    }
    return (
        <div className="profile-nav">
            <div className="main-navbar-left">
                <h2>Bookly client</h2>
            </div>
            <div className="main-navbar-right">
                {!props.isUserAuthenticated ? "" :
                    <Button variant="contained" color="secondary" size="large" onClick={routeChange}>Logout</Button>
                }
            </div>
        </div>
    );
}

export default Nav;