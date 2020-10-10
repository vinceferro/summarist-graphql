BEGIN;

CREATE TABLE book_star
(
    book_id    UUID REFERENCES book (id),
    account_id UUID REFERENCES account (id),
    PRIMARY KEY(book_id, account_id)
);

CREATE INDEX ON book_star(account_id);

END;
