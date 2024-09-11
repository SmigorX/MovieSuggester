CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    director VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    tags TEXT[],
    watched BOOLEAN NOT NULL DEFAULT FALSE
);
