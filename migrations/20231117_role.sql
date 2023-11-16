-- +goose Up
CREATE TABLE role (
                         id uuid NOT NULL,
                         name varchar(255),
                         PRIMARY KEY(id)
);

INSERT INTO role (id, name) VALUES
    ('7647316e-22e2-4f94-93bc-c0459dcb54de', 'Owner'),
    ('6f306b39-0ef6-4dc1-bab4-a050a3ce6f8c', 'Editor');

-- +goose Down
DROP TABLE role;