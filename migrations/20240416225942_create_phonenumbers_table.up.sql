
SET TIME ZONE 'Europe/Moscow';

CREATE TABLE IF NOT EXISTS "phonenumbers" (
    "uid"           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "phonenumber"   CHAR(10) NOT NULL UNIQUE,

    "is_registered" BOOLEAN  NOT NULL DEFAULT false,

    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION phonenumbers_set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW."updated_at" = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER phonenumbers_set_updated_at
BEFORE UPDATE ON "phonenumbers"
FOR EACH ROW
EXECUTE PROCEDURE phonenumbers_set_updated_at();