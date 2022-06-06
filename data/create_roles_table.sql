CREATE TABLE IF NOT EXISTS roles
(
    name character varying(255)  NOT NULL,
    CONSTRAINT roles_pkey PRIMARY KEY (name)
);

CREATE TABLE IF NOT EXISTS users_roles
(
    user_id integer  REFERENCES users (id),
    name character varying(255)  REFERENCES roles (name),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    CONSTRAINT users_roles_pkey PRIMARY KEY (user_id, name)
);


INSERT INTO roles (name) VALUES ('ADMIN');
INSERT INTO roles (name) VALUES ('EDITOR');


INSERT INTO users_roles (user_id, name, created_at, updated_at) VALUES ('9','EDITOR','2022-03-31 00:00:00', '2022-03-31 00:00:00');
INSERT INTO users_roles (user_id, name, created_at, updated_at) VALUES ('8','ADMIN','2022-03-31 00:00:00', '2022-03-31 00:00:00');






