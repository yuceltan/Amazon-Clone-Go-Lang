CREATE TABLE IF NOT EXISTS users
(
    id SERIAL NOT NULL,
    first_name character varying(255)  NOT NULL,
    last_name character varying(255)  NOT NULL,
    email character varying(255)  NOT NULL,
    password character varying(255)  NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX IF NOT EXISTS users_email_idx
    ON users USING btree
    (email ASC NULLS LAST);



