CREATE TABLE users (
    id serial4 NOT NULL,
    is_deleted bool NOT NULL DEFAULT false,
    email text NOT NULL DEFAULT ''::text,
    is_registered bool NOT NULL DEFAULT false,
    "name" text NOT NULL DEFAULT ''::text,
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_pkey PRIMARY KEY (id)
);