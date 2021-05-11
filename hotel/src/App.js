import React, { useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import {BrowserRouter as Router, Switch, Route, Link, useHistory} from 'react-router-dom';
import StartPage from './StartPage';
import HotelInfo from './HotelInfo';
import Offers from './Offers';
import Reservations from './Reservations';
import Rooms from './Rooms';
import LogIn from './LogIn';
import HomeIcon from '@material-ui/icons/Home';
import CreateOffer from './CreateOffer';
import HotelInfoEdit from './HotelInfoEdit';
import EditOffer from './EditOffer';
import { PrivateRoute } from './PrivateRoute';
import { TryGetHotelInfo } from './FetchUtils';
import OfferDetails from './OfferDetails';
const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  }
}));

export default function App() {
  const classes = useStyles();

  const [hotelName, setHotelName] = useState('');

  useEffect(()=>{
    TryGetHotelInfo().then(function (response) {
      if(response!=='') setHotelName(response.hotelName)
    })
  },[])

  function LogOut(){
    localStorage.removeItem("x-hotel-token");
    
  }
  return (
    <Router basename={process.env.PUBLIC_URL}>
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Link to='/' style={{color:'white'}}>
            <HomeIcon style={{marginRight:20, fontSize: 30}}/>
          </Link>
          <Typography variant="h6" className={classes.title}>
            Hello, {hotelName}
          </Typography>
          <div>
            <Button component={Link} to='/HotelInfo' color="inherit" style={{marginRight:20}}>Hotel Info</Button>
            <Button component={Link} to='/LogIn' color="inherit" onClick={() => {LogOut()}}>Log out</Button>
          </div>
        </Toolbar>
      </AppBar>
    </div>
      <Switch>
        <PrivateRoute path='/' exact component={StartPage}/>
        <PrivateRoute path='/hotelInfo' exact component={HotelInfo}/>
        <PrivateRoute path='/hotelInfo/edit' exact component={HotelInfoEdit}/>


        <PrivateRoute path='/offers' exact component={Offers}/>
        <PrivateRoute path='/offers/create' exact component={CreateOffer}/>
        <PrivateRoute path='/offers/:offerId' exact component={OfferDetails}/>
        <PrivateRoute path='/offers/edit/:offerId' exact component={EditOffer}/>


        <PrivateRoute path='/reservations' component={Reservations}/>
        <PrivateRoute path='/rooms' component={Rooms}/>
        <Route path='/login' component={LogIn}/>
      </Switch>
    </Router>
  );
}

