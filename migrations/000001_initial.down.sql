BEGIN;

DROP TRIGGER update_usr_modtime ON public.usr;
DROP FUNCTION update_modified_column();
DROP TABLE public.usr;

END;