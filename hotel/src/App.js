
import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import {BrowserRouter as Router, Switch, Route, Link} from 'react-router-dom';
import StartPage from './StartPage';
import HotelInfo from './HotelInfo';
import Offers from './Offers';
import OfferDetails from './OfferDetails';
import Reservations from './Reservations';
import Rooms from './Rooms';
import LogIn from './LogIn';
import HomeIcon from '@material-ui/icons/Home';

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

  return (
    <Router>
    <div className={classes.root}>
      <AppBar position="static">

        <Toolbar>
          <Link to='/' style={{color:'white'}}>
            <HomeIcon style={{marginRight:20, fontSize: 30}}/>
          </Link>
          <Typography variant="h6" className={classes.title}>
            Hello, Hotel Name
          </Typography>
          <div>
            <Button component={Link} to='/HotelInfo' color="inherit" style={{marginRight:20}}>Hotel Info</Button>
            <Button component={Link} to='/LogIn' color="inherit">Log out</Button>
          </div>
        </Toolbar>
      </AppBar>
    </div>
      <Switch>
        <Route path='/' exact component={StartPage}/>
        <Route path='/hotelInfo' component={HotelInfo}/>
        <Route path='/offers' exact component={Offers}/>
        <Route path='/offers/:offerId' component={OfferDetails}/>
        <Route path='/reservations' component={Reservations}/>
        <Route path='/rooms' component={Rooms}/>
        <Route path='/logIn' component={LogIn}/>
      </Switch>
    </Router>
  );
}
