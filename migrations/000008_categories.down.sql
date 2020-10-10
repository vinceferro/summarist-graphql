BEGIN;

ALTER TABLE book
DROP COLUMN category_id;

DROP TABLE category;

END;