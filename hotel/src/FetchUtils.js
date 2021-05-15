import axios from 'axios';

export let HOTEL_TOKEN_NAME = 'x-hotel-token'

export async function TryLogIn(login, password){
    const res = await axios({
      method: 'post',
      url: '/api-client/client/login',
      headers: {
        'accept': 'application/json',
        'Content-Type': 'application/json',
      },
      data: {
        "login": login,
        "password": password
      }, 
    })
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
    });
    if(res !== undefined) return JSON.stringify(res.data);
    return "";
  };


  export async function TryGetHotelInfo(){
    const res = await axios({
      method: 'get',
      url: '/api-hotel/hotelInfo',
      headers: {
        'accept': 'application/json',
        'x-hotel-token': localStorage.getItem(HOTEL_TOKEN_NAME)
      }, 
    })
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
    });

    if(res !== undefined) return res.data;
    return "";
  };

  export async function TryPatchHotelInfo(hotelName,hotelDesc,hotelPreviewPucture='',hotelPictures=[]){
    const res = await axios({
      method: 'PATCH',
      url: '/api-hotel/hotelInfo',
      headers: {
        'accept': 'application/json',
        'x-hotel-token': localStorage.getItem(HOTEL_TOKEN_NAME),
        'Content-Type': 'application/json',
      }, 
      data:{
        "hotelName": hotelName,
        "hotelDesc": hotelDesc,
        "hotelPreviewPicture": hotelPreviewPucture,
        "hotelPictures": hotelPictures,
        // "city":city,   has no added on the back 
        // "country":country,   has no added on the back 
      },
    })
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error)
    });

    if(res !== undefined) return res;
    return "";
  };

  export async function TryPostOffer(offerTitle,costPerChild,costPerAdult,maxGuests,activeStatus,rooms,description,pictures=[],previewPicture=''){
    const res = await axios({
      method: 'post',
      url: '/api-hotel/offers',
      headers: {
        'accept': 'application/json',
        'Content-Type': 'application/json',
        'x-hotel-token': localStorage.getItem(HOTEL_TOKEN_NAME),
      },
      data: {
        "isActive": activeStatus,
        "offerTitle": offerTitle,
        "costPerChild": costPerChild,
        "costPerAdult": costPerAdult,
        "maxGuests": maxGuests,
        "description": description,
        "offerPreviewPicture": previewPicture,
        "pictures": pictures,
        "rooms": rooms
      }, 
    })
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
    });
    if(res !== undefined) return res.data.offerID;
    return -1;
  };

  export async function TryGetHotelOffers(pageNumber = 1, pageSize = 10,isActive = null){
    const res = await axios({
      method: 'get',
      url: '/api-hotel/offers',
      headers: {
        'accept': 'application/json',
        'x-hotel-token': localStorage.getItem(HOTEL_TOKEN_NAME)
      }, 
      params:{
        isActive,
        pageNumber,
        pageSize,
      },
    })
    .then(function (response) {
      return response;
    })
    .catch(function (error) {
      console.log(error);
    });
    if(res !== undefined) return res.data.offerPreview;
    return "";
  };

  export async function TryGetHotelOffer(offerID){
    const res = await axios({
      method: 'get',
      url: '/api-hotel/offers/'+offerID,
      headers: {
        'accept': 'application/json',
        'x-hotel-token': localStorage.getItem(HOTEL_TOKEN_NAME)
      }, 
    })
    .then(function (response) {
      console.log(response)
      return response;
    })
    .catch(function (error) {
      console.log(error);
    });
    if(res !== undefined) return res.data;
    return "";
  };

  export async function TryDeleteHotelOffer(offerID){
    const res = await axios({
      method: 'delete',
      url: '/api-hotel/offers/'+offerID,
      headers: {
        'accept': 'application/json',
        'x-hotel-token': localStorage.getItem(HOTEL_TOKEN_NAME)
      }, 
    })
    .then(function (response) {
      console.log(response)
      return response;
    })
    .catch(function (error) {
      console.log(error);
    });
    if(res !== undefined) return res.data;
    return "";
  };
  export async function TryEditHotelOffer(offerID,offerTitle,maxGuests,activeStatus,description,pictures=[],previewPicture=''){
    const res = await axios({
      method: 'PATCH',
      url: '/api-hotel/offers/'+offerID,
      headers: {
        'accept': 'application/json',
        'x-hotel-token': localStorage.getItem(HOTEL_TOKEN_NAME),
        'Content-Type': 'application/json',
      }, 
      data:{
        "isActive": activeStatus,
        "offerTitle": offerTitle,
        "maxGuests": maxGuests,
        "description": description,
        "offerPreviewPicture": previewPicture,
        "offerPictures": pictures,
      }
    })
    .then(function (response) {
      console.log(response)
      return response;
    })
    .catch(function (error) {
      console.log(error);
    });
    if(res !== undefined) return res.status;
    return "";
  };