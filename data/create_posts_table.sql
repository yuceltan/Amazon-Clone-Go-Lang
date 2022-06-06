CREATE TABLE IF NOT EXISTS posts
(
    id SERIAL NOT NULL,
    title character varying(255)  NOT NULL,
    body character varying(550)  NOT NULL,
    owner_id integer  REFERENCES users (id),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    CONSTRAINT posts_pkey PRIMARY KEY (id)
);