import './App.css';
import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import MainPage from './components/MainPage';
import Nav from './components/Nav';
import Hotels from './components/Hotels';
import Client from './components/Client';

function App() {
  return (
    <Router>
      <Nav/>
      <Switch>
        <Route path="/" exact component={MainPage}/>
        <Route path="/client" exact component={Client}/>
        <Route path="/hotels" exact component={Hotels}/>
      </Switch>
    </Router>
  );
}

export default App;
