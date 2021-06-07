import { Box, Button, Grid, makeStyles, Typography } from "@material-ui/core";
import React from "react";
import { Link } from "react-router-dom";
import "./Reservations.css";

const useStyles = makeStyles((theme) => ({
  title:{
    marginTop:10,
    fontSize: 25
  },
}));

function Reservations() {
  const classes = useStyles();

  return (
    <>

      <Grid className="Reservations">
        <Typography className={classes.title}>
          Reservations
        </Typography>
        <Box className="ReservationRow">
          <Box>
            <Typography>Date from:[Date]</Typography>
            <Typography>Date to:[Date]</Typography>
          </Box>
          <Box>
            <Typography>Count of adults:[count]</Typography>
            <Typography>Count of children:[count]</Typography>
          </Box>
          <Box>
            <Typography>Room numbers:</Typography>
            <Typography>[room number list]</Typography>
          </Box>
        </Box>
        <Box className="ReservationRow">
          <Box>
            <Typography>Date from:[Date]</Typography>
            <Typography>Date to:[Date]</Typography>
          </Box>
          <Box>
            <Typography>Count of adults:[count]</Typography>
            <Typography>Count of children:[count]</Typography>
          </Box>
          <Box>
            <Typography>Room numbers:</Typography>
            <Typography>[room number list]</Typography>
          </Box>
        </Box>
        <Box className="ReservationRow">
          <Box>
            <Typography>Date from:[Date]</Typography>
            <Typography>Date to:[Date]</Typography>
          </Box>
          <Box>
            <Typography>Count of adults:[count]</Typography>
            <Typography>Count of children:[count]</Typography>
          </Box>
          <Box>
            <Typography>Room numbers:</Typography>
            <Typography>[room number list]</Typography>
          </Box>
        </Box>
        <Box className="ReservationRow">
          <Box>
            <Typography>Date from:[Date]</Typography>
            <Typography>Date to:[Date]</Typography>
          </Box>
          <Box>
            <Typography>Count of adults:[count]</Typography>
            <Typography>Count of children:[count]</Typography>
          </Box>
          <Box>
            <Typography>Room numbers:</Typography>
            <Typography>[room number list]</Typography>
          </Box>
        </Box>
      </Grid>
    </>
  );
}

export default Reservations;
