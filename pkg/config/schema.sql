DROP TABLE IF EXISTS bookdata;

CREATE TABLE bookdata (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    Name TEXT NOT NULL,
    Author TEXT NOT NULL,
    Publication TEXT NOT NULL
);