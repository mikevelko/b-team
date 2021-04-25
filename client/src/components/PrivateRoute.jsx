import React, { Component, useEffect } from 'react';
import { BrowserRouter as Router, Route, Switch, Redirect } from 'react-router-dom';

function PrivateRoute ({component: Component, authed, ...rest}) {
    return (
      <Route {...rest}
        render={(props) => authed 
          ? <Component {...props} {...rest}/>
          : <Redirect to={{pathname: '/login'}} />}
      />
    )

    
  }

export default PrivateRoute;