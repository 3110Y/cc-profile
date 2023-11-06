-- +goose Up
create unique index profile_email_uidx on profile using btree (email);
create unique index profile_phone_uidx on profile using btree (phone);
-- +goose Down
DROP INDEX profile_email_uidx;
DROP INDEX profile_phone_uidx;