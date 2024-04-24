CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    service_name TEXT NOT NULL,
    service_user_id TEXT NOT NULL,
    display_name TEXT NOT NULL DEFAULT '',
    created_on TIMESTAMP WITH TIME ZONE DEFAULT (now() AT TIME ZONE('utc'))
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_service_user ON users (service_name, service_user_id);

