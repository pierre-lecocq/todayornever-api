CREATE TABLE IF NOT EXISTS task (
    id INTEGER,
    user_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    state TEXT,
    due_at TIMESTAMP DEFAULT (datetime('now', '+1 day', 'start of day')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    position INTEGER,
    PRIMARY KEY(id AUTOINCREMENT)
);
