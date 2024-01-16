CREATE TABLE IF NOT EXISTS tiger_sightings(
    tiger_sighting_id UUID PRIMARY KEY,
    tiger_id UUID NOT NULL REFERENCES tigers(tiger_id),
    created_by UUID NOT NULL REFERENCES users(user_id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_seen TIMESTAMP NOT NULL,
    lat DOUBLE PRECISION NOT NULL,
    lng DOUBLE PRECISION NOT NULL
);
