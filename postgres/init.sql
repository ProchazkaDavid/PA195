CREATE TABLE messages (
  id SERIAL PRIMARY KEY,
  sender VARCHAR(256) NOT NULL,
  date VARCHAR(256),
  room VARCHAR(256) NOT NULL,
  text TEXT
);
grant
INSERT,
SELECT,
UPDATE,
  DELETE ON messages TO docker;
grant
INSERT,
SELECT,
UPDATE,
  DELETE ON messages_id_seq TO docker;