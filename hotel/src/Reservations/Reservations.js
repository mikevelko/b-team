import { Box, Button, Grid, makeStyles, Typography } from "@material-ui/core";
import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { TryGetReservations } from "../Utils/FetchUtils";
import "./Reservations.css";

const useStyles = makeStyles((theme) => ({
  title:{
    marginTop:10,
    fontSize: 25
  },
}));

function Reservations() {
  const classes = useStyles();

  const [reservations, setReservations] = useState([])
  useEffect(()=>{
    TryGetReservations().then(function (response) {
      if(response.length != 0){
        setReservations(response)
      }
    })
  },[])
  return (
    <>

      <Grid className="Reservations">
        <Typography className={classes.title}>
          Reservations
        </Typography>
        {reservations.map((reservation) => {return (
          <Box key={reservation.reservation.reservationID} className="ReservationRow">
            <Box>
              <Typography>Name:{reservation.client.name}</Typography>
              <Typography>Surname:{reservation.client.surname}</Typography>
            </Box>
            <Box>
              <Typography>Date from:{reservation.reservation.from}</Typography>
              <Typography>Date to:{reservation.reservation.to}</Typography>
            </Box>
            <Box>
              <Typography>Count of adults:{reservation.reservation.numberOfAdults}</Typography>
              <Typography>Count of children:{reservation.reservation.numberOfChildren}</Typography>
            </Box>
            <Box>
              <Typography>Room number:{reservation.room.hotelRoomNumber}</Typography>
            </Box>
          </Box>
        )})}
      </Grid>
    </>
  );
}

export default Reservations;
