import React, { Component, useEffect, useState } from 'react';

function Hotels() {
    useEffect(() => {
        fetchItems();
    }, []);

    const [items, setItems] = useState([]);

    const fetchItems = async () => {
        
    }
    
    return (
        <div>
            <h3>Hotels page</h3>
        </div>
    );
}

export default Hotels;