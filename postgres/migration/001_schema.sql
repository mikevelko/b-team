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
    offer_preview_picture_url text
);

CREATE TABLE IF NOT EXISTS hotels
(
    id                          bigserial primary key,
    api_token                   text not null
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
    user_role                   text DEFAULT 'CLIENT_CUSTOMER'
);

CREATE TABLE IF NOT EXISTS sessions
(
    id                          bigserial primary key,
    creation_date               timestamptz not null,
    expire_date                 timestamptz not null,
    user_id                     bigint not null
);

INSERT INTO sessions(
    id,
    creation_date,
    expire_date,
    user_id
)
VALUES(99999999,NOW(),'2050-01-01',1);





