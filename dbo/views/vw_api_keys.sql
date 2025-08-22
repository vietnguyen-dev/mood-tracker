CREATE VIEW vw_api_keys AS 
SELECT 
    id,
    hashed_key
FROM api_keys 
WHERE deleted_at IS NULL;
