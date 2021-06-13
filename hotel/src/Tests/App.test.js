import '@testing-library/jest-dom';
import { render, screen } from '@testing-library/react';
import React from 'react';
import { BrowserRouter as Router, Route, MemoryRouter  } from 'react-router';
import App from '../MainComponents/App';
import CreateOffer from '../Offers/CreateOffer';
import HotelInfo from '../HotelInfo/HotelInfo';
import HotelInfoEdit from '../HotelInfo/HotelInfoEdit';
import Offers from '../Offers/Offers';
import Reservations from '../Reservations/Reservations';
import Rooms from '../Rooms/Rooms';
import StartPage from '../MainComponents/StartPage';

test('renders learn react link', async () => {
  render(<App />);
  const linkElement = screen.getByText("Hotel Info");
  expect(linkElement).toBeInTheDocument();
});

test('StartPage', async() =>{
  render(<MemoryRouter><StartPage/></MemoryRouter>);
  const linkElement = screen.getByText("Offers");
  expect(linkElement).toBeInTheDocument();
});

test('HotelInfoEdit',async() =>{
  render(<MemoryRouter><HotelInfoEdit/></MemoryRouter>);
  const linkElement = screen.getByText("Save changes");
  expect(linkElement).toBeInTheDocument();
});

test('HotelInfo',async() =>{
  render(<MemoryRouter><HotelInfo/></MemoryRouter>);
  const linkElement = screen.getByText("Edit Hotel Info");
  expect(linkElement).toBeInTheDocument();
});

test('Offers',async() =>{
  render(<MemoryRouter><Offers/></MemoryRouter>);
  const linkElement = screen.getByText("Add new offer");
  expect(linkElement).toBeInTheDocument();
});

test('CreateOffer',async() =>{
  render(<MemoryRouter><CreateOffer/></MemoryRouter>);
  const linkElement = screen.getByText("Create offer");
  expect(linkElement).toBeInTheDocument();
});

test('Reservations',async() =>{
  render(<MemoryRouter><Reservations/></MemoryRouter>);
  const linkElement = screen.getByText("Reservations");
  expect(linkElement).toBeInTheDocument();
});

test('Rooms',async() =>{
  render(<MemoryRouter><Rooms/></MemoryRouter>);
  const linkElement = screen.getByText("Rooms");
  expect(linkElement).toBeInTheDocument();
});
