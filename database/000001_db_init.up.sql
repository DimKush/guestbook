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
	description varchar(255) not null
);

CREATE TABLE users_lists
(
	id serial not null unique,
	user_id int references users(id) on delete cascade  not null,
	event_list_id int references events_lists(id) on delete cascade not null
);

CREATE TABLE event_item
(
	id serial not null unique,
	list_id int references events_lists(id) on delete cascade not null,
	title varchar(255) not null,
	description varchar(255),
	img_path varchar(255) not null
);


CREATE TABLE audit_events (
  eventid serial not null unique,
  eventtype varchar(255) NOT NULL,
  eventdate timestamptz NOT NULL,
  servicename varchar(255) NULL,
  initiator varchar(255) NOT NULL,
  is_panic bool NULL,
  description text NULL
);