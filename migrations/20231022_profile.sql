-- +goose Up
CREATE TABLE profile (
                      id uuid NOT NULL,
                      email varchar(255),
                      phone decimal(11, 0),
                      password text,
                      surname varchar(255),
                      name varchar(255),
                      patronymic varchar(255),
                      create_at timestamp DEFAULT CURRENT_TIMESTAMP,
                      update_at timestamp DEFAULT CURRENT_TIMESTAMP,
                      PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE profile;