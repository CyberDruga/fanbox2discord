-- name: GetSavedPosts :many
select * from posts where creator_id = ?;  

-- name: GetPost :one
select * from posts where post_id = ?;  

-- name: SaveNewPost :one
insert into posts (post_id, creator_id)
values (?, ?) returning *;
