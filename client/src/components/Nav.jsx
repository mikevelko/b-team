import React, { Component } from 'react';
import './Nav.css';
import Button from '@material-ui/core/Button';
import { Link, useHistory } from 'react-router-dom';
import logo from './logo.png';

function Nav(props) {
    const history = useHistory();

    const routeChange = () => {
        props.Logout();
        let path = `/login`;
        history.push(path);
    }
    return (
        <div className="profile-nav">
            <div className="main-navbar-left">
                <Link to="/home" className="link">
                    <h2>
                        <img src={logo} className="photo" alt="Bookly client logo" />
                    </h2>
                </Link>
            </div>
            <div className="main-navbar-right">
                {!props.token ? "" :
                    <Button variant="contained" color="secondary" size="large" onClick={routeChange}>Logout</Button>
                }
            </div>
        </div>
    );
}

export default Nav;