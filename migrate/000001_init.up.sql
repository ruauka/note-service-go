CREATE TABLE IF NOT EXISTS users
(
    id       serial PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    password varchar(255)        NOT NULL
);

CREATE TABLE IF NOT EXISTS notes
(
    id      serial PRIMARY KEY,
    title   varchar(255) UNIQUE                             NOT NULL,
    info    text,
    user_id integer REFERENCES users (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS tags
(
    id      serial PRIMARY KEY,
    tagname varchar(255)                                    NOT NULL,
    user_id integer REFERENCES users (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS notes_tags
(
    note_id integer REFERENCES notes (id) ON DELETE CASCADE,
    tag_id  integer REFERENCES tags (id) ON DELETE CASCADE,

    CONSTRAINT notes_tags_pkey PRIMARY KEY (note_id, tag_id)
);
