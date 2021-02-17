CREATE EXTENSION "pgcrypto";


-- Table asset is the single source of truth for all assets handled at DIA.
CREATE TABLE asset (
    asset_id UUID DEFAULT gen_random_uuid(),
    -- I think we can stick to integer primary key. The table is going to hold at most order 10^5 entries.
    -- asset_id integer primary key generated always as identity,
    symbol text not null,
    name text not null,
    decimals text,
    blockchain text,
    address text not null,
    UNIQUE (asset_id),
    UNIQUE (address, blockchain)
);

-- Table exchangepair holds all trading pairs for the pair scrapers.
-- The format has to be the same as emitted by the exchange's API in order
-- for the pair scrapers to be able to scrape trading data from the API.
CREATE TABLE exchangepair (
    exchangepair_id UUID DEFAULT gen_random_uuid(),
    -- exchangepair_id integer primary key generated always as identity,
    symbol text not null,
    foreignname text not null,
    exchange text not null,
    UNIQUE (foreignname, exchange),
    -- These fields reference asset table and should be verified by pairdiscoveryservice.
    -- Only trades with verified pairs are processed further and thereby enter price calculation.
    verified boolean default false,
    id_basetoken uuid REFERENCES asset(asset_id),
    id_quotetoken uuid REFERENCES asset(asset_id)
);

-- blockchain table stores all blockchains available in our databases
CREATE TABLE blockchain (
    blockchain_id integer primary key generated always as identity,
    name text not null,
    genesisdate timestamp,
    nativetoken text,
	verificationmechanism text
);

-- Comments:

-- We use text instead of char:
-- https://stackoverflow.com/questions/4848964/difference-between-text-and-varchar-character-varying

-- For the use of primary keys and the advantage of integers. A bit old, but maybe still true?
-- https://stackoverflow.com/questions/337503/whats-the-best-practice-for-primary-keys-in-tables