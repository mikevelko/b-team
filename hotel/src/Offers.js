import { Button, ButtonBase, Grid, List, Typography } from '@material-ui/core';
import React, { useEffect, useState } from 'react';
import { Link, matchPath, Redirect, useHistory } from 'react-router-dom';
import './Offers.css'
import offerImage from './offer.png';
import { makeStyles } from '@material-ui/core/styles';
import { TryDeleteHotelOffer, TryGetHotelOffers } from './FetchUtils';

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
    color:'white',
    '&:hover': {
      background: "#ffcc80",
    },
  },
  deleteOfferButton:{
    backgroundColor:'#cc0000',
    color:'white',
    '&:hover': {
      background: "#cc0000",
    },
  },
  offerPreviewImage:{
    width:'150px', 
    borderRadius:5,
  },
}));

function Offers() {
  const classes = useStyles();
  const history = useHistory();
  const [offersList,setOffersList] = useState([]);

  useEffect(()=>{
    TryGetHotelOffers()
      .then(function (response) {
        setOffersList(response)
      });
  },[])

  useEffect(()=>{
    console.log(offersList)
  },[offersList])
  function EditOfferButton(event, offerID) {
    console.log("aasfafsasafs")

    event.stopPropagation();
    history.push(`/offers/edit/${offerID}`)
  }
  function OfferDetails(offerID){
    console.log(offerID)
    history.push(`/offers/${offerID}`)
  }
  function DeleteOfferButton(event,offerID) {
    event.stopPropagation()
    TryDeleteHotelOffer(offerID).then(function (response) {
      TryGetHotelOffers()
      .then(function (response) {
        setOffersList(response)
      });
    })

  }
  return (
    <div className="offers">

      <div className="filterButtons">
        <Button style={{backgroundColor:'#ffcc80', color:'white' }} component={Link} to='/offers/create' >Add new offer</Button>
        <Button style={{backgroundColor:'#3f51b5', color:'white' }}>All offers:[count]</Button>
        <Button style={{backgroundColor:'#bfa1de', color:'white' }}>Active offers:[count]</Button>
        <Button style={{backgroundColor:'#b4e4e4', color:'white' }}>Inactive offers:[count]</Button>
      </div>
      <div className="offersList">     
        {offersList.map((item,_) => {
          return item.isActive ? 
          (<Grid className={classes.activeOfferItem} onClick={() => {OfferDetails(item.offerID)}} container key={item.offerID}>
            <Grid item className={classes.partOfOfferItem}>
                <ButtonBase>
                  <img src={offerImage} className={classes.offerPreviewImage}/>
                </ButtonBase>
            </Grid>
            <Grid item className={classes.partOfOfferItem}>
              <Typography>{item.offerTitle}</Typography>
              <Typography>Cost per child:{item.costPerChild}</Typography>
              <Typography>Cost per adult:{item.costPerAdult}</Typography>
              <Typography>Max guests:{item.maxGuests}</Typography>
            </Grid>
            <Grid item style={{display:'flex',flexDirection:'column'}}>
              <Grid item style={{display:'flex', justifyContent:'space-around', marginBottom:10}}>
                <Button className={classes.editOfferButton} onClick={(event) => {EditOfferButton(event,item.offerID)}}>Edit offer</Button>
                <Button className={classes.deleteOfferButton}  onClick={(event) => {DeleteOfferButton(event,item.offerID)}}>Delete offer</Button>

              </Grid>
              <Grid item>
  
                <Typography>Room numbers:[list of room numbers]</Typography>
              </Grid>
            </Grid>
            
          </Grid>) 
        :
        (<Grid className={classes.inactiveOfferItem} container onClick={() => {OfferDetails(item.offerID)}} key={item.offerID}>
          <Grid item className={classes.partOfOfferItem}>
              <ButtonBase>
                <img src={offerImage} className={classes.offerPreviewImage}/>
              </ButtonBase>
          </Grid>
          <Grid item className={classes.partOfOfferItem}>
            <Typography>{item.offerTitle}</Typography>
            <Typography>Cost per child:{item.costPerChild}</Typography>
            <Typography>Cost per adult:{item.costPerAdult}</Typography>
            <Typography>Max guests:{item.maxGuests}</Typography>
          </Grid>

          <Grid item style={{display:'flex',flexDirection:'column'}}>
            <Grid item style={{display:'flex', justifyContent:'space-around', marginBottom:10}}>
            <Button className={classes.editOfferButton} onClick={(event) => {EditOfferButton(event,item.offerID)}}>Edit offer</Button>
                <Button className={classes.deleteOfferButton}  onClick={(event) => {DeleteOfferButton(event,item.offerID)}}>Delete offer</Button>

            </Grid>
            <Grid item>

              <Typography>Room numbers:[list of room numbers]</Typography>
            </Grid>
          </Grid>
          
        </Grid>)
        })}
      </div>
    </div>
  );
}

export default Offers;