CREATE TABLE users
(
    id       serial PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    password varchar(255)        NOT NULL
);

CREATE TABLE notes
(
    id      serial PRIMARY KEY,
    title   varchar(255) NOT NULL,
    info    text,
    user_id integer REFERENCES users (id) ON DELETE CASCADE NOT NULL
);
