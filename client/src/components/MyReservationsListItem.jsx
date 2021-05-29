import { Button, Input } from '@material-ui/core';
import React, { useEffect, useState } from 'react';
import './MyReservationsListItem.css';
import Reservation from './Reservation';
import Rating from '@material-ui/lab/Rating';
import axios from 'axios';

function MyReservationListItem(props) {

    const [newReview, setNewReview] = useState("");
    const [editing, setEditing] = useState(false);
    const [buttonText, setButtonText] = useState("Edit review");
    const [newRating, setNewRating] = useState();

    useEffect(() => {
        fetchItems();
    }, []);

    const fetchItems = async () => {
        //get review for offer of client if exists
    }


    function ReviewEdit() {
        if (editing === true) {
            setEditing(false);
            setButtonText("Edit review");
        }
        else {
            setEditing(true);
            setButtonText("Save review");
        }
    }

    const DeleteReview = () => {
        //axios delete
    };

    const UpdateReview = () => {
        //axios put
    };

    const CancelReservation = () => {
        const url = `/api-client/client/reservations/${props.item.reservationInfo.reservationID}`;
        axios.delete(url, { headers: { 'accept': 'application/json', 'x-session-token': window.localStorage.getItem("token") } })
            .then(response => {
                console.log(response.data);
                props.fetchReservations();
            })
            .catch(error => {
                //console.error('There was an error!', error.response);
            });
    }


    let today = new Date();
    let CurrentDate = new Date(today.getFullYear(), today.getMonth(), today.getDate());
    console.log(CurrentDate)
    let LastDate = new Date(props.item.reservationInfo.from);
    console.log(LastDate > CurrentDate);

    return (
        <div className={LastDate > CurrentDate ? "container-green" : "container-blue"}>
            <div className="container-item">
                <p>{props.item.hotelInfoPreview.hotelName}</p>
                <p>{props.item.hotelInfoPreview.city}, {props.item.hotelInfoPreview.country}</p>
                <p>[{props.item.reservationInfo.from.substring(0, 10)}] — [{props.item.reservationInfo.to.substring(0, 10)}]</p>
                <p>Adults : {props.item.reservationInfo.numberOfAdults}</p>
                <p>Children : {props.item.reservationInfo.numberOfChildren}</p>

            </div>
            {LastDate > CurrentDate ? <Button variant="contained" color="secondary" size="medium" onClick={CancelReservation}>Cancel reservation</Button> :
                <div>
                    {props.item.offerReservations.offerReviewID === null ?
                        <div>
                            <div>
                                <Input value={newReview}
                                    onChange={(event, newValue) => {
                                        setNewReview(newValue);
                                    }}
                                    color='secondary'></Input>
                                <Rating value={newRating}
                                    onChange={(event, newValue) => {
                                        setNewRating(newValue);
                                    }}></Rating>
                                <div>
                                    <Button variant="contained" color="primary" size="medium">Add new review</Button>
                                </div>
                            </div>
                        </div>
                        :
                        <div>
                            <div>
                                <Input disabled={!editing} value={newReview}
                                    onChange={(event, newValue) => {
                                        setNewReview(newValue);
                                    }}
                                    color='secondary'></Input>
                                <Rating disabled={!editing} value={newRating}
                                    onChange={(event, newValue) => {
                                        setNewRating(newValue);
                                    }}></Rating>
                                <div>
                                    <Button variant="contained" color="primary" size="medium" onClick={ReviewEdit}>{buttonText}</Button>
                                    <Button variant="contained" color="secondary" size="medium">Delete review</Button>
                                </div>
                            </div>
                        </div>}
                </div>}

        </div>
    );
}

export default MyReservationListItem;