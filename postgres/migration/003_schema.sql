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

INSERT INTO offers
(
    hotel_id,
    is_active,
    offer_title,
    cost_per_child,
    cost_per_adult,
    max_guests,
    description,
    is_deleted
)
VALUES
(
    1,
    true,
    'Sample offer',
    20.00,
    35.50,
    5,
    'Sample description',
    false
),
(
    1,
    true,
    'Sample rich offer',
    500.00,
    500.50,
    4,
    'Sample rich description',
    false
),
(
    1,
    true,
    'Sample big offer',
    50.00,
    20.50,
    40,
    'Sample rich description',
    false
),
(
    1,
    false,
    'Sample inactive offer',
    500.00,
    500.50,
    4,
    'If you can see this offer in listing, contact Michal',
    false
)
