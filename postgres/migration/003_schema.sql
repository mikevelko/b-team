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

INSERT INTO hotels(
    id,
    hotel_name,
    hotel_desc,
    city,
    country
)
VALUES
(
    99999999,
    '999_999_99 hotel',
    'Hotel for testing hotel 999',
    'Moscow',
    'Russia'
),
(
    99999998,
    'Grand Arizona Grounds Hotel',
    'See the Grand Canyon or something, idk',
    'Berlin',
    'Germany'
),
(
    99999997,
    'Grand Budapest Hotel',
    'Quite good movie btw',
    'Budapest',
    'Hungary'
);