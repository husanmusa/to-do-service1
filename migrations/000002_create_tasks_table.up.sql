alter table tasks add column created_at timestamp default current_timestamp,
            add column updated_at timestamp,
            add column deleted_at timestamp;