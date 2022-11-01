-- This script was generated by the Schema Diff utility in pgAdmin 4
-- For the circular dependencies, the order in which Schema Diff writes the objects is not very sophisticated
-- and may require manual changes to the script to ensure changes are applied in the correct order.
-- Please report an issue for any failure with the reproduction steps.

CREATE OR REPLACE FUNCTION public.setteamcode()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    VOLATILE
    COST 100
AS $BODY$
BEGIN
  NEW.code = (select CONCAT((select value as adjective from adjectives order by random() limit 1), '-', (select value as colour from colours order by random() limit 1), '-', (select value as animal from animals order by random() limit 1)) as code);
  RETURN NEW;
END;
$BODY$;

CREATE TABLE IF NOT EXISTS public.adjectives
(
    value text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT adjectives_pkey PRIMARY KEY (value),
    CONSTRAINT adjectives_value_key UNIQUE (value)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.adjectives
    OWNER to supabase_admin;

ALTER TABLE IF EXISTS public.adjectives
    ENABLE ROW LEVEL SECURITY;

GRANT ALL ON TABLE public.adjectives TO anon;

GRANT ALL ON TABLE public.adjectives TO postgres;

GRANT ALL ON TABLE public.adjectives TO supabase_admin;

GRANT ALL ON TABLE public.adjectives TO authenticated;

GRANT ALL ON TABLE public.adjectives TO service_role;

CREATE TABLE IF NOT EXISTS public.animals
(
    value text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT animals_pkey PRIMARY KEY (value)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.animals
    OWNER to supabase_admin;

ALTER TABLE IF EXISTS public.animals
    ENABLE ROW LEVEL SECURITY;

GRANT ALL ON TABLE public.animals TO anon;

GRANT ALL ON TABLE public.animals TO postgres;

GRANT ALL ON TABLE public.animals TO supabase_admin;

GRANT ALL ON TABLE public.animals TO authenticated;

GRANT ALL ON TABLE public.animals TO service_role;

CREATE TABLE IF NOT EXISTS public.colours
(
    value text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT colours_pkey PRIMARY KEY (value)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.colours
    OWNER to supabase_admin;

ALTER TABLE IF EXISTS public.colours
    ENABLE ROW LEVEL SECURITY;

GRANT ALL ON TABLE public.colours TO anon;

GRANT ALL ON TABLE public.colours TO postgres;

GRANT ALL ON TABLE public.colours TO supabase_admin;

GRANT ALL ON TABLE public.colours TO authenticated;

GRANT ALL ON TABLE public.colours TO service_role;

REVOKE ALL ON TABLE public.guesses FROM anon;
REVOKE ALL ON TABLE public.guesses FROM authenticated;
REVOKE ALL ON TABLE public.guesses FROM postgres;
REVOKE ALL ON TABLE public.guesses FROM service_role;
REVOKE ALL ON TABLE public.guesses FROM supabase_admin;
GRANT ALL ON TABLE public.guesses TO anon;

GRANT ALL ON TABLE public.guesses TO postgres;

GRANT ALL ON TABLE public.guesses TO supabase_admin;

GRANT ALL ON TABLE public.guesses TO authenticated;

GRANT ALL ON TABLE public.guesses TO service_role;

REVOKE ALL ON TABLE public.adventures FROM anon;
REVOKE ALL ON TABLE public.adventures FROM authenticated;
REVOKE ALL ON TABLE public.adventures FROM postgres;
REVOKE ALL ON TABLE public.adventures FROM service_role;
REVOKE ALL ON TABLE public.adventures FROM supabase_admin;
GRANT ALL ON TABLE public.adventures TO anon;

GRANT ALL ON TABLE public.adventures TO postgres;

GRANT ALL ON TABLE public.adventures TO supabase_admin;

GRANT ALL ON TABLE public.adventures TO authenticated;

GRANT ALL ON TABLE public.adventures TO service_role;

REVOKE ALL ON TABLE public.games FROM anon;
REVOKE ALL ON TABLE public.games FROM authenticated;
REVOKE ALL ON TABLE public.games FROM postgres;
REVOKE ALL ON TABLE public.games FROM service_role;
REVOKE ALL ON TABLE public.games FROM supabase_admin;
GRANT ALL ON TABLE public.games TO anon;

GRANT ALL ON TABLE public.games TO postgres;

GRANT ALL ON TABLE public.games TO supabase_admin;

GRANT ALL ON TABLE public.games TO authenticated;

GRANT ALL ON TABLE public.games TO service_role;
CREATE OR REPLACE TRIGGER "SetTeamCode"
    BEFORE INSERT
    ON public.games
    FOR EACH ROW
    EXECUTE FUNCTION public.setteamcode();

REVOKE ALL ON TABLE public.admins FROM anon;
REVOKE ALL ON TABLE public.admins FROM authenticated;
REVOKE ALL ON TABLE public.admins FROM postgres;
REVOKE ALL ON TABLE public.admins FROM service_role;
REVOKE ALL ON TABLE public.admins FROM supabase_admin;
GRANT ALL ON TABLE public.admins TO anon;

GRANT ALL ON TABLE public.admins TO postgres;

GRANT ALL ON TABLE public.admins TO supabase_admin;

GRANT ALL ON TABLE public.admins TO authenticated;

GRANT ALL ON TABLE public.admins TO service_role;

REVOKE ALL ON TABLE public.puzzles FROM anon;
REVOKE ALL ON TABLE public.puzzles FROM authenticated;
REVOKE ALL ON TABLE public.puzzles FROM postgres;
REVOKE ALL ON TABLE public.puzzles FROM service_role;
REVOKE ALL ON TABLE public.puzzles FROM supabase_admin;
GRANT ALL ON TABLE public.puzzles TO anon;

GRANT ALL ON TABLE public.puzzles TO postgres;

GRANT ALL ON TABLE public.puzzles TO supabase_admin;

GRANT ALL ON TABLE public.puzzles TO authenticated;

GRANT ALL ON TABLE public.puzzles TO service_role;
