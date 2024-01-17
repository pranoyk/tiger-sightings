CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

ALTER TABLE users RENAME COLUMN name TO username;
ALTER TABLE users ALTER COLUMN user_id SET DEFAULT uuid_generate_v4();
ALTER TABLE tigers ALTER COLUMN tiger_id SET DEFAULT uuid_generate_v4();
ALTER TABLE tiger_sightings ALTER COLUMN tiger_sighting_id SET DEFAULT uuid_generate_v4();
