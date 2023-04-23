CREATE TABLE company (
	id uuid NOT NULL,
	name varchar(80) NOT NULL UNIQUE,
	description varchar(3000)
);
CREATE INDEX company_id_idx ON company USING HASH (id);