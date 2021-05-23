import React, { Component, useEffect, useState } from 'react';
import axios from 'axios'
import { Link, useHistory, BrowserRouter as Router, Route } from 'react-router-dom';import ReviewsListItem from './ReviewsListItem';
;

function Reviews() {
    useEffect(() => {
        fetchItems();
    }, []);



    const fetchItems = () => {

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
                items.map(item =>
                    (<ReviewsListItem item={item} key={item.reviewID}></ReviewsListItem>))
            }
        </div>
    );
}

export default Reviews;