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
                        <h3>Hello, username</h3>
                    </div>
                    <div className="profile-navbar-right">
                        <Button variant="contained" color="primary" size="large" component={Link} to="/client">My profile</Button>
                    </div>
                </div>
                <div>
                    <ul className="ul-profile">
                        <li className="ul-li-profile">
                            <Route render={({ history }) => (
                                <button className="btn" onClick={() => { history.push('/hotels') }}>
                                    Hotels
                                </button>
                            )} />
                        </li>
                        <li className="ul-li-profile">
                            <Route render={({ history }) => (
                                <button className="btn" onClick={() => { history.push('/reservations') }}>
                                    My reservations
                                </button>
                            )} />
                        </li>
                        <li className="ul-li-profile">
                            <Route render={({ history }) => (
                                <button className="btn" onClick={() => { history.push('/reviews') }}>
                                    My reviews
                                </button>
                            )} />
                        </li>
                    </ul>
                </div>
            </div>
        );
    }
}

export default MainPage;