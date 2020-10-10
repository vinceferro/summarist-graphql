BEGIN;

ALTER TABLE account
DROP COLUMN role role_t;

DROP TYPE role_t;

END;