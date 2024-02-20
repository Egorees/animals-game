CREATE TABLE IF NOT EXISTS users
(
    id serial not null unique,
    username varchar(255) not null unique,
    telegram_id varchar(255) not null,
    password_hash varchar(255) not null,
    animal_id integer
);

CREATE TABLE IF NOT EXISTS animals
(
    id serial not null unique,
    type smallint,
    exp bigint,
    owner_user_id integer
);