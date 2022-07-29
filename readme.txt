brew install golang-migrate - утилита для миграций
migrate create -ext sql -dir ./migrate -seq init - создание файлов миграций
migrate -path ./migrate -database 'postgres://pg:pass@localhost:5432/crud?sslmode=disable' up
migrate -path ./migrate -database 'postgres://pg:pass@localhost:5432/crud?sslmode=disable' down


CREATE TABLE users
(
    id       serial primary key,
    username varchar(255) not null unique,
    password varchar(255) not null
);

CREATE TABLE notes
(
    id      serial primary key,
    note    varchar(255),
    user_id integer REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);

select *
FROM users;

DROP TABLE users CASCADE;

INSERT INTO notes (note, user_id)
VALUES ('note 1', 1);

INSERT INTO notes (note, user_id)
VALUES ('note 2', 1);

select *
FROM notes;

DROP TABLE notes