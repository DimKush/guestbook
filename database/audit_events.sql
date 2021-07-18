-- public.audit_events2 definition

-- Drop table

-- DROP TABLE public.audit_events2;

CREATE TABLE public.audit_events (
	eventid numeric(19) NOT NULL DEFAULT nextval('event_id_seq'::regclass),
	eventtype varchar(255) NOT NULL,
	eventdate timestamptz NOT NULL,
	servicename varchar(255) NULL,
	is_panic bool NULL,
	description text NULL
);