import React, { useState, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import './Offer.css';
import { makeStyles } from '@material-ui/core/styles';
import GridList from '@material-ui/core/GridList';
import GridListTile from '@material-ui/core/GridListTile';
import GridListTileBar from '@material-ui/core/GridListTileBar';
import TextField from '@material-ui/core/TextField';
import axios from 'axios';
import { Link, useHistory } from 'react-router-dom';

function Offer(props) {
    let hotelId = props.match.params.hotelId;
    let offerId = props.match.params.offerId;
    const [offer, setOffer] = useState([]);
    const history = useHistory();

    const [adults, setAdults] = useState(1);
    const [children, setChildren] = useState(0);
    const [reservationPossible, setReservationPossible] = useState(true);

    const handleChangeAdults = (event) => {
        if((parseInt(event.target.value) + parseInt(children)) > offer.maxGuests) setReservationPossible(false);
        else setReservationPossible(true);
        if(event.target.value < 0)
        {
            setAdults(0);
            if(parseInt(children) > offer.maxGuests) setReservationPossible(false);
            else setReservationPossible(true);
            return;
        }
        if(event.target.value === "" || children === "") setReservationPossible(false);
        setAdults(event.target.value);
    };
    const handleChangeChildren = (event) => {
        if((parseInt(event.target.value) + parseInt(adults)) > offer.maxGuests) setReservationPossible(false);
        else setReservationPossible(true);
        if(event.target.value < 0)
        {
            setChildren(0);
            if(parseInt(adults) > offer.maxGuests) setReservationPossible(false);
            else setReservationPossible(true);
            return;
        }
        if(event.target.value === "" || adults === "") setReservationPossible(false);
        setChildren(event.target.value);

    };




    useEffect(() => {
        fetchItems();
    }, []);

    const tileData = [
        {
            img: "https://images.all-free-download.com/images/graphiclarge/simple_room_picture_167607.jpg",
            title: "hotel",
            cols: 2
        },
        {
            img: "https://images.all-free-download.com/images/graphiclarge/simple_room_picture_167607.jpg",
            title: "",
            cols: 1
        },
        {
            img: "https://images.all-free-download.com/images/graphiclarge/simple_room_picture_167607.jpg",
            title: "",
            cols: 1
        },
        {
            img: "https://images.all-free-download.com/images/graphiclarge/simple_room_picture_167607.jpg",
            title: "",
            cols: 1
        },
        {
            img: "https://images.all-free-download.com/images/graphiclarge/simple_room_picture_167607.jpg",
            title: "",
            cols: 1
        },
        {
            img: "https://images.all-free-download.com/images/graphiclarge/simple_room_picture_167607.jpg",
            title: "",
            cols: 1
        }
    ];


    const fetchItems = () => {
        const url = `/api-client/hotels/${hotelId}/offers/${offerId}`;
        axios.get(url, { headers: { 'accept': 'application/json', 'x-session-token': window.localStorage.getItem("token") } })
            .then(response => {
                console.log(response.data);
                setOffer(response.data);
            })
            .catch(error => {
                // if hotel not exist then redirect to hotels page
                if (error.response.status === 404) {
                    let path = `/hotels/${hotelId}/offers`;
                    history.push(path);
                }
                console.error('There was an error!', error.response);
            });
    }

    

    const Reserve = () => {
        const body = 
        {
            from: "2021-10-02T10:00:00-05:00",
            to: "2021-12-02T10:00:00-05:00",
            numberOfChildren: children,
            numberOfAdults: adults
        };

        const url = `/api-client/hotels/${hotelId}/offers/${offerId}/reservations`;
        axios.post(url,body, { headers: { 'accept': 'application/json', 'x-session-token': window.localStorage.getItem("token"), 'Content-Type': 'application/json' } })
            .then(response => {
                console.log(response.data);
                fetchItems();
            })
            .catch(error => {
                // if hotel not exist then redirect to hotels page
                if (error.response.status === 404) {
                    let path = `/hotels/${hotelId}/offers`;
                    history.push(path);
                }
                console.error('There was an error!', error.response);
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

    const CheckReviews = () => {
        let path = `/hotels/${hotelId}/offers/${offerId}/reviews`;
        history.push(path);
    }


    const classes = useStyles();


    return (
            <div className="offer-container">
                <div className="offer-container-item">
                    <p>{offer.offerTitle}</p>
                    <p>{offer.offerDescription}</p>
                    <p>Max guests: {offer.maxGuests}</p>
                    <p>Price per adult: {offer.costPerAdult}€</p>
                    <p>Price per child: {offer.costPerChild}€</p>
                    <Button variant="contained" color="primary" size="large" onClick={CheckReviews}>Check reviews</Button>
                    <div className={classes.root}>
                        <GridList className={classes.gridList} cols={3} cellHeight='300'>
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
                    <TextField id="outlined-basic" label="Adults" variant="outlined" type="number" InputProps={{ inputProps: { min: 1, max: 10 } }} size="small" value={adults} onChange={handleChangeAdults} />
                    <p></p>
                    <TextField id="outlined-basic" label="Children" variant="outlined" type="number" InputProps={{ inputProps: { min: 1, max: 10 } }} size="small" value={children} onChange={handleChangeChildren} />
                    <p>Total: {offer.costPerAdult * adults + offer.costPerChild * children}€ per night</p>
                    <Button variant="contained" color="primary" size="large" disabled={!reservationPossible || !offer.isActive} onClick={Reserve}>Reserve</Button>
                </div>
            </div>

    );
}

export default Offer;