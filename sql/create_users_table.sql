CREATE TABLE users (
                              id serial4 NOT NULL,
                              email text NOT NULL DEFAULT ''::text,
                              "name" text NOT NULL DEFAULT ''::text,
                              is_deleted bool NOT NULL DEFAULT false,
                              CONSTRAINT users_pkey PRIMARY KEY (id)
);