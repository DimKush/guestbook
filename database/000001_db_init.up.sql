CREATE TABLE users
(
	id serial not null unique,
	name varchar(255) not null,
	username varchar(255) not null,
	email varchar(255) not null,
	password_hash varchar(255) not null,
	registration_date timestamptz not null 
);

CREATE TABLE events_lists 
(
	id  serial not null unique,
	title varchar(255) not null,
	description text not null,
	owner_user_id int references users(id) on delete cascade not null
);


CREATE TABLE item_type (
	type_id serial not null unique,
	systemname varchar(50) not null,
	fullname varchar(255) not null 
);

CREATE TABLE items
(
	id serial not null unique,
	list_id int references events_lists(id) on delete cascade not null,
	item_type_id int references item_type(type_id) on delete cascade not null,
	description text not null
);


CREATE TABLE audit_events (
  event_id serial not null unique,
  service_name varchar(255) NULL,
  initiator varchar(255) NOT NULL,
  event_type varchar(255) NOT NULL,
  event_date timestamptz NOT NULL,
  is_panic bool NULL,
  description text NULL
);

CREATE TABLE email_events (
	event_id serial not null unique,
	sender varchar(255) not null,
	receiver varchar(255) not null,
	email_body text not null
);


