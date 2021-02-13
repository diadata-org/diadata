-- This table is the single source of truth for all assets handled at DIA.
create table asset (
    id serial primary key,
    symbol text not null,
    name text not null,
    address text,
    decimals text,
    blockchain text,
    unique (symbol, address)
);

-- This table holds all trading pairs for the exchange scraper.
-- The format has to be the same as emitted by the exchange's API
-- in order to be able to scrape trading data from the API.
create table exchangepair (
    id serial primary key,
    symbol text not null,
    foreignname text not null,
    exchange text not null
);