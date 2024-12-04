BEGIN;
CREATE TABLE tree_folders (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    is_folder boolean default true,
    parent_id INTEGER REFERENCES tree_folders(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_parent_id ON tree_folders(parent_id);

COMMIT;