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
    offer_preview_picture_url,
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
    'qwehasihdl.cop',
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
    'qwehasihdl.cop',
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
    'qwehasihdl.cop',
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
    'qwehasihdl.cop',
    false
);

INSERT INTO rooms
(
    hotel_id,
    room_number
)
VALUES
(
    1,
    '12F'
),
(
    1,
    '14A'
),
(
    1,
    '16F'
),
(
    1,
    '112'
);

INSERT INTO offers_rooms
(
    room_id,
    offer_id
)
VALUES
(1,1),(2,1),(3,2),(4,2);