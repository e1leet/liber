BEGIN;

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = ON;
SET check_function_bodies = FALSE;
SET client_min_messages = WARNING;
SET search_path = public, extensions;
SET default_tablespace = '';
SET default_with_oids = FALSE;

CREATE OR REPLACE FUNCTION update_modified_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE public.usr
(
    id         serial PRIMARY KEY,
    email      varchar(256) UNIQUE       NOT NULL,
    username   varchar(32) UNIQUE        NOT NULL,
    password   varchar(64)               NOT NULL,
    created_at timestamptz DEFAULT now() NOT NULL,
    updated_at timestamptz DEFAULT now() NOT NULL
);

CREATE TRIGGER update_usr_modtime
    BEFORE UPDATE
    ON public.usr
    FOR EACH ROW
EXECUTE PROCEDURE update_modified_column();

COMMIT;