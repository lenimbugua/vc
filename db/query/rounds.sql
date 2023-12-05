-- name: CreatePeriod :one
INSERT INTO periods (
  competition_id,
  end_time,
  round_id,
  start_time
) VALUES(
    $1, $2, $3, $4
) RETURNING *;
