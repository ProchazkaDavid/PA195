CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(256) NOT NULL ,
    password VARCHAR(256) NOT NULL
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    sender INTEGER REFERENCES users,
    channel SMALLINT NOT NULL CHECK ( channel between 0 and 9),
    content VARCHAR(1024) NOT NULL
);

grant INSERT, SELECT, UPDATE, DELETE ON users to docker;
grant INSERT, SELECT, UPDATE, DELETE ON messages to docker;