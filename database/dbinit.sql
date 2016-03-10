CREATE TYPE visibility AS ENUM ('public', 'private');
CREATE DOMAIN username AS TEXT;
CREATE DOMAIN avatar AS TEXT;
CREATE DOMAIN email AS VARCHAR(254);
CREATE DOMAIN channelname AS VARCHAR(20);
CREATE DOMAIN title AS VARCHAR(140);
CREATE DOMAIN comment AS VARCHAR(500);
CREATE DOMAIN base64uuid AS VARCHAR(22);
CREATE DOMAIN userid AS base64uuid;
CREATE DOMAIN postid AS base64uuid;
CREATE DOMAIN commentid AS base64uuid;
CREATE DOMAIN thumbnail AS VARCHAR(60);

CREATE TABLE users (
	u_userid userid PRIMARY KEY,
	u_email email NOT NULL,
	u_username username NOT NULL,
	u_avatar avatar
);
CREATE UNIQUE INDEX lower_email_idx ON users (lower(u_email));

CREATE TABLE admins (
	a_userid userid PRIMARY KEY,
	FOREIGN KEY (a_userid) REFERENCES users (u_userid)
);

CREATE TABLE channels (
	ch_channelname channelname PRIMARY KEY,
	ch_userid_creator userid NOT NULL,
	ch_visibility visibility NOT NULL,
	ch_fields jsonb NOT NULL,
	FOREIGN KEY (ch_userid_creator) REFERENCES users (u_userid)
);

CREATE TABLE posts (
	p_postid postid PRIMARY KEY,
	p_userid_creator userid NOT NULL,
	p_channelname channelname NOT NULL,
	p_title title NOT NULL,
	p_thumbnail thumbnail NOT NULL,
	p_time TIMESTAMP NOT NULL,
	p_location POINT NOT NULL,
	p_fields jsonb NOT NULL,
	FOREIGN KEY (p_userid_creator) REFERENCES users (u_userid),
	FOREIGN KEY (p_channelname) REFERENCES channels (ch_channelname)
);
CREATE INDEX post_gix ON posts USING GIST (p_location); 

CREATE TABLE comments (
	cmt_commentid commentid PRIMARY KEY,
	cmt_postid postid NOT NULL,
	cmt_userid_creator userid NOT NULL,
	cmt_commentid_parent commentid,
	cmt_comment comment NOT NULL,
	cmt_time TIMESTAMP NOT NULL,
	FOREIGN KEY (cmt_userid_creator) REFERENCES users (u_userid), 
	FOREIGN KEY (cmt_postid) REFERENCES posts (p_postid),
	FOREIGN KEY (cmt_commentid_parent) REFERENCES comments (cmt_commentid)
);

CREATE TABLE user_favorites (
	uf_userid userid NOT NULL,
	uf_postid postid NOT NULL,
	FOREIGN KEY (uf_userid) REFERENCES users (u_userid),
	FOREIGN KEY (uf_postid) REFERENCES posts (p_postid),
	PRIMARY KEY (uf_userid, uf_postid)
);

CREATE TABLE user_subscriptions (
	us_userid userid NOT NULL,
	us_channelname channelname NOT NULL,
	FOREIGN KEY (us_userid) REFERENCES users (u_userid),
	FOREIGN KEY (us_channelname) REFERENCES channels (ch_channelname),
	PRIMARY KEY (us_userid, us_channelname)
);

CREATE TABLE channel_moderators (
	chm_userid userid NOT NULL,
	chm_channelname channelname NOT NULL,
	FOREIGN KEY (chm_userid) REFERENCES users (u_userid),
	FOREIGN KEY (chm_channelname) REFERENCES channels (ch_channelname),
	PRIMARY KEY (chm_userid, chm_channelname)
);

CREATE TABLE channel_viewers (
	chv_userid userid NOT NULL,
	chv_channelname channelname NOT NULL,
	FOREIGN KEY (chv_userid) REFERENCES users (u_userid),
	FOREIGN KEY (chv_channelname) REFERENCES channels (ch_channelname),
	PRIMARY KEY (chv_userid, chv_channelname)
);

CREATE TABLE channel_bans (
	chb_userid userid NOT NULL,
	chb_channelname channelname NOT NULL,
	FOREIGN KEY (chb_userid) REFERENCES users (u_userid),
	FOREIGN KEY (chb_channelname) REFERENCES channels (ch_channelname),
	PRIMARY KEY (chb_userid, chb_channelname)
);
