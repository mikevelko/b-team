import React, { useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import './Hotel.css';
import { makeStyles } from '@material-ui/core/styles';
import GridList from '@material-ui/core/GridList';
import GridListTile from '@material-ui/core/GridListTile';
import GridListTileBar from '@material-ui/core/GridListTileBar';
import IconButton from '@material-ui/core/IconButton';
import StarBorderIcon from '@material-ui/icons/StarBorder';
import axios from 'axios';

function Hotel(props) {
    let hotelId = props.match.params.hotelId;
    const [hotel, setHotel] = useState([]);


    useEffect(() => {
        fetchItems();
    }, []);

    const tileData = [
        {
            img: "https://pix10.agoda.net/hotelImages/124/1246280/1246280_16061017110043391702.jpg?s=1024x768",
            title: "hotel",
            cols: 2
        },
        {
            img: "https://pix10.agoda.net/hotelImages/124/1246280/1246280_16061017110043391702.jpg?s=1024x768",
            title: "",
            cols: 1
        },
        {
            img: "https://pix10.agoda.net/hotelImages/124/1246280/1246280_16061017110043391702.jpg?s=1024x768",
            title: "",
            cols: 1
        },
        {
            img: "https://pix10.agoda.net/hotelImages/124/1246280/1246280_16061017110043391702.jpg?s=1024x768",
            title: "",
            cols: 1
        },
        {
            img: "https://pix10.agoda.net/hotelImages/124/1246280/1246280_16061017110043391702.jpg?s=1024x768",
            title: "",
            cols: 1
        },
        {
            img: "https://pix10.agoda.net/hotelImages/124/1246280/1246280_16061017110043391702.jpg?s=1024x768",
            title: "",
            cols: 1
        }
    ];


    const fetchItems = () => {
        const url = '/api-client/hotels/' + hotelId.toString();
        axios.get(url, { headers: { 'accept': 'application/json', 'x-session-token': window.localStorage.getItem("token") } })
            .then(response => {
                console.log(response);
                setHotel(response.data);
            })
            .catch(error => {
                console.error('There was an error!', error);
            });


    }


    const useStyles = makeStyles((theme) => ({
        root: {
            display: 'flex',
            flexWrap: 'wrap',
            justifyContent: 'space-around',
            overflow: 'hidden',
            backgroundColor: theme.palette.background.paper,
            margin: '50px',
        },
        gridList: {
            flexWrap: 'nowrap',
            // Promote the list into his own layer on Chrome. This cost memory but helps keeping high FPS.
            transform: 'translateZ(0)',
        },
        title: {
            color: theme.palette.primary.light,
        },
        titleBar: {
            background:
                'linear-gradient(to top, rgba(0,0,0,0.7) 0%, rgba(0,0,0,0.3) 70%, rgba(0,0,0,0) 100%)',
        },
    }));



    const classes = useStyles();
    return (

        <div className="hotel-container">
            <div className="hotel-container-item">
                <p>{hotel.hotelName}</p>
                <p>{hotel.hotelDesc}</p>
                <p>{hotel.city}, {hotel.country}</p>
                <Button variant="contained" color="primary" style={{ fontSize: '42px', maxWidth: '30%', maxHeight: '100px', minWidth: '30%', minHeight: '100px' }}
                    size="large" >Check offers</Button>
                <div className={classes.root}>
                    <GridList className={classes.gridList} cols={3} cellHeight='auto'>
                        {tileData.map((tile) => (
                            <GridListTile key={tile.img}>
                                    <img src={tile.img} alt={tile.title} />
                                <GridListTileBar
                                    classes={{
                                        root: classes.titleBar,
                                        title: classes.title,
                                    }}
                                />
                            </GridListTile>
                        ))}
                    </GridList>
                </div>
                
            </div>
        </div>

    );
}

export default Hotel;