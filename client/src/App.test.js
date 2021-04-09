import React from 'react';
import { render, screen } from '@testing-library/react';
import App from './App';
import ReactDOM from 'react-dom';
import Button from '@material-ui/core/Button';
import "@testing-library/jest-dom/extend-expect";
import MainPage from './components/MainPage';
import PrivateRoute from './components/PrivateRoute';
import { BrowserRouter as Router } from 'react-router-dom';
import Hotels from './components/Hotels';
import MyReservations from './components/MyReservations';
import LoginPage from './components/LoginPage';


test('bookly client label logo test', () => {
  render(<App />);
  expect(screen.getByText("Bookly client")).toBeInTheDocument();
});

test('button my profile test', () => {
  render(<Button variant="contained" color="primary" size="large" data-testid="my-profile">My profile</Button>);
  expect(screen.getByText("My profile")).toBeInTheDocument();
});


test('login component test', () => {
  render(<App />);
  expect(screen.getByText("Login")).toBeInTheDocument();
});

test('login component test', () => {
  render(<App />);
  expect(screen.getByText("username")).toBeInTheDocument();
});

test('login component test', () => {
  render(<App />);
  expect(screen.getByText("password")).toBeInTheDocument();
});

test('main page test', () => {
  render(<Router><PrivateRoute authed={true} exact path='/home' component={MainPage}/></Router>);
  expect(screen.getByText("Hotels")).toBeInTheDocument();
});

test('main page test', () => {
  render(<Router><PrivateRoute authed={true} exact path='/home' component={MainPage}/></Router>);
  expect(screen.getByText("My profile")).toBeInTheDocument();
});

test('hotels view test', () => {
  render(<Router><PrivateRoute authed={true} exact path='/hotels' component={Hotels}/></Router>);
  expect(screen.getByText("Search")).toBeInTheDocument();
});

test('my reservtaions view test', () => {
  render(<Router><PrivateRoute authed={true} exact path='/client/reservations' component={MyReservations} /></Router>);
  expect(screen.getByText("My reservations:")).toBeInTheDocument();
});

