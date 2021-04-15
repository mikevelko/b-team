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
    api_token
)
VALUES
(
    99999999,
    '{"ID": 99999999, "CreatedAt":"2020-01-01"}'
)