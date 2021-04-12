import { Button, Input } from '@material-ui/core';
import React, { useEffect, useState } from 'react';
import './MyReservationsListItem.css';
import Reservation from './Reservation';
import Rating from '@material-ui/lab/Rating';

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

    let tempReview = {
        "reviewID": 3,
        "content": "so bad place",
        "rating": 2,
        "creationDate": "1",
        "reviewerUsername": "alex"
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

    const DeleteReview = () =>
    {
        //axios delete
    };

    const UpdateReview = () => 
    {
        //axios put
    };


    let today = new Date();
    let CurrentDate = new Date(today.getFullYear(), today.getMonth(), today.getDate());
    let LastDateOfReservationString = props.item.offerReservations.reservationsInfo[props.item.offerReservations.reservationsInfo.length - 1].to.split("/");
    let LastDate = new Date(parseInt(LastDateOfReservationString[2]), parseInt(LastDateOfReservationString[1]), parseInt(LastDateOfReservationString[0]));

    return (
        <div className={LastDate > CurrentDate ? "container-green" : "container-blue"}>
            <div className="container-item">
                <p1>Hotel Name: {props.item.hotelInfoPreview.hotelName}</p1>
                <p>City: {props.item.hotelInfoPreview.city}</p>
                <p1>Country: {props.item.hotelInfoPreview.country}</p1>
                <h3>Reservations:</h3>
                {
                    props.item.offerReservations.reservationsInfo.map(item =>
                        (<Reservation key={item.id} item={item}></Reservation>))
                }
            </div>

            {LastDate > CurrentDate ? <Button variant="contained" color="secondary" size="large">Cancel reservation</Button> :
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
                                    <Button variant="contained" color="primary" size="large">Add new review</Button>
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
                                    <Button variant="contained" color="primary" size="large" onClick={ReviewEdit}>{buttonText}</Button>
                                    <Button variant="contained" color="secondary" size="large">Delete review</Button>
                                </div>
                            </div>
                        </div>}
                </div>}
        </div>
    );
}

export default MyReservationListItem;