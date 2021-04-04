import React from 'react';
import './MyReservationsListItem.css';
import Reservation from './Reservation';

function MyReservationListItem(props) {
    console.log("element is rendering");
    return (
        <div className="container">
            <div className="container-item">
                <p1>Hotel Name: {props.item.hotelInfoPreview.hotelName}</p1>
                <p>City: {props.item.hotelInfoPreview.city}</p>
                <p1>Country: {props.item.hotelInfoPreview.country}</p1>
                <h3>Offers in reservation:</h3>
                {
                    props.item.offerReservations.reservationsInfo.map(item => 
                    (<Reservation key={item.id} item={item}></Reservation>))
                }
            </div>
        </div>
    );
}

export default MyReservationListItem;