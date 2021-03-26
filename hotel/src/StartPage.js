import React from 'react';
import "./StartPage.css";
import { Button } from '@material-ui/core';
import { Link } from 'react-router-dom';


function StartPage() {


  return (
      <div className='buttonsList'>
          <Button component={Link} to='/Offers' className='generalButton' variant="contained" color='primary'>
              Offers
          </Button>
          <Button component={Link} to='/Rooms' className='generalButton' variant="contained" color='primary'>
              Rooms
          </Button>
          <Button component={Link} to='/Reservations' className='generalButton' variant="contained" color='primary'>
              Reservations
          </Button>
      </div>
  );
}

export default StartPage;