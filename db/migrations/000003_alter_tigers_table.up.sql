ALTER TABLE TIGERS ADD created_by UUID NOT NULL REFERENCES users(user_id);
ALTER TABLE TIGERS DROP COLUMN last_seen, DROP COLUMN lat, DROP COLUMN lng;