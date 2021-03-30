CREATE SCHEMA bookly
    CREATE TABLE bookly.offers
    (
        id                        bigserial primary key,
        hotel_id                  bigint not null,
        is_active                 bool    not null,
        offer_title               text    not null,
        cost_per_child            decimal not null,
        cost_par_adult            decimal not null,
        max_guests                integer not null,
        description               text    not null,
        offer_preview_picture_url text
    );

