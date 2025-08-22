DELETE FROM moods;
DELETE FROM sqlite_sequence WHERE name='moods';

DELETE FROM api_keys;
DELETE FROM sqlite_sequence WHERE name='api_keys';
VACUUM;

