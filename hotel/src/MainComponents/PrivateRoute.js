import React from "react";
import { Redirect, Route } from "react-router";
import {HOTEL_TOKEN_NAME} from '../Utils/FetchUtils.js';
export function PrivateRoute ({component: Component}) {
    function isAuthorised(){
        try{
            const token = JSON.parse(localStorage.getItem(HOTEL_TOKEN_NAME));
            return token !== null; 
        }catch(error){
            return false
        }
    }
    return (
      <Route
        render={(props) => isAuthorised()
          ? <Component {...props} />
          : <Redirect to={{pathname: '/login'}} />}
      />
    )
  }