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

function App() {

  
  const [isUserAuthenticated, SetIsUserAuthenticated] = useState(JSON.parse(window.localStorage.getItem("isUserAuthenticated")));
  const [token, setToken] = useState(JSON.parse(window.localStorage.getItem("token")));

  useEffect(() => {
    const userAuth =  (window.localStorage.getItem("isUserAuthenticated"));
    const token2 =  (window.localStorage.getItem("token"));
    console.log(JSON.parse(userAuth));
    SetIsUserAuthenticated(JSON.parse(userAuth));
    setToken(JSON.parse(token2));
    console.log(isUserAuthenticated);
  }, [])


  useEffect(() => {
    window.localStorage.setItem("isUserAuthenticated", JSON.stringify(isUserAuthenticated));
    window.localStorage.setItem("token", JSON.stringify(token));
  })

  
  function Logout() {
    window.localStorage.setItem("isUserAuthenticated", null);
    window.localStorage.setItem("token", null);
    SetIsUserAuthenticated(null);
    setToken(null);
  }

  function Login(dataToken) {

    setToken(dataToken);
    //response from server here 
    SetIsUserAuthenticated(true);
  }

  //if(!isUserAuthenticated) { return <div></div> } 
  return (
    <Router basename={process.env.PUBLIC_URL}>
      <Nav isUserAuthenticated={isUserAuthenticated} Logout={Logout} />
      <Switch>
        <Route
          exact
          path="/"
          render={() => {
              return (
                isUserAuthenticated ?
                  <Redirect to="/login"/> :
                  <Redirect to="/home"/>
              )
          }}
        ></Route>
        <PrivateRoute authed={isUserAuthenticated} exact path='/home' component={MainPage} />
        <PrivateRoute authed={token} exact path='/hotels' component={Hotels} />
        <PrivateRoute authed={isUserAuthenticated} exact path='/client' component={Client} />
        <PrivateRoute authed={isUserAuthenticated} exact path='/reservations' component={MyReservations} />
        <Route exact path="/login" component={() => <LoginPage Login={Login} isUserAuthenticated={isUserAuthenticated}/>} />
      </Switch>
    </Router>
  );

}


export default App;
