CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'inactive',  
    creator_id INT NOT NULL,
    assignee_id INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

ALTER TABLE tasks
    ADD CONSTRAINT fk_creator
    FOREIGN KEY (creator_id)
    REFERENCES users(id)
    ON DELETE CASCADE;

ALTER TABLE tasks
    ADD CONSTRAINT fk_assignee
    FOREIGN KEY (assignee_id)
    REFERENCES users(id)
    ON DELETE CASCADE;

CREATE INDEX IF NOT EXISTS idx_tasks_creator_id ON tasks(creator_id);
CREATE INDEX IF NOT EXISTS idx_tasks_assignee_id ON tasks(assignee_id);
CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks(status);
