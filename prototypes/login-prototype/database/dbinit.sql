CREATE DOMAIN username AS VARCHAR(20);
CREATE DOMAIN email AS VARCHAR(254);

CREATE TABLE users (
	u_email email NOT NULL UNIQUE,
	u_username username PRIMARY KEY
);
CREATE UNIQUE INDEX lower_username_idx ON users ((lower(u_username)));
