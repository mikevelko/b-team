import React, { Component } from 'react';
import './HotelsListItem.css'

function HotelsListItem(props) {
    return (
        <div className="container">
            <div className="container-item-hotel">
                <p>{props.item.hotelName}</p>
                <p>{props.item.city}, {props.item.country}</p>
                <img src="https://whatsanswer.com/wp-content/uploads/2020/04/wawlc-exterior-7823-hor-wide.jpg" width="80%" height="auto"></img>
            </div>


        </div>
    );
}

export default HotelsListItem;