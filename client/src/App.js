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

  const [isUserAuthenticated, SetIsUserAuthenticated] = useState(false);
  const [token, setToken] = useState();


  useEffect(() => {
    const data = (window.localStorage.getItem("isUserAuthenticated"));
    console.log(JSON.parse(data));
    SetIsUserAuthenticated(JSON.parse(data));
    console.log(isUserAuthenticated);
  }, [])

  useEffect(() => {
    window.localStorage.setItem("isUserAuthenticated", JSON.stringify(isUserAuthenticated));
  })

  
  function Logout() {
    SetIsUserAuthenticated(false);
  }

  function Login() {

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
        <PrivateRoute authed={isUserAuthenticated} exact path='/hotels' component={Hotels} />
        <PrivateRoute authed={isUserAuthenticated} exact path='/client' component={Client} />
        <PrivateRoute authed={isUserAuthenticated} exact path='/client/reservations' component={MyReservations} />
        <Route exact path="/login" component={() => <LoginPage Login={Login} isUserAuthenticated={isUserAuthenticated}/>} />
      </Switch>
    </Router>
  );

}


export default App;
