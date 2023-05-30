create table if not exists "home" (
    id serial primary key,
    type varchar(10) not null,
    address varchar(255) not null,
    longitude numeric(20, 2) not null,
    latitude numeric(20, 2) not null,
    country varchar(100) not null,
    regency varchar(100) not null,
    subdistrict varchar(100) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz,
    deleted_at timestamptz
);