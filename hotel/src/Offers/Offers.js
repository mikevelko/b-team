import { Button } from '@material-ui/core';
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import './Offers.css'

import {  TryGetHotelOffers } from '../Utils/FetchUtils';
import OffersListItem from './OffersListItem';


function Offers() {

  const [offersList,setOffersList] = useState([]);
  const [isActive, setIsActive] = useState(null);
  useEffect(()=>{
    TryGetHotelOffers(isActive)
      .then(function (response) {
        setOffersList(response)
      });
  },[isActive])
  return (
    <div className="offers">

      <div className="filterButtons">
        <Button style={{backgroundColor:'#ffcc80', color:'white' }} component={Link} to='/offers/create' >Add new offer</Button>
        <Button style={{backgroundColor:'#3f51b5', color:'white' }} onClick={() => {setIsActive(null)}}>All offers</Button>
        <Button style={{backgroundColor:'#bfa1de', color:'white' }} onClick={() => {setIsActive(true)}}>Active offers</Button>
        <Button style={{backgroundColor:'#b4e4e4', color:'white' }} onClick={() => {setIsActive(false)}}>Inactive offers</Button>
      </div>
      <div className="offersList">     
        {offersList.map((item,_) => {
          return <OffersListItem key={item.offerID} offer={item} setOffersList={setOffersList}/>
        })}
      </div>
    </div>
  );
}

export default Offers;