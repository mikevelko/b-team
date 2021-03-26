import React, { Component } from 'react';
import './HotelsListItem.css'

function HotelsListItem(props) {
    return (
        <div className="container">
            <div className="container-item">
                <p1>Hotel Name: {props.item.HotelName}</p1>
                <p>City: {props.item.city}</p>
                <p1>Country: {props.item.country}</p1>
            </div>


        </div>
    );
}

export default HotelsListItem;