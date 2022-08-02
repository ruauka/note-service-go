CREATE TABLE users
(
    id       serial PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    password varchar(255)        NOT NULL
);

CREATE TABLE notes
(
    id      serial PRIMARY KEY,
    title   varchar(255)                                    NOT NULL,
    info    text,
    user_id integer REFERENCES users (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE tags
(
    id      serial PRIMARY KEY,
    tagname varchar(255) UNIQUE                             NOT NULL,
    user_id integer REFERENCES users (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE notes_tags
(
    note_id integer REFERENCES notes (id),
    tag_id  integer REFERENCES tags (id),

    CONSTRAINT notes_tags_pkey PRIMARY KEY (note_id, tag_id)
);
