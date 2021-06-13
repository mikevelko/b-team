INSERT INTO users(
    first_name,
    surname,
    email,
    user_name,
    password,
    hotel_id,
    user_role
)
VALUES
(
    'client',
    'client',
    'client@client.client',
    'client',
    'client',
    null,
    'CLIENT_CUSTOMER'
),
(
    'hotel',
    'hotel',
    'hotel@hotel.hotel',
    'hotel',
    'hotel',
    1,
    'HOTEL_MANAGER'
);

INSERT INTO hotels(
    id,
    hotel_name,
    hotel_desc,
    city,
    country
)
VALUES
(
    1,
    'hotel',
    'Hotel for testing hotel operations',
    'Warsaw',
    'Poland'
);




INSERT INTO users(
    id,
    first_name,
    surname,
    email,
    user_name,
    password,
    hotel_id,
    user_role
)
VALUES
(
    99999998,
    'adminC',
    'adminC',
    'adminC@adminC.adminC',
    'adminC',
    'adminC',
    null,
    'CLIENT_CUSTOMER'
),
(
    99999999,
    'adminH',
    'adminH',
    'adminH@adminH.adminH',
    'adminH',
    'adminH',
    99999999,
    'HOTEL_MANAGER'
),
(
    99999990,
    'client0',
    'client0',
    'client0@client.client',
    'client0',
    'client0',
    null,
    'CLIENT_CUSTOMER'
),
(
    99999991,
    'client1',
    'client1',
    'client1@client.client',
    'client1',
    'client1',
    null,
    'CLIENT_CUSTOMER'
),
(
    99999992,
    'client2',
    'client2',
    'client2@client.client',
    'client2',
    'client2',
    null,
    'CLIENT_CUSTOMER'
);



INSERT INTO sessions(
    id,
    creation_date,
    expire_date,
    user_id
)
VALUES
(
    99999998,
    NOW(),
    '2050-01-01',
    99999998
),
(
    99999999,
    NOW(),
    '2050-01-01',
    99999999
);
