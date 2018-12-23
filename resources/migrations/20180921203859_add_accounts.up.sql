--BEGIN;
CREATE TABLE accounts
(
	id UUID PRIMARY KEY,
  name VARCHAR(64),
	account_type VARCHAR(16),
	email VARCHAR(64),
	owner_id UUID,
	parent_id UUID,
	description TEXT NULL
);
--
ALTER TABLE accounts
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
-- ALTER TABLE accounts
-- 	ADD CONSTRAINT owner_id_fkey
-- 	FOREIGN KEY (owner_id)
-- 	REFERENCES users
-- --	ON DELETE CASCADE;
--
-- ALTER TABLE accounts
-- 	ADD CONSTRAINT parent_id_fkey
-- 	FOREIGN KEY (parent_id)
-- 	REFERENCES accounts
-- 	ON DELETE CASCADE;
--
-- ALTER TABLE accounts
-- 	ADD CONSTRAINT created_by_id_fkey
-- 	FOREIGN KEY (created_by_id)
-- 	REFERENCES users
-- 	ON DELETE CASCADE;
--
-- ALTER TABLE accounts
-- 	ADD CONSTRAINT updated_by_id_fkey
-- 	FOREIGN KEY (updated_by_id)
-- 	REFERENCES accounts
-- 	ON DELETE CASCADE;
--
--COMMIT;


