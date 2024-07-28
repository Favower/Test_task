CREATE USER myuser WITH PASSWORD 'mypassword';
CREATE DATABASE messages_db;
GRANT ALL PRIVILEGES ON DATABASE messages_db TO myuser;

\connect messages_db
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    processed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
