--BEGIN;
CREATE TABLE users
(
  id UUID PRIMARY KEY,
  username VARCHAR(32) UNIQUE,
  password_hash CHAR(128),
  email VARCHAR(255) UNIQUE,
  given_name VARCHAR(64),
  middle_names VARCHAR(128) NULL,
  family_name VARCHAR(128),
  context_id UUID
);
--
ALTER TABLE users
ADD COLUMN geolocation geography
(Point,4326),
ADD COLUMN starts_at TIMESTAMP
WITH TIME ZONE,
ADD COLUMN ends_at TIMESTAMP
WITH TIME ZONE,
ADD COLUMN is_active BOOLEAN,
ADD COLUMN is_logical_deleted BOOLEAN,
ADD COLUMN created_by_id UUID,
ADD COLUMN updated_by_id UUID,
ADD COLUMN created_at TIMESTAMP
WITH TIME ZONE,
ADD COLUMN updated_at TIMESTAMP
WITH TIME ZONE;
--COMMIT;
