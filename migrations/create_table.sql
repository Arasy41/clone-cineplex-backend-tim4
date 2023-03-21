
    CREATE TABLE author
    (
        id BIGSERIAL PRIMARY KEY,
        full_name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT timezone('utc', now()),
        updated_at TIMESTAMP
    );
    CREATE INDEX ON author (full_name);

    CREATE TABLE book
    (
        id BIGSERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        synopsis VARCHAR(3000),
        cover_url VARCHAR(512),
        content TEXT NOT NULL,
        author_id BIGINT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT timezone('utc', now()),
        updated_at TIMESTAMP
    );
    CREATE INDEX ON book (title);
    ALTER TABLE book ADD FOREIGN KEY (author_id) REFERENCES author (id);









