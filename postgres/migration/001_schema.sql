-- clear everything before setting up tables
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;

CREATE TABLE IF NOT EXISTS offers
(
    id                        bigserial primary key,
    hotel_id                  bigint not null,
    is_active                 bool    not null,
    offer_title               text    not null,
    cost_per_child            decimal not null,
    cost_per_adult            decimal not null,
    max_guests                integer not null,
    description               text    not null,
    offer_preview_picture_url text,
    is_deleted                bool DEFAULT false
);

CREATE TABLE IF NOT EXISTS rooms
(
    id                          bigserial primary key,
    room_number                 text not null,
    hotel_id                    bigint not null
);

CREATE TABLE IF NOT EXISTS offers_rooms
(
    offer_id                    bigint not null,
    room_id                     bigint not null
);


CREATE TABLE IF NOT EXISTS hotels
(
    id                          bigserial primary key,
    api_token                   text,
    hotel_name                  text not null,
    hotel_desc                  text,
    city                        text not null,
    country                     text not null,
    preview_picture_url         text
);


CREATE TABLE IF NOT EXISTS users
(
    id                          bigserial primary key,
    first_name                  text not null,
    surname                     text not null,
    email                       text not null,
    user_name                   text not null,
    password                    text not null,
    hotel_id                    bigint,
    user_role                   text DEFAULT 'CLIENT_CUSTOMER' not null
);


CREATE TABLE IF NOT EXISTS sessions
(
    id                          bigserial primary key,
    creation_date               timestamptz not null,
    expire_date                 timestamptz not null,
    user_id                     bigint not null
);






