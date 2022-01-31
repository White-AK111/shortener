CREATE DATABASE shortener WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';

CREATE TABLE IF NOT EXISTS links
(
    id serial PRIMARY KEY,
    short_url VARCHAR (1024) UNIQUE NOT NULL,
    long_url VARCHAR (1024),
    counter INT NOT NULL
    );

CREATE ROLE shortenerapi WITH LOGIN PASSWORD 'Q1w2e3r4t5';
GRANT ALL PRIVILEGES ON DATABASE shortener TO shortenerapi;
GRANT ALL ON links to shortenerapi;
GRANT ALL ON links_id_seq to shortenerapi;
