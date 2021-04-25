import React, { Component, useEffect, useState } from 'react';
import Button from '@material-ui/core/Button';
import './Hotels.css';
import TextField from '@material-ui/core/TextField';
import HotelsListItem from './HotelsListItem';
import axios from 'axios'
import { Link, useHistory, BrowserRouter as Router, Route } from 'react-router-dom';;

function Hotels() {
    useEffect(() => {
        fetchItems();
    }, []);

    const [items, setItems] = useState([]);
    const [city, setCity] = useState("");
    const [country, setCountry] = useState("");
    const [hotelName, setHotelName] = useState("");

    const handleChangeCity = (event) => {
        setCity(event.target.value);
    };
    const handleChangeCountry = (event) => {
        setCountry(event.target.value);
    };
    const handleChangeHotelName = (event) => {
        setHotelName(event.target.value);
    };

    const fetchItems = () => {

        const body = {
            city: city,
            country: country,
            hotelName: hotelName
        };
        const headers = {
            'accept': 'application/json',
            'x-session-token': window.localStorage.getItem("token")
        };
        axios.get('/api-client/hotels', { headers: { 'accept': 'application/json', 'x-session-token': window.localStorage.getItem("token") } })
            .then(response => {
                setItems(response.data);
                console.log(response);
            })
            .catch(error => {
                console.error('There was an error!', error);
            });
        console.log(items);
    }

    const Search = () => {
        let url = "";
        if(city !== "" || country !== "" || hotelName !== "") url += "?";
        if(country !== "") url += "country=" + country;
        if(url.length > 2 && city !== "") url += "&city=" + city; 
        else if(city !== "") url += "city=" + city;
        if(url.length > 2 && hotelName !== "") url += "&hotelName=" + hotelName; 
        else if(hotelName !== "") url += "hotelName=" + hotelName;

        const GET_URL = '/api-client/hotels' + url;
        console.log(GET_URL);
        const headers = {
            'accept': 'application/json',
            'x-session-token': window.localStorage.getItem("token")
        };
        axios.get(GET_URL , { headers: { 'accept': 'application/json', 'x-session-token': window.localStorage.getItem("token") } })
            .then(response => {
                setItems(response.data);
                console.log(response);
            })
            .catch(error => {
                console.error('There was an error!', error);
            });
        console.log(items);
    }

    return (
        <div>
            <div className="filters">

                <ul className="ul-filters">
                    <li className="ul-li-filters">
                        <TextField id="outlined-basic" label="Hotel name" variant="outlined" size="small" value={hotelName} onChange={handleChangeHotelName}/>
                    </li>
                    <li className="ul-li-filters">
                        <TextField id="outlined-basic" label="City" variant="outlined" size="small" value={city} onChange={handleChangeCity}/>
                    </li>
                    <li className="ul-li-filters">
                        <TextField id="outlined-basic" label="Country" variant="outlined" size="small" value={country} onChange={handleChangeCountry}/>
                    </li>
                    <li className="ul-li-filters">
                        <Button variant="contained" color="primary" size="large" onClick={Search}>Search</Button>
                    </li>
                </ul>
            </div>
            <div>
                {
                    items.map(item =>
                        (<Link key={item.hotelId} to={`/hotels/${item.hotelId}`} className="hotel-link"><HotelsListItem item={item}></HotelsListItem></Link>))
                }
            </div>
        </div>
    );
}

export default Hotels;