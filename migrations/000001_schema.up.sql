-- importing uuid-ossp extension to be able to use uuid_generate_v4 function
CREATE EXTENSION "uuid-ossp";

CREATE TABLE book
(
    id         UUID PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    title      TEXT NOT NULL,
    author     TEXT,
    audio_path TEXT,
    isbn       BIGINT
);
