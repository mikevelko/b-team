import React, { Component } from 'react';
import './OffersListItem.css'

function OffersListItem(props) {
    return (
        <div className="container">
            <div className="container-item-offer">
                <p>{props.item.offerTitle}</p>
                <p>maxGuests: {props.item.maxGuests}</p>
                <img src={props.item.offerPreviewPicture} width="400" height="auto"></img>
                <p>Cost for child: {props.item.costPerChild}€</p>
                <p>Cost for adult: {props.item.costPerAdult}€</p>
            </div>


        </div>
    );
}

export default OffersListItem;