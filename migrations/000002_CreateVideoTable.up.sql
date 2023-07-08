CREATE TABLE "videos" (
    id SERIAL PRIMARY KEY,
    title VARCHAR(20) NOT NULL,
    description VARCHAR(100),
    url VARCHAR(255) NOT NULL,
    author_id INT,
    CONSTRAINT fk_authors_video FOREIGN KEY (author_id) REFERENCES authors(id)
);