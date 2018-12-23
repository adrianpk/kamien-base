--BEGIN;
CREATE TABLE profiles
(
  id UUID PRIMARY KEY,
	name VARCHAR(64) NULL,
	description TEXT NULL,
	email VARCHAR(255),
	location VARCHAR(255) NULL,
	bio VARCHAR(255),
	moto VARCHAR(255),
	website VARCHAR(255),
	aniversary_date TIMESTAMP,
	avatar VARCHAR(255),
	host VARCHAR(255),
	avatar_path VARCHAR(255),
	header_path VARCHAR(255),
  owner_id UUID
);
--
ALTER TABLE profiles
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
-- ALTER TABLE profiles
-- 	ADD CONSTRAINT owner_id_fkey
-- 	FOREIGN KEY (owner_id)
-- 	REFERENCES users
-- --	ON DELETE CASCADE;
--
-- ALTER TABLE profiles
-- 	ADD CONSTRAINT created_by_id_fkey
-- 	FOREIGN KEY (created_by_id)
-- 	REFERENCES users
-- 	ON DELETE CASCADE;
--
-- ALTER TABLE profiles
-- 	ADD CONSTRAINT updated_by_id_fkey
-- 	FOREIGN KEY (updated_by_id)
-- 	REFERENCES profiles
-- 	ON DELETE CASCADE;
--
--COMMIT;


