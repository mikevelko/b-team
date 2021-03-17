CREATE SCHEMA rently
    CREATE TABLE rently.offers
    (
        id                        bigserial primary key,
        is_active                 bool    not null,
        offer_title               text    not null,
        cost_per_child            decimal not null,
        cost_par_adult            decimal not null,
        max_guests                integer not null,
        description               text    not null,
        offer_preview_picture_url text
    );
