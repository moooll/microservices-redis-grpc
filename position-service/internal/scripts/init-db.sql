create user $(DB_USER);
create database $(DB_NAME);
grant all privileges on $(DB_NAME) to $(DB_USER);

drop table positions;
create table positions(
    id uuid primary key,
    server_id smallint,
    company_name varchar,
    open boolean,
    buy_price real,
    sell_price real,
    pnl real
);

create or replace function notify_posit() 
returns trigger 
LANGUAGE 'plpgsql'
as $$
BEGIN
    PERFORM (
        with payload(id, server_id, company_name, open, buy_price, sell_price, pnl) as (
        select NEW."id",
        NEW."server_id",
        NEW."company_name",
        NEW."open",
        NEW."buy_price",
        NEW."sell_price",
        NEW."pnl"
    )
    select pg_notify('notification', row_to_json(payload)::text) from payload
    );
    RETURN null;
END;
$$;

drop trigger posit_upd on positions;
CREATE trigger posit_upd 
after insert or update
on positions
for each row 
execute procedure notify_posit();
