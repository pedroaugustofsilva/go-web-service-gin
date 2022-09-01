CREATE TABLE albums (
    id serial PRIMARY KEY,
    title VARCHAR (50) NOT NULL,
    artist VARCHAR (50) NOT NULL,
    price FLOAT NOT NULL
);

INSERT INTO albums (title, artist, price) VALUES ('Blue Train', 'John Coltrane', 56.99);
INSERT INTO albums (title, artist, price) VALUES ('Jeru', 'Gerry Mulligan', 17.99);
INSERT INTO albums (title, artist, price) VALUES ('Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 39.99);
