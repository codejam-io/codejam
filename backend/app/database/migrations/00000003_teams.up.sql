CREATE TABLE IF NOT EXISTS teams (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    event_id UUID references events(id),
    owner_user_id UUID references users(id) ON DELETE NO ACTION ON UPDATE NO ACTION,
    name TEXT NOT NULL,
    visibility TEXT,
    timezone TEXT,
    technologies TEXT,
    availability TEXT NOT NULL,
    description TEXT NOT NULL,
    created_on TIMESTAMP WITH TIME ZONE DEFAULT (now() AT TIME ZONE('utc'))
);