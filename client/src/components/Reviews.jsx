import React, { Component, useEffect, useState } from 'react';
import axios from 'axios'
import { Link, useHistory, BrowserRouter as Router, Route } from 'react-router-dom';import ReviewsListItem from './ReviewsListItem';
;

function Reviews(props) {

    let hotelId = props.match.params.hotelId;
    let offerId = props.match.params.offerId;
    const [reviews, setReviews] = useState([]);

    useEffect(() => {
        fetchItems();
    }, []);



    const fetchItems = () => {
        const url = `/api-client/hotels/${hotelId}/offers/${offerId}/reviews`;
        axios.get(url, { headers: { 'accept': 'application/json', 'x-session-token': window.localStorage.getItem("token") } })
            .then(response => {
                console.log(response.data);
                setReviews(response.data);
            })
            .catch(error => {
                if (error.response.status === 404) {
                    let path = `/hotels/${hotelId}/offers`;
                    history.push(path);
                }
                console.error('There was an error!', error.response);
            });

        
    }

    const items = [
        {
            "reviewID": 1,
            "content": "very nice room , wifi is fast",
            "rating": 5,
            "creationDate": "2021-03-03",
            "reviewerUsername": "mailo001"
        },
        {
            "reviewID": 2,
            "content": "broken shower",
            "rating": 4,
            "creationDate": "2021-08-05",
            "reviewerUsername": "areknoster"
        }
    ]

    return (
        <div>
            {
                reviews.map(item =>
                    (<ReviewsListItem item={item} key={item.reviewID}></ReviewsListItem>))
            }
        </div>
    );
}

export default Reviews;