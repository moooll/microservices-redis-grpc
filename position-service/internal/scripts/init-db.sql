create user position_admin;
create database positions;
grant all privileges on  positions to position_admin;

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

truncate positions;
INSERT INTO positions(
    id,
    server_id,
    company_name,
    open,
    buy_price,
    sell_price,
    pnl) 
    VALUES('805fe1cc-4702-4dc8-a20a-d35cbfc3faa8', 1, 'apple', true, 145.0, 0, 0.05);