import { Button, ButtonBase, Grid, List, Typography } from '@material-ui/core';
import React, { useEffect, useState } from 'react';
import { Link, Redirect } from 'react-router-dom';
import './Offers.css'
import offerImage from './offer.png';
import { makeStyles } from '@material-ui/core/styles';
import { TryGetHotelOffers } from './FetchUtils';
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

  const [offersList,setOffersList] = useState([    {
    "IsActive": true,
    "OfferTitle": "string",
    "CostPerChild": "0",
    "CostPerAdult": "0",
    "MaxGuests": 1,
    "Description": "string",
    "OfferPreviewPicture": "string",
    "Pictures": null,
    "Rooms": null
  }]);

  useEffect(()=>{
    TryGetHotelOffers()
      .then(function (response) {
        setOffersList(response)
      });
  },[])
  useEffect(()=>{
    console.log(offersList.map((item,_) => item))
  },[offersList])

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
          return item.IsActive ? 
          (<Grid className={classes.activeOfferItem} container>
            <Grid item className={classes.partOfOfferItem}>
                <ButtonBase>
                  <img src={offerImage} className={classes.offerPreviewImage}/>
                </ButtonBase>
            </Grid>
            <Grid item className={classes.partOfOfferItem}>
              <Typography>{item.OfferTitle}</Typography>
              <Typography>Cost per child:{item.CostPerChild}</Typography>
              <Typography>Cost per adult:{item.CostPerAdult}</Typography>
              <Typography>Max guests:{item.MaxGuests}</Typography>
            </Grid>
            <Grid item style={{display:'flex',flexDirection:'column'}}>
              <Grid item style={{display:'flex', justifyContent:'space-around', marginBottom:10}}>
                <Button className={classes.editOfferButton} component={Link} to='/offers/edit/5' onClick={(event) => {event.stopPropagation()}}>Edit offer</Button>
                <Button className={classes.deleteOfferButton}  onClick={(event) => {event.stopPropagation()}}>Delete offer</Button>
              </Grid>
              <Grid item>
  
                <Typography>Room numbers:[list of room numbers]</Typography>
              </Grid>
            </Grid>
            
          </Grid>) 
        :
        (<Grid className={classes.inactiveOfferItem} container>
          <Grid item className={classes.partOfOfferItem}>
              <ButtonBase>
                <img src={offerImage} className={classes.offerPreviewImage}/>
              </ButtonBase>
          </Grid>
          <Grid item className={classes.partOfOfferItem}>
            <Typography>{item.OfferTitle}</Typography>
            <Typography>Cost per child:{item.CostPerChild}</Typography>
            <Typography>Cost per adult:{item.CostPerAdult}</Typography>
            <Typography>Max guests:{item.MaxGuests}</Typography>
          </Grid>

          <Grid item style={{display:'flex',flexDirection:'column'}}>
            <Grid item style={{display:'flex', justifyContent:'space-around', marginBottom:10}}>
              <Button className={classes.editOfferButton} component={Link} to='/offers/edit/' onClick={(event) => {event.stopPropagation()}}>Edit offer</Button>
              <Button className={classes.deleteOfferButton}  onClick={(event) => {event.stopPropagation()}}>Delete offer</Button>
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