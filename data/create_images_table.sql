CREATE TABLE IF NOT EXISTS profile_images
(
    user_id integer  REFERENCES users (id),
    name character varying(255)  NOT NULL,
    value BYTEA NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    CONSTRAINT profile_images_pkey PRIMARY KEY (user_id)
);