import React, { Component, useEffect } from 'react';
import { BrowserRouter as Router, Route, Switch, Redirect } from 'react-router-dom';

function PrivateRoute ({component: Component, authed}) {
    return (
      <Route
        render={(props) => authed === true
          ? <Component/>
          : <Redirect to={{pathname: '/login'}} />}
      />
    )
  }

export default PrivateRoute;