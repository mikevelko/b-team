import React, { Component, useEffect, useState } from 'react';
import Button from '@material-ui/core/Button';
import './Offers.css';
import TextField from '@material-ui/core/TextField';
import OffersListItem from './OffersListItem';
import axios from 'axios'
import { Link, useHistory, BrowserRouter as Router, Route } from 'react-router-dom';;

function Offers(props) {

    let hotelId = props.match.params.hotelId;
    let temp = [{
        "offerID": 5,
        "offerTitle": 'string',
        "offerPreviewPicture": null,
        "maxGuests": 2,
        "costPerChild": 2.0,
        "costPerAdult": 3.4
      }];

    useEffect(() => {
        fetchItems();
    }, []);

    let today = new Date();

    let tomorrow = new Date(today)
    tomorrow.setDate(tomorrow.getDate() + 1)
    let dd = today.getDate();
    let mm = today.getMonth() + 1;
    let yyyy = today.getFullYear();
    if (dd < 10) {
        dd = '0' + dd;
    }
    if (mm < 10) {
        mm = '0' + mm;
    }
    today = yyyy + '-' + mm + '-' + dd;

    dd = tomorrow.getDate();
    mm = tomorrow.getMonth() + 1;
    yyyy = tomorrow.getFullYear();
    if (dd < 10) {
        dd = '0' + dd;
    }
    if (mm < 10) {
        mm = '0' + mm;
    }
    tomorrow = yyyy + '-' + mm + '-' + dd;





    const [items, setItems] = useState([]);
    const [from, setFrom] = useState(today);
    const [to, setTo] = useState(tomorrow);
    const [guests, setGuests] = useState(1);
    const [minCost, setMinCost] = useState();
    const [maxCost, setMaxCost] = useState();


    const handleChangeFrom = (event) => {
        console.log(event.target.value)
        setFrom(event.target.value);
    };
    const handleChangeTo = (event) => {
        setTo(event.target.value);
    };
    const handleChangeGuests = (event) => {
        setGuests(event.target.value);
    };
    const handleChangeMinCost = (event) => {
        if(event.target.value < 0)
        {
            setMinCost(0);
            return;
        }
        setMinCost(event.target.value);
    };
    const handleChangeMaxCost = (event) => {
        if(event.target.value < 0)
        {
            setMaxCost(0);
            return;
        }
        setMaxCost(event.target.value);
    };


    const fetchItems = () => {

    }



    return (
        <div>
            <div className="filters">
                <ul className="ul-filters">
                    <li className="ul-li-filters">
                        <TextField id="outlined-basic" label="Check in (mm/dd/yyyy)" type="date" variant="outlined" size="small" value={from} onChange={handleChangeFrom} />
                    </li>
                    <li className="ul-li-filters">
                        <TextField id="outlined-basic" label="Check out (mm/dd/yyyy)" type="date" variant="outlined" size="small" value={to} onChange={handleChangeTo} />
                    </li>
                    <li className="ul-li-filters">
                        <TextField id="outlined-basic" label="Guests" variant="outlined" type="number" InputProps={{ inputProps: { min: 1, max: 10 } }} size="small" value={guests} onChange={handleChangeGuests} />
                    </li>
                    <li className="ul-li-filters">
                        <TextField id="outlined-basic" label="Min cost" variant="outlined" type="number" InputProps={{ inputProps: { min: 0, max: 1000000 } }} size="small" value={minCost} onChange={handleChangeMinCost} />
                    </li>
                    <li className="ul-li-filters">
                        <TextField id="outlined-basic" label="Max cost" variant="outlined" type="number" InputProps={{ inputProps: { min: 0, max: 1000000 } }} size="small" value={maxCost} onChange={handleChangeMaxCost} />
                    </li>
                    <li className="ul-li-filters">
                        <Button variant="contained" color="primary" size="large">Search</Button>
                    </li>
                </ul>
            </div>
            <div>
                {
                    items.map(item =>
                        (<Link key={item.offerID} to={`/hotels/${hotelId}/offers/${item.offerID}`} className="hotel-link"><OffersListItem item={item}></OffersListItem></Link>))
                }
            </div>
        </div>
    );
}

export default Offers;