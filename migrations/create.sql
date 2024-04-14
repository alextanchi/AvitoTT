CREATE TABLE feature
(
    id int NOT NULL,
    CONSTRAINT feature_pkey PRIMARY KEY (id)

);

CREATE TABLE banner
(
    id         serial,
    title      varchar                  NOT NULL,
    text       varchar                  NOT NULL,
    url        varchar                  NOT NULL,
    is_active  bool                     NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    feature_id int                      NOT NULL,
    CONSTRAINT banner_pkey PRIMARY KEY (id),
    CONSTRAINT feature_fk
        FOREIGN KEY (feature_id)
            REFERENCES feature (id)

);

CREATE TABLE tag
(
    id        int NOT NULL,
    banner_id int NOT NULL,
    CONSTRAINT tag_fk
        FOREIGN KEY (banner_id)
            REFERENCES banner (id)
            ON DELETE CASCADE
);

CREATE TABLE role
(
    user_id uuid    NOT NULL,
    role    varchar NOT NULL

);

INSERT INTO feature (id)
VALUES (1),(2),(3)
        ;

INSERT INTO role (user_id, role)
VALUES ('3b1b19f3-455d-474d-92fc-65a76551b16f', 'user'),
    ('0d97e23b-089b-4237-9c2c-d1e100576920', 'admin')
;