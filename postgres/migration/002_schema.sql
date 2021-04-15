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
    api_token
)
VALUES
(
    '{"ID": 1, "CreatedAt":"2020-01-01"}'
)