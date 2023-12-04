CREATE TABLE user (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255),
    email       VARCHAR(255),
    mobile      VARCHAR(255),
    password    VARCHAR(255),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP NULL
);