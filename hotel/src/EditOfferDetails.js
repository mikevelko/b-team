import { Button, ButtonBase, makeStyles, TextField, Typography, } from '@material-ui/core';
import ClearIcon from '@material-ui/icons/Clear';
import React, { useState } from 'react';
import templatePicture from './offer.png'; 
import './CreateOffer.css'
import { Link } from 'react-router-dom';
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
    marginBottom:15,
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
    alignItems:'flex-start',
  },
}));


function EditOfferDetails({match}) {

    console.log(match)
  const classes = useStyles();
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
          <Button className={classes.saveButton} component={Link} to='/offers'>
            Save changes
          </Button>
        </>
      </div>
      <div className='offerDetails'>
          <div className={classes.fieldRow}>
            <Typography className={classes.offerDetailsItem}>
              Offer title:
            </Typography>
            <TextField size='small'>
            </TextField>  
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            Cost per child:
          </Typography>
            <TextField size='small'>
            </TextField>  
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            Cost per adult:
          </Typography>
            <TextField size='small'>
            </TextField>  
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            Max guests:
          </Typography>
            <TextField size='small'>
            </TextField>  
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            Active status:
          </Typography>
            <TextField size='small'>
            </TextField>  
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}> 
            Rooms:
          </Typography>
            <TextField size='small'>
            </TextField>  
          </div>


        </div>
    </div>
  );
}

export default EditOfferDetails;