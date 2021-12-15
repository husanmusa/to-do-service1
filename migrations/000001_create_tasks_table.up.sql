CREATE TABLE IF NOT EXISTS tasks
(
    id         SERIAL Primary Key,
    assignee   VARCHAR(50),
    title      VARCHAR(50),
    summary    VARCHAR(500),
    deadline   timestamp NOT NULL,
    status     VARCHAR(50)
);
-- export POSTGRESQL_URL='postgres://husanmusa:pass@localhost:5432/tododb?sslmode=disable'
-- migrate -database ${POSTGRESQL_URL} -path migrations up/down