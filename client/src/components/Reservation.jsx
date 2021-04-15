import React from 'react';
import './Reservation.css';

function Reservation(props) {
    return (
        <div className="container-3">
            <p1>Time from: {props.item.from}</p1>
            <p>Time to: {props.item.to}</p>
            <p1>{props.item.numberOfChildren} children and {props.item.numberOfAdults} adults</p1>
        </div>
    );
}

export default Reservation;