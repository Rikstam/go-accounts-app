-- name: CreateEntry :one
INSERT INTO entries (
    account_id,
    amount

) VALUES (
             $1, $2
         ) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY id
LIMIT $2
    OFFSET $3;

-- name: GetEntryWithAccount :one
SELECT
    e.*,
    a.owner,
    a.balance
FROM
    entries as e
        JOIN accounts as a
             ON e.account_id = a.id
WHERE e.id = $1
LIMIT 1;