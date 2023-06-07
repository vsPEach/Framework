CREATE TABLE IF NOT EXISTS users(
    id uuid primary key,
    username varchar(255),
    email varchar(255),
    password varchar(255),
    role integer,
    is_confirmed boolean,
    created_at TIMESTAMP,
    user_id uuid
);

CREATE TABLE IF NOT EXISTS articles (
    id uuid primary key,
    author_id uuid,
    title varchar(255),
    text text,
    created_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS comments (
    id uuid primary key,
    author_id uuid,
    article_id uuid,
    text text,
    created_at TIMESTAMP
);