CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(256) NOT NULL,
    password VARCHAR(256) NOT NULL
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    sender INTEGER REFERENCES users,
    time_sent TIMESTAMP (6) WITH TIME ZONE,
    channel SMALLINT NOT NULL CHECK ( channel between 0 and 9),
    content VARCHAR(1024) NOT NULL
);

grant INSERT, SELECT, UPDATE, DELETE ON users TO docker;
grant INSERT, SELECT, UPDATE, DELETE ON messages TO docker;
grant INSERT, SELECT, UPDATE, DELETE ON users_id_seq TO docker;
grant INSERT, SELECT, UPDATE, DELETE ON messages_id_seq TO docker;
