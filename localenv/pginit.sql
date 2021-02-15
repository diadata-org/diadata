-- Table asset is the single source of truth for all assets handled at DIA.
CREATE TABLE asset (
    asset_id integer primary key generated always as identity,
    symbol text not null,
    name text not null,    
    decimals text,
    blockchain text,
    address text not null,
    UNIQUE (address, blockchain)
);

-- Table exchangepair holds all trading pairs for the pair scrapers.
-- The format has to be the same as emitted by the exchange's API in order
-- for the pair scrapers to be able to scrape trading data from the API.
CREATE TABLE exchangepair (
    exchangepair_id integer primary key generated always as identity,
    symbol text not null,
    foreignname text not null,
    exchange text not null,
    UNIQUE (foreignname, exchange),
    -- These fields reference asset table and should be verified by pairdiscoveryservice.
    -- Only trades with verified pairs are processed further and thereby enter price calculation.
    verified boolean default false,
    id_basetoken integer REFERENCES asset(asset_id),
    id_quotetoken integer REFERENCES asset(asset_id)
);

-- blockchain table stores all blockchains available in our databases
CREATE TABLE blockchain (
    blockchain_id integer primary key generated always as identity,
    name text not null,
    genesisdate timestamp,
    nativetoken text,
	verificationmechanism text
);