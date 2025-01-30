CREATE TABLE public.history (
    id   BIGSERIAL PRIMARY KEY,
    money VARCHAR(10) NOT NULL,
    source VARCHAR(255) NOT NULL,
    source_web VARCHAR(255) NOT NULL,
    change DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- Tabla de usuarios
CREATE TABLE public.users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Tabla de suscripciones
CREATE TABLE public.subscriptions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES public.users(id),
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    status VARCHAR(50) NOT NULL, -- e.g., 'active', 'inactive', 'cancelled'
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Tabla de pagos
CREATE TABLE public.payments (
    id BIGSERIAL PRIMARY KEY,
    subscription_id BIGINT NOT NULL REFERENCES public.subscriptions(id),
    amount DECIMAL(10, 2) NOT NULL,
    payment_date TIMESTAMP NOT NULL,
    payment_method VARCHAR(50) NOT NULL, -- e.g., 'credit_card', 'paypal'
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);