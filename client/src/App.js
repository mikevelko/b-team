import './App.css';
import React, { Component, useEffect, useState } from 'react';
import { BrowserRouter as Router, Route, Switch, Redirect } from 'react-router-dom';
import MainPage from './components/MainPage';
import Nav from './components/Nav';
import Hotels from './components/Hotels';
import Client from './components/Client';
import PrivateRoute from './components/PrivateRoute'
import LoginPage from './components/LoginPage';
import MyReservations from './components/MyReservations';
import Hotel from './components/Hotel';
import Offers from './components/Offers';

function App() {

  
  const [isUserAuthenticated, SetIsUserAuthenticated] = useState(JSON.parse(window.localStorage.getItem("isUserAuthenticated")));
  const [token, setToken] = useState(JSON.parse(window.localStorage.getItem("token")));

  useEffect(() => {
    const token2 =  (window.localStorage.getItem("token"));
    setToken(JSON.parse(token2));
  }, [])


  useEffect(() => {
    window.localStorage.setItem("token", JSON.stringify(token));
  })

  
  function Logout() {
    window.localStorage.setItem("token", null);
    setToken(null);
  }

  function Login(dataToken) {

    setToken(dataToken);
    //response from server here 
  }

  //if(!isUserAuthenticated) { return <div></div> } 
  return (
    <Router basename={process.env.PUBLIC_URL}>
      <Nav token={token} Logout={Logout} />
      <Switch>
        <Route
          exact
          path="/"
          render={() => {
              return (
                token ?
                  <Redirect to="/login"/> :
                  <Redirect to="/home"/>
              )
          }}
        ></Route>
        <PrivateRoute authed={token} exact path='/home' component={MainPage} />
        <PrivateRoute authed={token} exact path='/hotels' component={Hotels} />
        <PrivateRoute authed={token} exact path='/hotels/:hotelId' component={Hotel} />
        <PrivateRoute authed={token} exact path='/hotels/:hotelId/offers' component={Offers} />
        <PrivateRoute authed={token} exact path='/client' component={Client} />
        <PrivateRoute authed={token} exact path='/reservations' component={MyReservations} />
        
        <Route exact path="/login" component={() => <LoginPage Login={Login} token={token}/>} />
      </Switch>
    </Router>
  );

}


export default App;
