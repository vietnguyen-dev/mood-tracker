CREATE VIEW vw_moods AS
SELECT 
  * 
FROM moods 
WHERE deleted_at IS NULL;