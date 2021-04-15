import React, { Component, useEffect, useState } from 'react';
import Button from '@material-ui/core/Button';
import './Hotels.css';
import TextField from '@material-ui/core/TextField';
import HotelsListItem from './HotelsListItem';


function Hotels() {
    useEffect(() => {
        fetchItems();
    }, []);

    const [items, setItems] = useState([]);

    const data = [
        { 'city': 'Warsaw', 'country': 'Poland', 'HotelName': 'Grand Hotel 1' },
        { city: 'Warsaw', 'country': 'Poland', 'HotelName': 'Grand Hotel 2' },
        { city: 'Warsaw', 'country': 'Poland', 'HotelName': 'Grand Hotel 3' },
        { city: 'Warsaw', 'country': 'Poland', 'HotelName': 'Grand Hotel 4' },
        { city: 'Warsaw', 'country': 'Poland', 'HotelName': 'Grand Hotel 5' },
        { city: 'Minsk', 'country': 'Belarus', 'HotelName': 'Grand Hotel 6' }];

    const fetchItems = async () => {

    }

    return (
        <div>
            <div className="filters">

                <ul className="ul-filters">
                    <li className="ul-li-filters">
                        <TextField id="outlined-basic" label="Hotel name" variant="outlined" size="small" />
                    </li>
                    <li className="ul-li-filters">
                        <TextField id="outlined-basic" label="City" variant="outlined" size="small" />
                    </li>
                    <li className="ul-li-filters">
                        <TextField id="outlined-basic" label="Country" variant="outlined" size="small" />
                    </li>
                    <li className="ul-li-filters">
                        <Button variant="contained" color="primary" size="large">Search</Button>
                    </li>
                </ul>
            </div>
            <div>
                {
                    data.map(item =>
                    (<HotelsListItem key={item.id} item={item}></HotelsListItem>)) 
                }
            </div>
        </div>
    );
}

export default Hotels;