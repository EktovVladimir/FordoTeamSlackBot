-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.slack_post
    ADD COLUMN meta JSONB NOT NULL DEFAULT '{}'::jsonb;

COMMENT ON COLUMN public.slack_post.meta IS 'Дополнительные мета-данные сообщения';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.slack_post
    DROP COLUMN IF EXISTS meta;
-- +goose StatementEnd