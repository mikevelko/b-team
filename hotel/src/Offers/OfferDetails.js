import { Button, ButtonBase, Checkbox, makeStyles, TextField, Typography, } from '@material-ui/core';
import ClearIcon from '@material-ui/icons/Clear';
import React, { useEffect, useState } from 'react';
import templatePicture from './offer.png'; 
import './CreateOffer.css'
import { Link, useHistory } from 'react-router-dom';
import { TryGetHotelOffer, TryGetHotelRooms, TryGetRoomsForOffer } from '../Utils/FetchUtils';
const useStyles = makeStyles((theme) => ({
  offerPreviewImage:{
    width:'300px', 
    borderRadius:5,
  },
  offerImage:{
    width:'90px',
    borderRadius:10,
    margin:5,
  },
  allImages:{
    display:'flex',
    flexDirection:'column',

  },
  offerImages:{
    display: 'flex',
  },
  deletePreviewImageButton:{
    backgroundColor: 'red',
    position: 'absolute',
    borderRadius:5,
    margin:5,
    top: 0,
    right: 0,
  },
  deleteImageButton:{
    backgroundColor: 'red',
    position: 'absolute',
    borderRadius:5,
    margin:10,
    top: 0,
    right: 0,
  },
  offerDetailsItem:{
  },
  editOfferButton:{
    backgroundColor:'#ffcc80', 
    color:'white',
    color:'white',
    '&:hover': {
      background: "#ffcc80",
    },
  },
  setImageButton:{
    backgroundColor:'#ffcc80', 
    color:'white',
    margin:5,
    '&:hover': {
        background: "#ffcc80",
      },
  },
  offerImageView:{
    width:'auto', 
    height:'auto', 
    position:'relative',
  },
  fieldRow:{
    display:'flex', 
    flexDirection:'row', 
    alignItems:'center',
    marginBottom:15,

  },
}));


function OfferDetails() {
  const history = useHistory();

  const [offerTitle,setOfferTitle] = useState('');
  const [costPerChild,setCostPerChild] = useState(5);
  const [costPerAdult,setCostPerAdult] = useState(5);
  const [maxGuests,setMaxGuests] = useState(1);
  const [activeStatus, setActiveStatus] = useState(false);
  const [description, setDescription] = useState('');

    // For feature
  const [pictures,setPictures] = useState(["string"]);
  const [previewPicture,setPreviewPicture] = useState('');

  useEffect(()=>{
    TryGetHotelOffer(history.location.pathname.split('/')[2]).then(function (response) {
      if(response!= ""){
        setActiveStatus(response.isActive)
        setOfferTitle(response.offerTitle)
        setCostPerChild(response.costPerChild)
        setCostPerAdult(response.costPerAdult)
        setDescription(response.description)
        setMaxGuests(response.maxGuests)
        setPreviewPicture(response.offerPreviewPicture)
        setPictures(response.pictures)
      }
    })
  },[])
  const classes = useStyles();

  function EditOfferButton() {
    history.push(`/offers/edit/${history.location.pathname.split('/')[2]}`)
  }

  return (
    <div className='createOffer'>
      <div className={classes.allImages}>
        <>
          <div className={classes.offerImageView}>
            <img src={templatePicture} className={classes.offerPreviewImage}/>
          </div>
          <div className={classes.offerImages}>
            <div className={classes.offerImageView}>
              <img src={templatePicture} className={classes.offerImage}/>

            </div>
            <div className={classes.offerImageView}>
              <img src={templatePicture} className={classes.offerImage}/>

            </div>
            <div className={classes.offerImageView}>
              <img src={templatePicture} className={classes.offerImage}/>
            </div>
          </div>
        </>
        <>
          <Button className={classes.editOfferButton} onClick={() => EditOfferButton()}>
            Edit offer
          </Button>
        </>
      </div>
      <div className='offerDetails'>
          <div className={classes.fieldRow}>
            <Typography className={classes.offerDetailsItem}>
              Offer title:
            </Typography>
            <Typography>
                {offerTitle}    
            </Typography> 
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            Cost per child:
          </Typography>
            <Typography>
                {costPerChild}    
            </Typography> 
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            Cost per adult:
          </Typography>
            <Typography>
                {costPerAdult}    
            </Typography> 
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            Max guests:
          </Typography>
            <Typography>
                {maxGuests}
            </Typography>
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            Active status:
          </Typography>
          <Typography>
              {activeStatus ? 'active' : 'inactive'}
          </Typography>

          </div>
          <div className={classes.fieldRowDescription}>
          <Typography className={classes.offerDetailsItem}> 
            Description:
          </Typography>
            <Typography>
                {description}
            </Typography>  
          </div>
        </div>
    </div>
  );
}

export default OfferDetails;