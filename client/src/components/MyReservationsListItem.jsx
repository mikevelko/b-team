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
    const [newRating, setNewRating] = useState(0);
    const [reviewItem, setReviewItem] = useState([])

    useEffect(() => {
        fetchItems();
    }, []);

    const fetchItems = async () => {
        if ((props.item.reservationInfo.hasOwnProperty('reviewID') && props.item.reservationInfo.reviewID !== null)) {
            const url = `/api-client/client/reservations/${props.item.reservationInfo.reservationID}/review`;
            axios.get(url, { headers: { 'accept': '*/*', 'x-session-token': window.localStorage.getItem("token") } })
                .then(response => {
                    console.log(response.data);
                    setReviewItem(response.data);
                    setNewReview(response.data.content);
                    setNewRating(response.data.rating);
                })
                .catch(error => {
                    //console.error('There was an error!', error.response);
                });
        }
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

    const AddEditReview = () => {
        const data =
        {
            content: newReview.toString(),
            rating: parseInt(newRating)
        }

        const url = `/api-client/client/reservations/${props.item.reservationInfo.reservationID}/review`;
        axios.put(url,
            data,
            { headers: { 'accept': '*/*', 'Content-Type': 'application/json', 'x-session-token': window.localStorage.getItem("token") } })
            .then(response => {
                console.log(response.data);
                props.fetchReservations();
            })
            .catch(error => {
                console.error('There was an error!', error.response);
            });
    }

    const DeleteReview = () => {
        const url = `/api-client/client/reservations/${props.item.reservationInfo.reservationID}/review`;
        axios.delete(url, { headers: { 'accept': '*/*', 'x-session-token': window.localStorage.getItem("token") } })
            .then(response => {
                console.log(response.data);
                props.fetchReservations();
                setNewReview("");
                setNewRating(0);
            })
            .catch(error => {
                //console.error('There was an error!', error.response);
            });
    };


    let today = new Date();
    let CurrentDate = new Date(today.getFullYear(), today.getMonth(), today.getDate());
    let LastDate = new Date(props.item.reservationInfo.from);


    return (
        <div className={LastDate > CurrentDate ? "container-green" : "container-blue"}>
            <div className="container-item">
                <p>{props.item.hotelInfoPreview.hotelName}</p>
                <p>{props.item.reservationInfo.reservationID}</p>
                <p>{props.item.hotelInfoPreview.city}, {props.item.hotelInfoPreview.country}</p>
                <p>[{props.item.reservationInfo.from.substring(0, 10)}] — [{props.item.reservationInfo.to.substring(0, 10)}]</p>
                <p>Adults : {props.item.reservationInfo.numberOfAdults}</p>
                <p>Children : {props.item.reservationInfo.numberOfChildren}</p>

            </div>
            {LastDate > CurrentDate ? <Button variant="contained" color="secondary" size="small" onClick={CancelReservation}>Cancel reservation</Button> :
                <div>
                    {console.log(!(props.item.hasOwnProperty('reviewID')))}
                    {(!(props.item.reservationInfo.hasOwnProperty('reviewID')) || props.item.reservationInfo.reviewID === null) ?
                        <div>
                            <div>
                                <Input value={newReview}
                                    onChange={(event, newValue) => {
                                        setNewReview(event.target.value);
                                    }}
                                    color='secondary'></Input>
                                <Rating value={newRating}
                                    onChange={(event, newValue) => {
                                        setNewRating(event.target.value);
                                    }}></Rating>
                                <div>
                                    <Button variant="contained" color="primary" size="small" onClick={AddEditReview}>Add new review</Button>
                                </div>
                            </div>
                        </div>
                        :
                        <div>
                            <div>
                                <Input disabled={!editing} value={newReview}
                                    onChange={(event, newValue) => {
                                        setNewReview(event.target.value);
                                    }}
                                    color='secondary'></Input>
                                <Rating disabled={!editing} value={newRating}
                                    onChange={(event, newValue) => {
                                        setNewRating(event.target.value);
                                    }}></Rating>
                                <div>
                                    <Button variant="contained" color="primary" size="small" onClick={ReviewEdit}>{buttonText}</Button>
                                    <Button variant="contained" color="primary" size="small" onClick={AddEditReview}>Update review</Button>
                                    <Button variant="contained" color="secondary" size="small" onClick={DeleteReview}>Delete review</Button>
                                </div>
                            </div>
                        </div>}
                </div>}

        </div>
    );
}

export default MyReservationListItem;