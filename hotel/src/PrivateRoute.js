import React from "react";
import { Redirect, Route } from "react-router";

export function PrivateRoute ({component: Component}) {
    function isAuthorised(){
        try{
            const token = JSON.parse(localStorage.getItem("x-hotel-token"));
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