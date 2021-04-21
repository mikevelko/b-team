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

  export async function TryPatchHotelInfo(hotelName,hotelDesc,city,country,hotelPreviewPucture='',hotelPictures=[]){
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