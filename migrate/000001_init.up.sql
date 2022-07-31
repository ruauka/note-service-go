CREATE TABLE users
(
    id       serial PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    password varchar(255)        NOT NULL
);

CREATE TABLE notes
(
    id        SERIAL PRIMARY KEY,
    note      text,
    author_id int references users (id) on DELETE CASCADE NOT NULL
);
