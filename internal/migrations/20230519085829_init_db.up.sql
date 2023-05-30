create table if not exists "home" (
    id serial primary key,
    type varchar(10) not null,
    address varchar(511) not null,
    formatted_address varchar(511) not null,
    address_line varchar(511) not null,
    postal_code varchar(31) not null,
    longitude float not null,
    latitude float not null,
    country varchar(100) not null,
    regency varchar(100) not null,
    subdistrict varchar(100) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz,
    deleted_at timestamptz
);