CREATE TABLE IF NOT EXISTS "sessions" (
    "user_uid" UUID,
    "token" VARCHAR(1000),
    "expires_at" BIGINT NOT NULL,
    "issued_at" BIGINT NOT NULL,
    PRIMARY KEY ("user_uid", "token")
);