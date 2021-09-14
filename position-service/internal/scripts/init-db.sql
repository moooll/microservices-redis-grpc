create user position_admin;
create database positions;
grant all privilleges on  positions to position_admin;


create table prices(
    id uuid primary key,
    company_name varchar,
    buy_price real,
    sell_price real
);

