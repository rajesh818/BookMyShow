CREATE TABLE user (
    id          INT PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(255),
    email       VARCHAR(255),
    mobile      VARCHAR(255),
    password    VARCHAR(255),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP NULL
);

CREATE TABLE location (
    id          INT PRIMARY KEY AUTO_INCREMENT,
    city        VARCHAR(255),
    state       VARCHAR(255),
    pincode      INT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP NULL
);


CREATE TABLE theatre (
    id          INT PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(255),
    owner_email       VARCHAR(255),
    owner_mobile      VARCHAR(255),
    location_id INT,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP NULL,
    FOREIGN KEY(location_id) REFERENCES location (id)
);


CREATE TABLE movies_available (
    id          INT PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(255),
    actor       VARCHAR(255),
    languaue      VARCHAR(255),
    director VARCHAR(255),
    active BOOLEAN,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP NULL
);

CREATE TABLE running_movies (
    id int PRIMARY key AUTO_INCREMENT,
    movie_id int,
    theater_id int,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP NULL,
    FOREIGN key (movie_id) REFERENCES  movies_available (id)
    FOREIGN key (theater_id) REFERENCES  theatre (id)
)