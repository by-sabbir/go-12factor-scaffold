CREATE TABLE blog (
	id uuid NOT NULL,
	title varchar(200) NOT NULL UNIQUE,
	slug varchar(250) NOT NULL UNIQUE,
	author varchar(80) NOT NULL,
	body varchar(3000) NOT NULL
);
CREATE INDEX blog_id_idx ON blog USING HASH (id);
CREATE INDEX blog_title_idx ON blog USING HASH (title);