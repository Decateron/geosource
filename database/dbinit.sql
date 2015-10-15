CREATE DOMAIN username AS VARCHAR(20);
CREATE DOMAIN email AS VARCHAR(254);
CREATE DOMAIN channelname AS VARCHAR(20);
CREATE TYPE visibility AS ENUM ('public', 'private');
CREATE DOMAIN title AS VARCHAR(140);
CREATE DOMAIN comment AS VARCHAR(500);

CREATE TABLE users (
	u_username username PRIMARY KEY,
	u_email email NOT NULL UNIQUE
);
CREATE UNIQUE INDEX lower_username_idx ON users (lower(u_username));

CREATE TABLE admins (
	a_username username PRIMARY KEY,
	FOREIGN KEY (a_username) REFERENCES users (u_username)
);

CREATE TABLE channels (
	ch_channelname channelname PRIMARY KEY,
	ch_username_creator username NOT NULL,
	ch_visibility visibility NOT NULL,
	ch_form jsonb NOT NULL,
	FOREIGN KEY (ch_username_creator) REFERENCES channels (ch_channelname)
);

CREATE TABLE posts (
	p_pid SERIAL PRIMARY KEY,
	p_username_creator username NOT NULL,
	p_channelname channelname NOT NULL,
	p_title title NOT NULL,
	p_time TIMESTAMP NOT NULL,
	p_location POINT NOT NULL,
	p_fields jsonb NOT NULL,
	FOREIGN KEY (p_username_creator) REFERENCES users (u_username),
	FOREIGN KEY (p_channelname) REFERENCES channels (ch_channelname)
);
CREATE INDEX post_gix ON posts USING GIST (p_location); 

CREATE TABLE comments (
	cmt_cid SERIAL PRIMARY KEY,
	cmt_pid INTEGER NOT NULL,
	cmt_cid_parent INTEGER,
	cmt_comment comment NOT NULL,
	FOREIGN KEY (cmt_pid) REFERENCES posts (p_pid),
	FOREIGN KEY (cmt_cid_parent) REFERENCES comments (cmt_cid)
);

CREATE TABLE user_favorites (
	uf_username username NOT NULL,
	uf_pid INTEGER NOT NULL,
	FOREIGN KEY (uf_username) REFERENCES users (u_username),
	FOREIGN KEY (uf_pid) REFERENCES posts (p_pid),
	PRIMARY KEY (uf_username, uf_pid)
);

CREATE TABLE user_subscriptions (
	us_username username NOT NULL,
	us_channelname channelname NOT NULL,
	FOREIGN KEY (us_username) REFERENCES users (u_username),
	FOREIGN KEY (us_channelname) REFERENCES channels (ch_channelname),
	PRIMARY KEY (us_username, us_channelname)
);

CREATE TABLE channel_moderators (
	chm_username username NOT NULL,
	chm_channelname channelname NOT NULL,
	FOREIGN KEY (chm_username) REFERENCES users (u_username),
	FOREIGN KEY (chm_channelname) REFERENCES channels (ch_channelname),
	PRIMARY KEY (chm_username, chm_channelname)
);

CREATE TABLE channel_viewers (
	chv_username username NOT NULL,
	chv_channelname channelname NOT NULL,
	FOREIGN KEY (chv_username) REFERENCES users (u_username),
	FOREIGN KEY (chv_channelname) REFERENCES channels (ch_channelname),
	PRIMARY KEY (chv_username, chv_channelname)
);

CREATE TABLE channel_bans (
	chb_username username NOT NULL,
	chb_channelname channelname NOT NULL,
	FOREIGN KEY (chb_username) REFERENCES users (u_username),
	FOREIGN KEY (chb_channelname) REFERENCES channels (ch_channelname),
	PRIMARY KEY (chb_username, chb_channelname)
);
