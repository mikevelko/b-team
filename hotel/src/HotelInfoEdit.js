import React, { useEffect, useState } from 'react';
import { Button, ButtonBase, makeStyles, TextField, Typography, } from '@material-ui/core';
import templatePicture from './offer.png'; 
import './HotelInfoEdit.css'
import { useHistory } from 'react-router-dom';
import ClearIcon from '@material-ui/icons/Clear';
import {TryGetHotelInfo, TryPatchHotelInfo} from './FetchUtils.js';

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
    const history = useHistory()


    const [hotelName,setHotelName] = useState('');
    const [hotelDescription,setHotelDescription] = useState('');
    const [city,setCity] = useState('');
    const [country,setCountry] = useState('');

      // For feature
    const [pictures,setPictures] = useState([]);
    const [previewPicture,setPreviewPicture] = useState('');

    function GetHotelInfo(){
      TryGetHotelInfo().then(function(response) {
        if(response!=''){

          setHotelName(response.hotelName)
          setCity(response.city)
          setCountry(response.country)
          setHotelDescription(response.hotelDesc)
          // For feature
          setPictures(response.hotelPictures)
          setPreviewPicture(response.hotelPreviewPicture)
        }
      })
    }

    useEffect(()=>{
      GetHotelInfo()
    },[])
    
    function OnClickSaveChangesButton() {
      if(hotelName !==''&&hotelDescription!==''&&city!==''&&country!==''&& hotelName!==undefined &&  hotelDescription !==undefined && city !== undefined && country !== undefined){
        TryPatchHotelInfo(hotelName,hotelDescription,city,country).then(function (response) {
          if(response.status === 200) history.push('/hotelInfo')
        })
      }else{
        alert("fulfil all fields")
      }
    } 
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
                <Button onClick={() => {OnClickSaveChangesButton()}} className={classes.saveHotelInfoButton} >
                    Save changes
                </Button>
            </>
          </div>
          <div className='hotelDetails'>
              <div className={classes.fieldRow}>
                <Typography className={classes.hotelDetailsItem}>
                  Hotel name:
                </Typography>
                <TextField size='small' value={hotelName} onChange={(e) =>{setHotelName(e.target.value)}}>
                </TextField>
              </div>
              <div className={classes.fieldRow}>
              <Typography className={classes.hotelDetailsItem}>
                Country:
              </Typography>
              <TextField size='small' value={country} onChange={(e) =>{setCountry(e.target.value)}}>
                </TextField>
              </div>
              <div className={classes.fieldRow}>
              <Typography className={classes.hotelDetailsItem}>
                City:
              </Typography>
              <TextField size='small' value={city} onChange={(e) =>{setCity(e.target.value)}}>
                </TextField>
              </div>
              <div className={classes.fieldRowDescription}>
              <Typography className={classes.hotelDetailsItem}>
                Description:
              </Typography>
                <TextField multiline fullWidth value={hotelDescription} onChange={(e) =>{setHotelDescription(e.target.value)}}>
                </TextField>
              </div>
            </div>
        </div>
      );
}

export default HotelInfoEdit;