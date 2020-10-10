BEGIN;

ALTER TABLE book
DROP COLUMN publisher,
DROP COLUMN cover_url,
DROP COLUMN overview,
DROP COLUMN key_insights;

END;