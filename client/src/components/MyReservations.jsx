import React, { useEffect, useState } from 'react';
import MyReservationsListItem from './MyReservationsListItem';
import './MyReservations.css'

function MyReservations (){

    useEffect(() => {
        fetchItems();
    }, []);

    const [items, setItems] = useState([]);

    const data = [
        {
          "hotelInfoPreview": {
            "hotelID": 4,
            "hotelName": "Grand",
            "country": "Poland",
            "city": "Warsaw"
          },
          "offerReservations": {
            "reservationsInfo": [
              {
                "reservationID": 3,
                "from": "2014-01-01T23:28:56.782Z",
                "to": "2014-01-03T23:28:56.782Z",
                "numberOfChildren": 1,
                "numberOfAdults": 2
              },
              {
                "reservationID": 3,
                "from": "2016-01-01T23:28:56.782Z",
                "to": "2016-01-03T23:28:56.782Z",
                "numberOfChildren": 0,
                "numberOfAdults": 4
              }
            ],
            "offerID": 15,
            "offerReviewID": 3,
          }
        },
        {
            "hotelInfoPreview": {
              "hotelID": 4,
              "hotelName": "Grand",
              "country": "Poland",
              "city": "Warsaw"
            },
            "offerReservations": {
              "reservationsInfo": [
                {
                  "reservationID": 3,
                  "from": "2014-01-01T23:28:56.782Z",
                  "to": "2014-01-03T23:28:56.782Z",
                  "numberOfChildren": 1,
                  "numberOfAdults": 2
                },
                {
                  "reservationID": 3,
                  "from": "2016-01-01T23:28:56.782Z",
                  "to": "2016-01-03T23:28:56.782Z",
                  "numberOfChildren": 3,
                  "numberOfAdults": 4
                }
              ],
              "offerID": 15,
              "offerReviewID": 3,
            }
          }
      ];

    const fetchItems = async () => {

    }


        return (
            <div>
                <div className="reservations-container">
                <h3>My reservations:</h3>
                {
                    data.map(item =>
                    (<MyReservationsListItem key={item.id} item={item}></MyReservationsListItem>)) 
                }
                </div>
            </div>
        );
}

export default MyReservations;