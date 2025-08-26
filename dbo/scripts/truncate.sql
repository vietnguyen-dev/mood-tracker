DELETE FROM moods;
DELETE FROM sqlite_sequence WHERE name='moods';
VACUUM

DELETE FROM api_keys;
DELETE FROM sqlite_sequence WHERE name='api_keys';
VACUUM;

