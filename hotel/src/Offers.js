import { Button, ButtonBase, Grid, Typography } from '@material-ui/core';
import React from 'react';
import { Link } from 'react-router-dom';
import './Offers.css'
import offerImage from './offer.png';
import { makeStyles } from '@material-ui/core/styles';
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
  },
  deleteOfferButton:{
    backgroundColor:'#cc0000',
     color:'white',
  },
  offerPreviewImage:{
    width:'150px', 
    borderRadius:5,
  },
}));

function Offers() {
  const classes = useStyles();
  return (
    <div className="offers">

      <div className="filterButtons">
        <Button style={{backgroundColor:'#ffcc80', color:'white' }} component={Link} to='/offers/create' >Add new offer</Button>
        <Button style={{backgroundColor:'#3f51b5', color:'white' }}>All offers:[count]</Button>
        <Button style={{backgroundColor:'#bfa1de', color:'white' }}>Active offers:[count]</Button>
        <Button style={{backgroundColor:'#b4e4e4', color:'white' }}>Inactive offers:[count]</Button>
      </div>
      <div className="offersList">     
        <Grid className={classes.activeOfferItem} container>
          <Grid item className={classes.partOfOfferItem}>
              <ButtonBase>
                <img src={offerImage} className={classes.offerPreviewImage}/>
              </ButtonBase>
          </Grid>
          <Grid item className={classes.partOfOfferItem}>
            <Typography>[Title offer]</Typography>
            <Typography>Cost per child:[cost]</Typography>
            <Typography>Cost per adult:[cost]</Typography>
            <Typography>Max guests:[count]</Typography>
          </Grid>
          <Grid item style={{display:'flex',flexDirection:'column'}}>
            <Grid item style={{display:'flex', justifyContent:'space-around', marginBottom:10}}>
              <Button className={classes.editOfferButton} component={Link} to='/offers/edit/:OfferId'>Edit offer</Button>
              <Button className={classes.deleteOfferButton}>Delete offer</Button>
            </Grid>
            <Grid item>
              <Typography>Room numbers:[list of room numbers]</Typography>
            </Grid>
          </Grid>
        </Grid>
        <Grid className={classes.inactiveOfferItem} container>
          <Grid item className={classes.partOfOfferItem}>
              <ButtonBase>
                <img src={offerImage} className={classes.offerPreviewImage}/>
              </ButtonBase>
          </Grid>
          <Grid item className={classes.partOfOfferItem}>
            <Typography>[Title offer]</Typography>
            <Typography>Cost per child:[cost]</Typography>
            <Typography>Cost per adult:[cost]</Typography>
            <Typography>Max guests:[count]</Typography>
          </Grid>
          <Grid item style={{display:'flex',flexDirection:'column'}}>
            <Grid item style={{display:'flex', justifyContent:'space-around', marginBottom:10}}>
              <Button className={classes.editOfferButton}  component={Link} to='/offers/edit/:OfferId'>Edit offer</Button>
              <Button className={classes.deleteOfferButton}>Delete offer</Button>
            </Grid>
            <Grid item>
              <Typography>Room numbers:[list of room numbers]</Typography>
            </Grid>
          </Grid>
        </Grid>
          
      </div>
    </div>
  );
}

export default Offers;