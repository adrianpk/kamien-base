CREATE TABLE permissions
(
	id UUID PRIMARY KEY,
	name VARCHAR(64) NULL,
	description TEXT NULL,
	organization_id UUID,
	is_active BOOLEAN NULL,
	is_logical_deleted BOOLEAN NULL,
	created_by_id UUID,
	updated_by_id UUID,
	created_at TIMESTAMP WITH TIME ZONE,
	updated_at TIMESTAMP WITH TIME ZONE
);
