BEGIN;

ALTER TABLE book
ADD COLUMN publisher TEXT,
ADD COLUMN cover_url TEXT,
ADD COLUMN overview TEXT,
ADD COLUMN key_insights TEXT[];

END;