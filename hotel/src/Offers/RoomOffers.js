import React, { useEffect, useState } from 'react';
import "./Rooms.css"
import GridList from '@material-ui/core/GridList';
import { Button, Grid, makeStyles, TextField, Typography } from '@material-ui/core';
import { TryAddHotelRoom, TryGetHotelOffer, TryGetHotelRooms } from '../Utils/FetchUtils';
import { useHistory } from 'react-router';
import OffersListItem from './OffersListItem';
import './RoomOffers.css'
const useStyles = makeStyles((theme) => ({
    title:{
        marginTop:10,
        fontSize: 25
      },
}));
function RoomOffers() {
    const classes = useStyles()
    const history = useHistory();
    const [roomNumber,setRoomNumber] = useState("");
    const [offersWithRoomNumber,setOffersWithRoomNumber] = useState([]);
    useEffect(()=>{
        setRoomNumber(history.location.pathname.split('/')[2])
        GetOffersWithRoomNumber()
    },[])
    function GetOffersWithRoomNumber(){
        let offerIds = []
        //param is roomNumber
        TryGetHotelRooms(roomNumber).then(function (response) {
            offerIds = response[0].offerID
            setOffersWithRoomNumber([])
            offerIds.map((offerId) => {
                TryGetHotelOffer(offerId).then(function (response) {
                    response.offerID = offerId
                    setOffersWithRoomNumber(prevState => [...prevState,response]) 
                })
            })
        })
    }
    return (
        <div className="offers">
            <div className="titleRoomOffers">
                <Typography className={classes.title}>
                Offers with room: {roomNumber}
                </Typography>
            </div>

            <div className="roomOffers">
                {offersWithRoomNumber.map((item,index) => {
                    return <OffersListItem GetOffersWithRoomNumber={GetOffersWithRoomNumber} key={index} offer={item} setOffersList={setOffersWithRoomNumber}/>
                })}
            </div>
        </div>

    )
}

export default RoomOffers;