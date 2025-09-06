CREATE VIEW vw_reports AS
SELECT 
    r.id,
    r.user_id,
    r.report,
    r.mood_data,
    r.start_date,
    r.end_date,
    r.created_at,
    r.updated_at,
    r.deleted_at
FROM reports r
WHERE r.deleted_at IS NULL;