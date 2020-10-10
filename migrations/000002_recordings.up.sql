BEGIN;

CREATE TABLE recording (
    id UUID PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
    book_id UUID REFERENCES book(id),
    path TEXT NOT NULL,
    format TEXT NOT NULL,
    version TEXT
);

ALTER TABLE book
DROP COLUMN audio_path;

END;
