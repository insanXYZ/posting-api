CREATE TABLE users (
    id varchar(100) not null,
    username varchar(50) not null,
    email varchar(50) unique not null,
    password varchar(100) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    deleted_at timestamp null,
    PRIMARY KEY(id)
);
