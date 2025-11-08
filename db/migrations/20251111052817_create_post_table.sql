-- migrate:up
CREATE TABLE posts (
	post_id text primary key,
	creator_id text not null
);

-- migrate:down
DROP TABLE posts;

