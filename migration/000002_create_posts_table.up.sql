CREATE TABLE posts (
    id VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_by VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP null,
    PRIMARY KEY (id),
    CONSTRAINT fk_user
        FOREIGN KEY(created_by)
            REFERENCES users(id)
);