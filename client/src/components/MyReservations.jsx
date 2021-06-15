import React, { useEffect, useState } from 'react';
import MyReservationsListItem from './MyReservationsListItem';
import './MyReservations.css';
import axios from 'axios';

function MyReservations() {

  useEffect(() => {
    fetchItems();
  }, []);

  const [items, setItems] = useState([]);


  const fetchItems = () => {
    const url = `/api-client/client/reservations?pageNumber=1&pageSize=10`;
    axios.get(url, { headers: { 'accept': 'application/json', 'x-client-token': window.localStorage.getItem("token") } })
      .then(response => {
        console.log(response.data);
        setItems(response.data);
      })
      .catch(error => {
        //console.error('There was an error!', error.response);
      });
  }
  



  return (
    <div>
      <div className="reservations-container">

        {
          items.map(item =>
            (<MyReservationsListItem key={item.reservationInfo.reservationId} item={item} fetchReservations={fetchItems}></MyReservationsListItem>))
        }
      </div>
    </div>
  );
}

export default MyReservations;