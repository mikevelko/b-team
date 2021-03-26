import React, { useEffect } from 'react';
import PropTypes from 'prop-types';


function OfferDetails({match}) {
    useEffect(() =>{
        console.log(match)
    },[]);
  return (
    <h1>Offer Details {match.params.offerId}</h1>
  );
}

OfferDetails.propTypes = {
    match: PropTypes.shape({
        params:PropTypes.shape({
            offerId : PropTypes.string
        }).isRequired
    })
};

export default OfferDetails;