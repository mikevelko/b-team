import { Button, ButtonBase, Checkbox, makeStyles, TextField, Typography, } from '@material-ui/core';
import ClearIcon from '@material-ui/icons/Clear';
import React, { useEffect, useState } from 'react';
import templatePicture from './offer.png'; 
import './CreateOffer.css'
import { Link, useHistory } from 'react-router-dom';
import { TryEditHotelOffer, TryGetHotelOffer,TryGetRoomsForOffer } from '../Utils/FetchUtils';
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
  saveButton:{
    backgroundColor:'#80ff80', 
    color:'white',
    margin:5,
    '&:hover': {
        background: "#80ff80",
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


function EditOffer() {
  const history = useHistory();

  const [offerTitle,setOfferTitle] = useState('');
  const [costPerChild,setCostPerChild] = useState(5);
  const [costPerAdult,setCostPerAdult] = useState(5);
  const [maxGuests,setMaxGuests] = useState(1);
  const [activeStatus, setActiveStatus] = useState(false);
  const [rooms, setRooms] = useState();
  const [description, setDescription] = useState('');

    // For feature
  const [pictures,setPictures] = useState(["string"]);
  const [previewPicture,setPreviewPicture] = useState('');

  useEffect(()=>{
    TryGetHotelOffer(history.location.pathname.split('/')[3]).then(function (response) {
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
    TryGetRoomsForOffer(history.location.pathname.split('/')[3]).then(function (response) {
      setRooms(response.map((item) => {return  item.roomID.toString()}).join(' '))
    })
  },[])
  const classes = useStyles();

  function SaveChangesButton(){
    TryEditHotelOffer(history.location.pathname.split('/')[3],offerTitle,maxGuests,activeStatus,description,rooms.length > 0 ? rooms.trim().split(' ') : [],pictures,previewPicture)
      .then(function (response) {
        if(response!==''){
          history.push('/offers')
        }
      })
  }
  return (
    <div className='createOffer'>
      <div className={classes.allImages}>
        <>
          <div className={classes.offerImageView}>
            <img src={templatePicture} className={classes.offerPreviewImage}/>

            <ButtonBase className={classes.deletePreviewImageButton} onClick={()=>{}}>
              <ClearIcon >
              </ClearIcon>
            </ButtonBase>
          </div>
          <div className={classes.offerImages}>
            <div className={classes.offerImageView}>
              <img src={templatePicture} className={classes.offerImage}/>
              <ButtonBase className={classes.deleteImageButton}>
                <ClearIcon >
                </ClearIcon>
              </ButtonBase>
            </div>
            <div className={classes.offerImageView}>
              <img src={templatePicture} className={classes.offerImage}/>
              <ButtonBase className={classes.deleteImageButton}>
                <ClearIcon >
                </ClearIcon>
              </ButtonBase>
            </div>
            <div className={classes.offerImageView}>
              <img src={templatePicture} className={classes.offerImage}/>
              <ButtonBase className={classes.deleteImageButton}>
                <ClearIcon >
                </ClearIcon>
              </ButtonBase>
            </div>
          </div>
        </>
        <>
          <Button className={classes.setImageButton}>
            Set preview image
          </Button>
          <Button className={classes.setImageButton}>
            Add image
          </Button>
          <Button className={classes.saveButton} onClick={()=>{SaveChangesButton()}}>
            Save changes
          </Button>
        </>
      </div>
      <div className='offerDetails'>
          <div className={classes.fieldRow}>
            <Typography className={classes.offerDetailsItem}>
              Offer title:
            </Typography>
            <TextField value={offerTitle} onChange={(e) => {setOfferTitle(e.target.value)}}>
            </TextField>  
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            Cost per child:
          </Typography>
            <TextField type='number' value={costPerChild} onChange={(e) => {setCostPerChild(e.target.value)}}>
            </TextField>  
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            Cost per adult:
          </Typography>
            <TextField type='number' value={costPerAdult} onChange={(e) => {setCostPerAdult(e.target.value)}}>
            </TextField>  
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            Max guests:
          </Typography>
            <TextField type='number' value={maxGuests} onChange={(e) => {setMaxGuests(e.target.value)}}>
            </TextField>  
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            Active status:
          </Typography>
          <Checkbox checked={activeStatus} color='primary' onChange={(e) =>{setActiveStatus(e.target.checked)}}>
          </Checkbox>

          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}> 
            Rooms:
          </Typography>
            <TextField value={rooms} onChange={(e)=>{setRooms(e.target.value)}}>
            </TextField>  
          </div>
          <div className={classes.fieldRowDescription}>
          <Typography className={classes.offerDetailsItem}> 
            Description:
          </Typography>
            <TextField multiline fullWidth value={description} onChange={(e) => {setDescription(e.target.value)}}>
            </TextField>  
          </div>
        </div>
    </div>
  );
}

export default EditOffer;