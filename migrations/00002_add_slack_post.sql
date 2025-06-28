-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.slack_post (
                                   id BIGSERIAL PRIMARY KEY,
                                   created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                   deleted_at TIMESTAMP WITH TIME ZONE,

                                   channel_id TEXT NOT NULL,
                                   thread_ts TEXT NOT NULL,
                                   type TEXT NOT NULL CHECK (type IN ('thread', 'reply'))
);

ALTER TABLE public.code_reviews
    ADD COLUMN slack_post_id BIGINT;

INSERT INTO public.slack_post (channel_id, thread_ts, type)
SELECT
    'legacy_channel' AS channel_id,
    cr.thread_ts AS thread_ts,
    'thread' AS type
FROM public.code_reviews cr
WHERE cr.thread_ts IS NOT NULL
GROUP BY cr.thread_ts;

UPDATE public.code_reviews cr
SET slack_post_id = sp.id
FROM public.slack_post sp
WHERE cr.thread_ts = sp.thread_ts;

ALTER TABLE public.code_reviews
    ALTER COLUMN slack_post_id SET NOT NULL,
    ADD CONSTRAINT fk_code_reviews_slack_post FOREIGN KEY (slack_post_id) REFERENCES public.slack_post(id);

COMMENT ON COLUMN public.code_reviews.thread_ts IS 'Устаревшее поле, использовать slack_post_id';
ALTER TABLE public.code_reviews
    ALTER COLUMN thread_ts DROP NOT NULL;

CREATE INDEX idx_code_reviews_slack_post_id ON public.code_reviews(slack_post_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Удаляем связь и колонку
ALTER TABLE public.code_reviews
    DROP CONSTRAINT IF EXISTS fk_code_reviews_slack_post,
    DROP COLUMN IF EXISTS slack_post_id;

ALTER TABLE public.code_reviews
    ALTER COLUMN thread_ts SET NOT NULL;
COMMENT ON COLUMN public.code_reviews.thread_ts IS '';

DROP TABLE IF EXISTS public.slack_post;
-- +goose StatementEnd