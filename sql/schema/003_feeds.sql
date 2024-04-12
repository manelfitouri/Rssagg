-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
    --   I'd recommend making the `url` field unique so that in the future we aren't downloading duplicate posts. 
    -- I'd also recommend using `ON DELETE CASCADE` on 
    -- the `user_id` foreign key so that if a user is deleted, all of their feeds are automatically deleted as well.
);

-- +goose Down
DROP TABLE feeds;