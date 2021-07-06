CREATE EXTENSION "pgcrypto";


-- Table asset is the single source of truth for all assets handled at DIA.
-- If a field is not case sensitive (such as address for Ethereum) it should
-- be all lowercase for consistency reasons.
-- Otherwise it must be as defined in the underlying contract.
CREATE TABLE asset (
    asset_id UUID DEFAULT gen_random_uuid(),
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
    symbol text not null,
    foreignname text not null,
    exchange text not null,
    UNIQUE (foreignname, exchange),
    -- These fields reference asset table and should be verified by pairdiscoveryservice.
    -- Only trades with verified pairs are processed further and thereby enter price calculation.
    verified boolean default false,
    id_quotetoken uuid REFERENCES asset(asset_id),
    id_basetoken uuid REFERENCES asset(asset_id)
);

CREATE TABLE exchangesymbol (
    exchangesymbol_id UUID DEFAULT gen_random_uuid(),
    symbol text not null,
    exchange text not null,
    UNIQUE (symbol,exchange),
    verified boolean default false,
    asset_id uuid REFERENCES asset(asset_id)
);

-- blockchain table stores all blockchains available in our databases
CREATE TABLE blockchain (
    blockchain_id integer primary key generated always as identity,
    name text not null,
    genesisdate timestamp,
    nativetoken text,
	verificationmechanism text,
    UNIQUE(name)
);


---------------------------------------
------- tables for NFT storage --------
---------------------------------------

-- collect all possible categories for nfts
CREATE TABLE nftcategory (
    category_id UUID DEFAULT gen_random_uuid(),
    category text not null,
    UNIQUE(category)
);

-- nftclass is uniquely defined by the pair (blockchain,address),
-- referring to the blockchain on which the nft was minted.
CREATE TABLE nftclass (
    nftclass_id UUID DEFAULT gen_random_uuid(),
    address text not null,
    symbol text,
    name text,
    blockchain text REFERENCES blockchain(name),
    contract_type text,
    category text REFERENCES nftcategory(category),
    UNIQUE(blockchain,address),
    UNIQUE(nftclass_id)
);

-- an element from nft is a specific non-fungible nft, unqiuely
-- identified by the pair (address(on blockchain), tokenID)
CREATE TABLE nft (
    nft_id UUID DEFAULT gen_random_uuid(),
    nftclass_id uuid REFERENCES nftclass(nftclass_id),
    tokenID text not null,
    creation_time timestamp,
    creator_address text,
    uri text,
    attributes jsonb,
    UNIQUE(nftclass_id, tokenID),
    UNIQUE(nft_id)
);

CREATE TABLE nfttrade (
    sale_id UUID DEFAULT gen_random_uuid(),
    nftclass_id uuid REFERENCES nftclass(nftclass_id),
    nft_id uuid REFERENCES nft(nft_id),
    price text,
    price_usd numeric,
    transfer_from text,
    transfer_to text,
    currency_symbol text,
    currency_address text,
    currency_decimals numeric,
    block_number numeric,
    trade_time timestamp,
    tx_hash text,    
    marketplace text,
    UNIQUE(sale_id),
    UNIQUE(nft_id, trade_time)
);

CREATE TABLE nftbid (
    bid_id UUID DEFAULT gen_random_uuid(),
    nft_id uuid REFERENCES nft(nft_id),
    bid_value text,
    from_address text,
    currency_symbol text,
    currency_address text,
    currency_decimals numeric,
    blocknumber numeric,
    blockposition numeric,
    bid_time timestamp,
    tx_hash text,
    marketplace text,
    UNIQUE(bid_id),
    UNIQUE(nft_id, from_address, bid_time)
);


CREATE TABLE nftoffer (
    offer_id UUID DEFAULT gen_random_uuid(),
    nft_id uuid REFERENCES nft(nft_id),
    start_value text,
    end_value text,
    duration numeric,
    from_address text,
    auction_type text,
    currency_symbol text,
    currency_address text,
    currency_decimals numeric,
    blocknumber numeric,
    blockposition numeric,
    offer_time timestamp,
    tx_hash text,
    marketplace text,
    UNIQUE(offer_id),
    UNIQUE(nft_id, from_address, offer_time)
);

CREATE TABLE IF NOT EXISTS scrapers (
    name character varying(255) NOT NULL,
	conf json,
	state json,
    CONSTRAINT pk_scrapers PRIMARY KEY(name)
);

CREATE TABLE blockdata (
    blockdata_id UUID DEFAULT gen_random_uuid(),
    blockchain text not null,
    block_number text not null,
    block_data jsonb,
    UNIQUE(blockchain, block_number),
    UNIQUE(blockdata_id)
);
