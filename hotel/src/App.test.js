import { render, screen } from '@testing-library/react';
import React from 'react';
import { BrowserRouter as Router, Route, MemoryRouter  } from 'react-router';
import App from './App';
import CreateOffer from './CreateOffer';
import HotelInfo from './HotelInfo';
import HotelInfoEdit from './HotelInfoEdit';
import OfferDetails from './OfferDetails';
import Offers from './Offers';
import Reservations from './Reservations';
import Rooms from './Rooms';
import StartPage from './StartPage';

test('renders learn react link', () => {
  render(<App />);
  const linkElement = screen.getByText("Hotel Info");
  expect(linkElement).toBeInTheDocument();
});

test('StartPage',() =>{
  render(<MemoryRouter><StartPage/></MemoryRouter>);
  const linkElement = screen.getByText("Offers");
  expect(linkElement).toBeInTheDocument();
});

test('HotelInfoEdit',() =>{
  render(<MemoryRouter><HotelInfoEdit/></MemoryRouter>);
  const linkElement = screen.getByText("Save changes");
  expect(linkElement).toBeInTheDocument();
});

test('HotelInfo',() =>{
  render(<MemoryRouter><HotelInfo/></MemoryRouter>);
  const linkElement = screen.getByText("Edit Hotel Info");
  expect(linkElement).toBeInTheDocument();
});

test('Offers',() =>{
  render(<MemoryRouter><Offers/></MemoryRouter>);
  const linkElement = screen.getByText("Add new offer");
  expect(linkElement).toBeInTheDocument();
});

test('CreateOffer',() =>{
  render(<MemoryRouter><CreateOffer/></MemoryRouter>);
  const linkElement = screen.getByText("Create offer");
  expect(linkElement).toBeInTheDocument();
});

test('OfferDetails',() =>{
  render(<MemoryRouter><OfferDetails match={({params:{offerId:'10'}})}/></MemoryRouter>);
  const linkElement = screen.getByText("Offer Details");
  expect(linkElement).toBeInTheDocument();
});

test('Reservations',() =>{
  render(<MemoryRouter><Reservations/></MemoryRouter>);
  const linkElement = screen.getByText("Reservations");
  expect(linkElement).toBeInTheDocument();
});

test('Rooms',() =>{
  render(<MemoryRouter><Rooms/></MemoryRouter>);
  const linkElement = screen.getByText("Rooms");
  expect(linkElement).toBeInTheDocument();
});
