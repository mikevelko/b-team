import './App.css';
import React, { Component, useState } from 'react';
import { BrowserRouter as Router, Route, Switch, Redirect } from 'react-router-dom';
import MainPage from './components/MainPage';
import Nav from './components/Nav';
import Hotels from './components/Hotels';
import Client from './components/Client';
import PrivateRoute from './components/PrivateRoute'
import LoginPage from './components/LoginPage';
import MyReservations from './components/MyReservations';

class App extends Component {
  constructor() {
    super();
    this.Logout = this.Logout.bind(this);
    this.Login = this.Login.bind(this);
    this.state = {
      isUserAuthenticated: false
    };
  }

  Logout() {
    this.setState({
      isUserAuthenticated: false
    });
  }

  Login() {
    this.setState({
      isUserAuthenticated: true
    });
  }

  render() {
    return (
      <Router>
        <Nav isUserAuthenticated={this.state.isUserAuthenticated} Logout={this.Logout}/>
        <Switch>
          <Route
            exact
            path="/"
            render={() => {
              return (
                this.isUserAuthenticated ?
                  <Redirect to="/client/login" /> :
                  <Redirect to="/home" />
              )
            }}
          />
          <PrivateRoute authed={this.state.isUserAuthenticated} exact path='/home' component={MainPage} />
          <PrivateRoute authed={this.state.isUserAuthenticated} exact path='/hotels' component={Hotels} />
          <PrivateRoute authed={this.state.isUserAuthenticated} exact path='/client' component={Client} />
          <PrivateRoute authed={this.state.isUserAuthenticated} exact path='/client/reservations' component={MyReservations} />
          <Route exact path="/client/login" component={() => <LoginPage Login={this.Login}/>}/>
        </Switch>
      </Router>
    );
  }
}

export default App;
