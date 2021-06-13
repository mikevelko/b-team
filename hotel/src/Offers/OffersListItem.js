import { Button, ButtonBase, Grid, makeStyles, Typography } from "@material-ui/core";
import React from 'react';
import { useEffect, useState } from "react";
import { useHistory } from "react-router";
import { TryDeleteHotelOffer, TryGetHotelOffers, TryGetRoomsForOffer } from "../Utils/FetchUtils";
import offerImage from './offer.png';

const useStyles = makeStyles((theme) => ({
    activeOfferItem:{
      marginBottom:20, 
      display:'flex', 
      justifyContent:'space-around', 
      backgroundColor:'#bfa1de', 
      padding:25, 
      borderRadius:20,
    },
    inactiveOfferItem:{
      marginBottom:20, 
      display:'flex', 
      justifyContent:'space-around', 
      backgroundColor:'#b4e4e4', 
      padding:25, 
      borderRadius:20,
      
    },
    partOfOfferItem:{
      marginRight:30,
    },
    editOfferButton:{
      backgroundColor:'#ffcc80', 
      color:'white',
      '&:hover': {
        background: "#ffcc80",
      },
      width:120,
      marginRight:20
    },
    offersListItemButtons:{
      display:'flex',
      justifyContent:'space-around',
      marginBottom:10,
      
    },
    deleteOfferButton:{
      backgroundColor:'#cc0000',
      color:'white',
      '&:hover': {
        background: "#cc0000",
      },
      width:120,
    },
    offerPreviewImage:{
      width:'150px', 
      borderRadius:5,
    },
    rightSide:{
      display:'flex',
      flexDirection:'column'
    },
  }));
export default function OffersListItem(props){
    const classes = useStyles();
    const [rooms,setRooms] = useState([])
    const history = useHistory();
    useEffect(() =>{
      GetRoomsForOffer(props.offer.offerID)
    },[])
    function EditOfferButton(event, offerID) {    
      event.stopPropagation();
      history.push(`/offers/edit/${offerID}`)
    }
    function OfferDetails(offerID){
      console.log(offerID)
      history.push(`/offers/${offerID}`)
    }
    function DeleteOfferButton(event,offerID) {
      event.stopPropagation()
      console.log(props)
      if(props.GetOffersWithRoomNumber != null){
        props.GetOffersWithRoomNumber()
      }else{
        TryDeleteHotelOffer(offerID).then(function (response) {
          TryGetHotelOffers()
          .then(function (response) {
            props.setOffersList(response)
          });
        })
      }
      
    }
    function GetRoomsForOffer(offerID) {
      TryGetRoomsForOffer(offerID).then(function (response) {
        setRooms(response)
      })
    }
    function OnClickOnRoom(event, hotelRoomNumber){
      event.stopPropagation()
      history.push(`/rooms/${hotelRoomNumber}`)
    }
    return(
        <Grid className={props.offer.isActive ? classes.activeOfferItem: classes.inactiveOfferItem} onClick={() => {OfferDetails(props.offer.offerID)}} container key={props.offer.offerID}>
            <Grid item className={classes.partOfOfferItem}>
                <ButtonBase>
                  <img src={offerImage} className={classes.offerPreviewImage}/>
                </ButtonBase>
            </Grid>
            <Grid className={classes.partOfOfferItem}>
              <Typography>{props.offer.offerTitle}</Typography>
              <Typography>Cost per child:{props.offer.costPerChild}</Typography>
              <Typography>Cost per adult:{props.offer.costPerAdult}</Typography>
              <Typography>Max guests:{props.offer.maxGuests}</Typography>
            </Grid>
            <Grid item className={classes.rightSide}>
              <Grid item className={classes.offersListItemButtons}>
                <Button className={classes.editOfferButton} onClick={(event) => {EditOfferButton(event,props.offer.offerID)}}>Edit offer</Button>
                <Button className={classes.deleteOfferButton}  onClick={(event) => {DeleteOfferButton(event,props.offer.offerID)}}>Delete offer</Button>

              </Grid>
              {rooms.length > 0 ? 
              (<Grid>
                <Typography>Rooms</Typography>
                <Grid className="roomList" >
                  {rooms.map((room)=>(
                    <Grid className="roomListItem" onClick={(event) => {OnClickOnRoom(event,room.hotelRoomNumber)}} key={room.roomID} >
                      <Typography>{room.hotelRoomNumber}</Typography>
                    </Grid>
                  ))
                  }
                </Grid>
              </Grid>) 
              : 
              (
                <></>
              )}
              
            </Grid>
          </Grid>
    )
}