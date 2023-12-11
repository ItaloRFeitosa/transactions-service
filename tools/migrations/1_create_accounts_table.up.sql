CREATE TABLE IF NOT EXISTS public.accounts (
    account_id BIGSERIAL PRIMARY KEY,
    document_type character varying COLLATE pg_catalog."default",
    document_number character varying COLLATE pg_catalog."default",
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ DEFAULT NULL
);