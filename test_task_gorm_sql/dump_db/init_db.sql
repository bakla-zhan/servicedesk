CREATE TABLE IF NOT EXISTS requests (
`id` SERIAL,
`head` CHARACTER VARYING(100) NOT NULL,
`body` TEXT,
`email` CHARACTER VARYING(100) NOT NULL);

INSERT INTO requests VALUES (NULL, "request1", "test1 test1", "bob@mail.ru");