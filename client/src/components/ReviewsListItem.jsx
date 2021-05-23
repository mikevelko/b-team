import React, { Component } from 'react';
import Rating from '@material-ui/lab/Rating';

function ReviewsListItem(props) {
    return (
        <div className="container">
            <div className="container-item-hotel">
                <p>{props.item.creationDate}</p>
                <p>{props.item.reviewerUsername}: &quot;{props.item.content}&quot;</p>
                <Rating disabled={true} value={props.item.rating}></Rating>
            </div>


        </div>
    );
}

export default ReviewsListItem;