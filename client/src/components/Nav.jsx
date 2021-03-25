import React, { Component } from 'react';
import './Nav.css';
import Button from '@material-ui/core/Button';

class Nav extends Component {
    render() {
        return (
            <div className="profile-nav">
                <div className="main-navbar-left">
                    <h2>Bookly client</h2>
                </div>
                <div className="main-navbar-right">
                    <Button variant="contained" color="secondary" size="large">Logout</Button>
                </div>
            </div>
        );
    }
}

export default Nav;