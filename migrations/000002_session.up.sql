BEGIN;

CREATE TABLE public.user_session
(
    id         serial PRIMARY KEY,
    token      uuid        NOT NULL,
    user_id    int         NOT NULL,
    expires_in bigint      NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    FOREIGN KEY (user_id) REFERENCES public.usr (id) ON DELETE CASCADE
);

END;