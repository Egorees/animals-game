CREATE TABLE IF NOT EXISTS users
(
    id serial not null unique,
    telegram_id bigint unique,
    username varchar(255),
    password_hash varchar(255),
    animal_id integer
);

CREATE TABLE IF NOT EXISTS animals
(
    id serial not null unique,
    name varchar(255),
    type smallint not null,
    exp bigint,
    owner_id integer not null,
    owner_telegram_id bigint not null
);