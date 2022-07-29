CREATE TABLE users
(
    id       serial primary key,
    username varchar(255) not null unique,
    password varchar(255) not null
);