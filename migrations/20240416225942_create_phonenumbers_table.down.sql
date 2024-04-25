DROP TABLE IF EXISTS "phonenumbers";

DROP TRIGGER IF EXISTS phonenumbers_set_updated_at ON "phonenumbers";
DROP FUNCTION IF EXISTS phonenumbers_set_updated_at;
