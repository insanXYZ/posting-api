CREATE TABLE user_like_post (
    user_id VARCHAR(100) NOT NULL,
    post_id VARCHAR(100) NOT NULL,
    CONSTRAINT post_user_pkey
        PRIMARY KEY(user_id , post_id),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_post
        FOREIGN KEY(post_id) REFERENCES posts(id)
)