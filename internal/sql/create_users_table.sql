CREATE TABLE users (
    id serial4 NOT NULL,
    email text NOT NULL DEFAULT ''::text unique,
    "name" text NOT NULL DEFAULT ''::text unique,
    is_deleted bool NOT NULL DEFAULT false,
    is_registered bool NOT NULL DEFAULT false,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);