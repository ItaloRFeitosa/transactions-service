ALTER TABLE public.accounts ADD COLUMN IF NOT EXISTS available_credit_limit BIGINT;

ALTER TABLE public.accounts ADD COLUMN IF NOT EXISTS version BIGINT DEFAULT 0;
