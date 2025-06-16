-- +goose Up
-- +goose StatementBegin
create table goods_logs (
    Id Int32,
    ProjectId Int32,
    Name String,
    Description String,
    Priority Int32,
    Removed Bool ,
    EventType DATETIME,

    index idx_id (Id) Type bloom_filter GRANULARITY 1,
    index idx_project_id (ProjectId) Type bloom_filter GRANULARITY 1,
    index idx_name (Name) Type bloom_filter GRANULARITY 1
)
ENGINE = MergeTree()
ORDER BY (ProjectId, Id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table goods_logs;
-- +goose StatementEnd
