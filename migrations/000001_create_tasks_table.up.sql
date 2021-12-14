CREATE TABLE IF NOT EXISTS tasks
(
    id SERIAL Primary Key,
    assignee VARCHAR(50),
    title VARCHAR(50),
    summary VARCHAR(500),
    deadline timestamp with time zone NOT NULL,
    status VARCHAR(50)
);