CREATE TABLE banner
(
    id         int                      NOT NULL,
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

CREATE TABLE feature
(
    id int NOT NULL,
    CONSTRAINT feature_pkey PRIMARY KEY (id)

);


CREATE TABLE tag
(
    id        int NOT NULL,
    banner_id int NOT NULL,

    CONSTRAINT tag_pkey PRIMARY KEY (id),
    CONSTRAINT tag_fk
        FOREIGN KEY (banner_id)
            REFERENCES banner (id)
);

CREATE TABLE role
(
    user_id uuid    NOT NULL,
    role    varchar NOT NULL

);