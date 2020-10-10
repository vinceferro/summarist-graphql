BEGIN;

CREATE TYPE role_t as enum('user', 'admin');

ALTER TABLE account
ADD COLUMN role role_t;

END;