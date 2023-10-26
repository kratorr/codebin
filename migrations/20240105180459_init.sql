-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(250) UNIQUE,
    username VARCHAR(250) UNIQUE,
    password VARCHAR(250),
    avatar VARCHAR(250),
    google_id VARCHAR(250) UNIQUE,
    github_id VARCHAR(250) UNIQUE,
    activation_status INT
);

CREATE TABLE IF NOT EXISTS languages(
    id SERIAL PRIMARY KEY,
    name VARCHAR(250) UNIQUE
);


CREATE TABLE IF NOT EXISTS snippets(
    id SERIAL PRIMARY KEY,
    title VARCHAR(250),
    text TEXT,
    language_id INT NOT NULL,
    user_id INT,
    url VARCHAR(10) UNIQUE,
    created_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (language_id) REFERENCES languages(id)
);


-- +goose Down
DROP TABLE snippets;
DROP TABLE languages;
DROP TABLE users;
