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

