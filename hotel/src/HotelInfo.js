import React from 'react';
import { Button, makeStyles, Typography, } from '@material-ui/core';
import templatePicture from './offer.png'; 
import './HotelInfo.css'
import { Link } from 'react-router-dom';

const useStyles = makeStyles((theme) => ({
  hotelPreviewPicture:{
    width:'300px', 
    borderRadius:5,
  },
  hotelPicture:{
    width:'90px',
    borderRadius:10,
    margin:5,

  },
  allImages:{
    display:'flex',
    flexDirection:'column',

  },
  hotelImages:{
    display: 'flex',
  },
  offerDetailsItem:{
    marginBottom:15,
  },
  editHotelInfoButton:{
    backgroundColor:'#ffcc80', 
    color:'white',
    margin:5,
  },

  fieldRow:{
    display:'flex', 
    flexDirection:'row', 
  },
  fieldRowDescription:{
    display:'flex', 
    flexDirection:'column', 
    alignItems:'flex-start',
    maxWidth:400,
  },

}));

function HotelInfo() {
  const classes = useStyles();

  return (
    <div className='hotelInfo'>
      <div className={classes.allImages}>
        <>
          <img src={templatePicture} className={classes.hotelPreviewPicture}/>
          <div className={classes.hotelImages}>
              <img src={templatePicture} className={classes.hotelPicture}/>
              <img src={templatePicture} className={classes.hotelPicture}/>
              <img src={templatePicture} className={classes.hotelPicture}/>
          </div>
        </>
        <>
          <Button component={Link} to='/hotelInfo/edit' className={classes.editHotelInfoButton} >
            Edit Hotel Info
          </Button>
        </>
      </div>
      <div className='hotelDetails'>
          <div className={classes.fieldRow}>
            <Typography className={classes.offerDetailsItem}>
              Hotel name:
            </Typography>
            <Typography>
              [HotelName]
            </Typography>
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            Country:
          </Typography>
          <Typography>
              [Country]
            </Typography>
          </div>
          <div className={classes.fieldRow}>
          <Typography className={classes.offerDetailsItem}>
            City:
          </Typography>
          <Typography>
              [City]
            </Typography>
          </div>
          <div className={classes.fieldRowDescription}>
          <Typography className={classes.offerDetailsItem}>
            Description:
          </Typography>
            <Typography variant='caption'>
              [Description]
            </Typography>
          </div>
        </div>
    </div>
  );
}

export default HotelInfo;