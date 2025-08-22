create table api_keys (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    hashed_key VARCHAR(64) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL
);

CREATE UNIQUE INDEX idx_api_keys_key ON api_keys(hashed_key);