import React, { Component } from 'react';
import './MainPage.css';
import Button from '@material-ui/core/Button';
import { Route, Link } from 'react-router-dom';

class MainPage extends Component {



    render() {
        return (
            <div>
                <div className="profile-nav">
                    <div className="profile-navbar-left">
                    </div>
                    <div className="profile-navbar-right">
                        <Button variant="contained" color="primary" size="large" component={Link} to="/client">My profile</Button>
                    </div>
                </div>
                <div>
                    <ul className="ul-profile">
                    <li className="ul-li-profile">
                            <Route render={({ history }) => (
                                <Button variant="contained" color="primary" style={{fontSize: '42px', maxWidth: '100%', maxHeight: '150px', minWidth: '100%', minHeight: '150px'}}
                                 size="large" component={Link} exact to="/hotels">Hotels</Button>
                            )} />
                        </li>
                        <li className="ul-li-profile">
                            <Route render={({ history }) => (
                                <Button variant="contained" color="primary" style={{fontSize: '42px', maxWidth: '100%', maxHeight: '150px', minWidth: '100%', minHeight: '150px'}}
                                 size="large" component={Link} exact to="/client/reservations">My reservations</Button>
                            )} />
                        </li>
                    </ul>
                </div>
            </div>
        );
    }
}

export default MainPage;