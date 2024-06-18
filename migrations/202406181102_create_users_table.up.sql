CREATE TABLE IF NOT EXISTS public.user_data
(
    id SERIAL PRIMARY KEY,
    username VARCHAR(20),
    description VARCHAR(30),
    email VARCHAR(256) NOT NULL UNIQUE,
    phone VARCHAR(15),
    avatar TEXT,
    status VARCHAR(30),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    role VARCHAR(20),
    password_hash TEXT NOT NULL,
    date_of_birth VARCHAR(10),
    privacy_settings JSONB,
    is_active BOOLEAN DEFAULT true,
    last_login TIMESTAMP,
    confirmation_token TEXT,
    social_profiles JSONB
);

ALTER TABLE IF EXISTS public.user_data
    OWNER TO postgres;