-- table history
-- name: GetLastHistory :one
SELECT id, money, source, source_web, change, created_at
FROM public.history
WHERE money = $1 AND source = $2 AND source_web = $3
ORDER BY id DESC
LIMIT 1;

-- name: FindHistory :one
SELECT id, money, source, source_web, change, created_at
FROM public.history
WHERE id = $1;

-- name: ListHistory :many
SELECT id, money, source, source_web, change, created_at
FROM public.history
WHERE source_web = $1
  AND created_at >= $2::timestamp - INTERVAL '8 hours'
ORDER BY id DESC
LIMIT $3;


-- name: CreateHistory :one
INSERT INTO public.history (money, source, source_web, change, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, money, source, source_web, change, created_at;


-- table users
-- name: CreateUser :one
INSERT INTO public.users (email, password, created_at)
VALUES ($1, $2, $3)
RETURNING id, email, password, created_at;

-- name: GetUser :one
SELECT * FROM public.users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE public.users
SET email = $2, password = $3
WHERE id = $1
RETURNING id, email, password, created_at;

-- name: DeleteUser :exec
DELETE FROM public.users
WHERE email = $1;

-- table subscriptions
-- name: CreateSubscription :one
INSERT INTO public.subscriptions (user_id, start_date, end_date, status, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, start_date, end_date, status, created_at;

-- name: GetSubscription :one
SELECT * FROM public.subscriptions
WHERE user_id = $1;

-- name: UpdateSubscription :one
UPDATE public.subscriptions
SET end_date = $2, status = $3
WHERE user_id = $1
RETURNING id, user_id, start_date, end_date, status, created_at;

-- name: DeleteSubscription :exec
DELETE FROM public.subscriptions
WHERE id = $1;

-- table payments
-- name: CreatePayment :one
INSERT INTO public.payments (subscription_id, amount, payment_date, payment_method, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, subscription_id, amount, payment_date, payment_method, created_at;

-- name: GetPayment :many
SELECT * FROM public.payments
WHERE subscription_id = $1;

-- name: UpdatePayment :one
UPDATE public.payments
SET amount = $2, payment_date = $3, payment_method = $4
WHERE id = $1
RETURNING id, subscription_id, amount, payment_date, payment_method, created_at;

-- name: DeletePayment :exec
DELETE FROM public.payments
WHERE id = $1;