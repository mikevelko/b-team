import axios from 'axios';


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