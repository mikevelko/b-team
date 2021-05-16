import React, { Component } from 'react';
import './HotelsListItem.css'

function HotelsListItem(props) {
    return (
        <div className="container">
            <div className="container-item-hotel">
                <p>Hotel Name: {props.item.hotelName}</p>
                <p>City: {props.item.city}</p>
                <p>Country: {props.item.country}</p>
            </div>


        </div>
    );
}

export default HotelsListItem;