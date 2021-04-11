import React from 'react';
import { Button, ButtonBase, makeStyles, TextField, Typography, } from '@material-ui/core';
import templatePicture from './offer.png'; 
import './HotelInfoEdit.css'
import { Link } from 'react-router-dom';
import ClearIcon from '@material-ui/icons/Clear';


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
    hotelAllImages:{
      display: 'flex',
      flexDirection: 'column',
    },
    hotelImages:{
        display:'flex',
    },
    hotelDetailsItem:{
      marginBottom:15,
    },
    fieldRow:{
      display:'flex', 
      flexDirection:'row', 
    },
    fieldRowDescription:{
      display:'flex', 
      flexDirection:'column', 
      alignItems:'flex-start',
      width:350,
    },
    saveHotelInfoButton:{
        backgroundColor:'#80ff80', 
        color:'white',
        margin:5,
    },
    hotelImageView:{
        width:'auto', 
        height:'auto', 
        position:'relative',
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
      setImageButton:{
        backgroundColor:'#ffcc80', 
        color:'white',
        margin:5,
      },
  }));
function HotelInfoEdit() {
    const classes = useStyles();

    return (
        <div className='hotelInfoEdit'>
            <div className={classes.hotelAllImages}>
            <>
                <div className={classes.hotelImageView}>
                    <img src={templatePicture} className={classes.hotelPreviewPicture}/>
                    <ButtonBase className={classes.deletePreviewImageButton} onClick={()=>{}}>
                    <ClearIcon >
                    </ClearIcon>
                    </ButtonBase>
                </div>
                <div className={classes.hotelImages}>
                    <div className={classes.hotelImageView}>
                        <img src={templatePicture} className={classes.hotelPicture}/>
                        <ButtonBase className={classes.deleteImageButton}>
                            <ClearIcon >
                            </ClearIcon>
                        </ButtonBase>
                        </div>
                    <div className={classes.hotelImageView}>
                        <img src={templatePicture} className={classes.hotelPicture}/>
                        <ButtonBase className={classes.deleteImageButton}>
                            <ClearIcon >
                            </ClearIcon>
                        </ButtonBase>
                    </div>
                    <div className={classes.hotelImageView}>
                        <img src={templatePicture} className={classes.hotelPicture}/>
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
                <Button component={Link} to='/hotelInfo' className={classes.saveHotelInfoButton} >
                    Save changes
                </Button>
            </>
          </div>
          <div className='hotelDetails'>
              <div className={classes.fieldRow}>
                <Typography className={classes.hotelDetailsItem}>
                  Hotel name:
                </Typography>
                <TextField size='small'>
                </TextField>
              </div>
              <div className={classes.fieldRow}>
              <Typography className={classes.hotelDetailsItem}>
                Country:
              </Typography>
              <TextField size='small'>
                </TextField>
              </div>
              <div className={classes.fieldRow}>
              <Typography className={classes.hotelDetailsItem}>
                City:
              </Typography>
              <TextField size='small'>
                </TextField>
              </div>
              <div className={classes.fieldRowDescription}>
              <Typography className={classes.hotelDetailsItem}>
                Description:
              </Typography>
                <TextField multiline fullWidth >
                </TextField>
              </div>
            </div>
        </div>
      );
}

export default HotelInfoEdit;