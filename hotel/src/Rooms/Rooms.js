import React, { useEffect, useState } from 'react';
import "./Rooms.css"
import GridList from '@material-ui/core/GridList';
import { Button, Grid, makeStyles, TextField, Typography } from '@material-ui/core';
import { TryAddHotelRoom, TryGetHotelRooms } from './FetchUtils';
import { useHistory } from 'react-router';
const useStyles = makeStyles((theme) => ({
  AddNewRoomButton:{
    backgroundColor:'#ffcc80', 
    color:'white',
    '&:hover': {
      background: "#ffcc80",
    },
  }
}));
function Rooms() {
  const history = useHistory();

  const [roomList,setRoomList] = useState([]);
  const [searchRoomNumberText,setSearchRoomNumberText] = useState("");
  const classes = useStyles();

  useEffect(() =>{
    TryGetHotelRooms(searchRoomNumberText == "" ? null : searchRoomNumberText).then(function (response) {
      if(response !== ""){

        setRoomList(response)
      }else{setRoomList([])}
    })
  },[searchRoomNumberText])
  function OnClickAddNewRoomButton() {
    TryAddHotelRoom(searchRoomNumberText).then(function (response){
      setSearchRoomNumberText("")
    })
  }
  function OnClickOnRoom(hotelRoomNumber){
    history.push(`/rooms/${hotelRoomNumber}`)
  }
  return (
    <div className='rootRooms'>
      <div className='roomsTitle'>
        <Typography variant="h4">Rooms</Typography>
      </div>

      <div  className="roomsButtons">
        <TextField className="roomsButton" label='Room number' size='small' variant='outlined' value={searchRoomNumberText} onChange={(e) => setSearchRoomNumberText(e.target.value)}/>
        {roomList.length == 0 && searchRoomNumberText != "" ?  
        <Button className={classes.AddNewRoomButton} onClick={()=>OnClickAddNewRoomButton()}>Add new room</Button>
        :
        <></>
        }
      </div>

      <Grid className="roomList" >
        {roomList.map((room)=>(
          <Grid className="roomListItem" onClick={()=>{OnClickOnRoom(room.hotelRoomNumber)}} key={room.roomID} >
            <Typography>{room.hotelRoomNumber}</Typography>
          </Grid>
        ))
        }
      </Grid>
      
    </div>
  );
}

export default Rooms;