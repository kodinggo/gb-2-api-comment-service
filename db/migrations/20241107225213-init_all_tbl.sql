
-- +migrate Up
CREATE TABLE comments (
    id      INT PRIMARY KEY AUTO_INCREMENT ,
    comment TEXT,
    story_id INT,
    user_id INT,
    created_at timestamp,
    updated_at timestamp
);
-- +migrate Down

