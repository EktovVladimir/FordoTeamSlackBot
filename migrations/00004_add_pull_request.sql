-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.pull_requests (
                                      id BIGSERIAL PRIMARY KEY,
                                      owner TEXT NOT NULL,
                                      repo TEXT NOT NULL,
                                      number INTEGER NOT NULL
);

CREATE TABLE public.code_review_pull_requests (
                                                  code_review_id BIGINT NOT NULL,
                                                  pull_request_id BIGINT NOT NULL,

                                                  PRIMARY KEY (code_review_id, pull_request_id),

                                                  CONSTRAINT fk_code_review
                                                      FOREIGN KEY (code_review_id)
                                                          REFERENCES public.code_reviews(id)
                                                          ON DELETE CASCADE,

                                                  CONSTRAINT fk_pull_request
                                                      FOREIGN KEY (pull_request_id)
                                                          REFERENCES public.pull_requests(id)
                                                          ON DELETE CASCADE
);

CREATE INDEX idx_crpr_code_review_id ON public.code_review_pull_requests(code_review_id);
CREATE INDEX idx_crpr_pull_request_id ON public.code_review_pull_requests(pull_request_id);

COMMENT ON COLUMN public.code_reviews.pull_request_number IS 'Устаревшее поле, использовать связь через code_review_pull_requests';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.code_review_pull_requests;
DROP TABLE IF EXISTS public.pull_requests;
COMMENT ON COLUMN public.code_reviews.pull_request_number IS '';
-- +goose StatementEnd