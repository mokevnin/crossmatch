-- name: CreateChannel :one
INSERT INTO channels (id, owner_id, username, topic, language, subscribers, avg_reach)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetChannel :one
SELECT * FROM channels WHERE id = $1;

-- name: ListCandidates :many
-- Channels of the same topic and language whose reach is within the given ratio of the
-- channel's own, excluding itself and every channel it already decided on.
SELECT c.*
FROM channels AS c
WHERE c.id <> sqlc.arg(channel_id)::bigint
  AND c.topic = sqlc.arg(topic)::text
  AND c.language = sqlc.arg(language)::text
  AND c.avg_reach::float8 BETWEEN sqlc.arg(own_reach)::float8 / sqlc.arg(max_ratio)::float8
                              AND sqlc.arg(own_reach)::float8 * sqlc.arg(max_ratio)::float8
  AND NOT EXISTS (
      SELECT 1 FROM decisions AS d
      WHERE d.from_channel_id = sqlc.arg(channel_id)::bigint AND d.to_channel_id = c.id
  )
ORDER BY abs(c.avg_reach - sqlc.arg(own_reach)::int)
LIMIT sqlc.arg(max_rows)::int;

-- name: RecordDecision :exec
INSERT INTO decisions (from_channel_id, to_channel_id, interested)
VALUES ($1, $2, $3)
ON CONFLICT (from_channel_id, to_channel_id)
DO UPDATE SET interested = EXCLUDED.interested, decided_at = now();

-- name: IsMutual :one
SELECT EXISTS (
    SELECT 1
    FROM decisions AS a
    JOIN decisions AS b
      ON b.from_channel_id = a.to_channel_id AND b.to_channel_id = a.from_channel_id
    WHERE a.from_channel_id = $1 AND a.to_channel_id = $2
      AND a.interested AND b.interested
);
