CREATE TABLE users (
    id SERIALIZABLE,
    name VARCHAR,
    site VARCHAR,
    password VARCHAR,
    status VARCHAR,
    create_at TIMESTAMP,
    update_at TIMESTAMP
);