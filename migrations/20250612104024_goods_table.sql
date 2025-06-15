-- +goose Up
-- +goose StatementBegin
create table projects(
    id int generated always as identity primary key,
    name varchar(255) not null,
    created_at timestamp not null default now()
);

create table goods(
    id int generated always as identity primary key,
    project_id int not null,
    name varchar(255) not null,
    description text,
    priority int not null,
    removed bool not null default false,
    created_at timestamp not null default now(),
    foreign key (project_id) references projects(id)
);

create index idx_goods_project_id on goods(project_id);
create index idx_goods_name on goods(name);


create function set_goods_priority() returns trigger as $$
begin
    select coalesce(max(priority),0) + 1 into new.priority
    from goods
    where project_id = new.project_id;
    return new;
end;
$$ language plpgsql;

create trigger trg_goods_priority before insert on goods
    for each row execute function set_goods_priority();


insert into projects(name) values ('Первая запись');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop trigger trg_goods_priority on goods;
drop function set_goods_priority();

drop index idx_goods_name;
drop index idx_goods_project_id;

drop table goods;
drop table projects;
-- +goose StatementEnd
