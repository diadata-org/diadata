CREATE EXTENSION "pgcrypto";


-- Table asset is the single source of truth for all assets handled at DIA.
-- If a field is not case sensitive (such as address for Ethereum) it should
-- be all lowercase for consistency reasons.
-- Otherwise it must be as defined in the underlying contract.
CREATE TABLE asset (
    asset_id UUID DEFAULT gen_random_uuid(),
    symbol text NOT NULL,
    name text NOT NULL,
    decimals text,
    blockchain text,
    address text NOT NULL,
    UNIQUE (asset_id),
    UNIQUE (address, blockchain)
);

-- Table exchangepair holds all trading pairs for the pair scrapers.
-- The format has to be the same as emitted by the exchange's API in order
-- for the pair scrapers to be able to scrape trading data from the API.
CREATE TABLE exchangepair (
    exchangepair_id UUID DEFAULT gen_random_uuid(),
    symbol text NOT NULL,
    foreignname text NOT NULL,
    exchange text NOT NULL,
    UNIQUE (foreignname, exchange),
    -- These fields reference asset table and should be verified by pairdiscoveryservice.
    -- Only trades with verified pairs are processed further and thereby enter price calculation.
    verified boolean default false,
    id_quotetoken UUID REFERENCES asset(asset_id),
    id_basetoken UUID REFERENCES asset(asset_id)
);

CREATE TABLE exchangesymbol (
    exchangesymbol_id UUID DEFAULT gen_random_uuid(),
    symbol text NOT NULL,
    exchange text NOT NULL,
    UNIQUE (symbol,exchange),
    verified boolean default false,
    asset_id UUID REFERENCES asset(asset_id)
);

CREATE TABLE exchange (
    exchange_id UUID DEFAULT gen_random_uuid(),
    name text NOT NULL,
    centralized boolean default false,
    bridge boolean default false,
    contract text,
    blockchain text,
    rest_api text,
    ws_api text,
    pairs_api text,
    watchdog_delay numeric NOT NULL,
    scraper_active boolean,
    UNIQUE(exchange_id),
    UNIQUE (name)
);

CREATE TABLE pool (
    pool_id UUID DEFAULT gen_random_uuid(),
    exchange text NOT NULL,
    blockchain text NOT NULL,
    address text NOT NULL,
    UNIQUE (pool_id),
    UNIQUE (blockchain,address)
);

CREATE TABLE poolasset (
    poolasset_id UUID DEFAULT gen_random_uuid(),
    pool_id UUID REFERENCES pool(pool_id) NOT NULL,
    asset_id UUID REFERENCES asset(asset_id) NOT NULL, 
    liquidity numeric,
    liquidity_usd numeric,
    time_stamp timestamp,
    token_index integer,
    UNIQUE (poolasset_id),
    UNIQUE(pool_id,asset_id)
);

CREATE TABLE scraper_cronjob_state (
    scraper_cronjob_state_id UUID DEFAULT gen_random_uuid(),
    scraper text NOT NULL,
    index_type text NOT NULL,
    index_value numeric,
    UNIQUE(scraper_cronjob_state_id),
    UNIQUE(scraper,index_type)
);

CREATE TABLE chainconfig (
    chain_config_id UUID DEFAULT gen_random_uuid(),
    rpcurl text NOT NULL,
    wsurl text NOT NULL,
    chainID text NOT NULL,
    UNIQUE (chainID)
);

-- blockchain table stores all blockchains available in our databases
CREATE TABLE blockchain (
    blockchain_id UUID DEFAULT gen_random_uuid(),
    name text NOT NULL,
    genesisdate numeric,
    nativetoken_id UUID REFERENCES asset(asset_id),
	verificationmechanism text,
    chain_id text,
    UNIQUE(blockchain_id),
    UNIQUE(name)
);

CREATE TABLE assetvolume (
    asset_id UUID primary key,
    volume decimal,
    time_stamp timestamp
);

---------------------------------------
------- tables for NFT storage --------
---------------------------------------

-- collect all possible categories for nfts
CREATE TABLE nftcategory (
    category_id UUID DEFAULT gen_random_uuid(),
    category text NOT NULL,
    UNIQUE(category)
);

-- nftclass is uniquely defined by the pair (blockchain,address),
-- referring to the blockchain on which the nft was minted.
CREATE TABLE nftclass (
    nftclass_id UUID DEFAULT gen_random_uuid(),
    blockchain text NOT NULL,
    address text NOT NULL,
    symbol text,
    name text,
    contract_type text,
    category text REFERENCES nftcategory(category),
    UNIQUE(blockchain, address),
    UNIQUE(nftclass_id)
);

-- historicalquotation collects USD quotes with lower frequency
-- for a selection of assets.
CREATE TABLE historicalquotation (
    historicalquotation_id UUID DEFAULT gen_random_uuid(),
    asset_id UUID REFERENCES asset(asset_id) NOT NULL, 
    price numeric,
    quote_time timestamp,
    source text,
    UNIQUE(asset_id,quote_time,source),
    UNIQUE(historicalquotation_id)
);

-- an element from nft is a specific non-fungible nft, unqiuely
-- identified by the pair (address(on blockchain), token_id)
CREATE TABLE nft (
    nft_id UUID DEFAULT gen_random_uuid(),
    nftclass_id UUID REFERENCES nftclass(nftclass_id),
    token_id text NOT NULL,
    creation_time timestamp,
    creator_address text,
    uri text,
    attributes jsonb,
    UNIQUE(nftclass_id, token_id),
    UNIQUE(nft_id)
);

CREATE TABLE nfttradecurrent (
    sale_id UUID DEFAULT gen_random_uuid(),
    nftclass_id UUID REFERENCES nftclass(nftclass_id),
    nft_id UUID REFERENCES nft(nft_id),
    price text,
    price_usd numeric,
    transfer_from text,
    transfer_to text,
    currency_symbol text,
    currency_address text,
    currency_decimals numeric,
    currency_id UUID REFERENCES asset(asset_id),
    bundle_sale boolean default false,
    block_number numeric,
    trade_time timestamp,
    tx_hash text,    
    marketplace text,
    UNIQUE(sale_id),
    UNIQUE(nft_id, trade_time)
);

CREATE TABLE nftbid (
    bid_id UUID DEFAULT gen_random_uuid(),
    nft_id UUID REFERENCES nft(nft_id),
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
    nft_id UUID REFERENCES nft(nft_id),
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
    blockchain text NOT NULL,
    block_number numeric NOT NULL,
    block_data jsonb,
    UNIQUE(blockchain, block_number),
    UNIQUE(blockdata_id)
);

CREATE TABLE assetpriceident (
    priceident_id UUID DEFAULT gen_random_uuid(),
    asset_id UUID REFERENCES asset(asset_id),
    group_id numeric NOT NULL,
    rank_in_group numeric NOT NULL,
    UNIQUE(asset_id),
    UNIQUE(group_id, rank_in_group)
);

CREATE TABLE synthassetdata (
    synthassetdata_id UUID DEFAULT gen_random_uuid(),
    synthasset_id UUID REFERENCES asset(asset_id),
    underlying_id UUID REFERENCES asset(asset_id),
    supply numeric,
    locked_underlying numeric,
    num_mints numeric,
    num_redeems numeric,
    block_number numeric,
    time_stamp timestamp,
    UNIQUE(synthassetdata_id),
    UNIQUE(synthasset_id,time_stamp)
);

CREATE TABLE nftexchange (
    exchange_id UUID DEFAULT gen_random_uuid(),
    name text NOT NULL,
    centralized boolean default false,
    contract text,
    blockchain text,
    rest_api text,
    ws_api text,
    watchdog_delay numeric NOT NULL,
    UNIQUE(exchange_id),
    UNIQUE (name)
);

CREATE TABLE oracleconfig (
    id uuid DEFAULT gen_random_uuid(),
    address text NOT NULL,
    feeder_id text NOT NULL,
    owner text NOT NULL,
    symbols text NOT NULL,
    chainid text NOT NULL,
    active boolean DEFAULT true,
    frequency text,
    sleepseconds text,
    deviationpermille text,
    blockchainnode text DEFAULT ''::text,
    feeder_address text,
    mandatory_frequency text,
    deleted boolean DEFAULT false,
    createddate timestamp without time zone DEFAULT now() NOT NULL,
    lastupdate timestamp without time zone,
    creation_block bigint,
    creation_block_time timestamp without time zone DEFAULT '1970-01-01 00:00:00'::timestamp without time zone,
    feedselection text,
    expired boolean DEFAULT false,
    expired_time timestamp without time zone DEFAULT '1970-01-01 00:00:00'::timestamp without time zone
);



-- CREATE TABLE oracleconfig (
--     id UUID DEFAULT gen_random_uuid(),
--     address text NOT NULL,
--     feeder_id text NOT NULL,
--     owner text NOT NULL,
--     symbols text NOT NULL,
--     feeder_address text NOT NULL,
--     chainID text NOT NULL,
--     active  boolean default true,
--     deleted  boolean default false,
--     frequency text ,
--     sleepseconds text,
--     deviationpermille text,
--     blockchainnode text,
--     mandatory_frequency text,
--     createddate TIMESTAMP NOT NULL DEFAULT NOW(),
--     lastupdate TIMESTAMP NOT NULL,
--     UNIQUE (id),
--     UNIQUE (feeder_id)
-- );

-- ALTER TABLE oracleconfig  ADD COLUMN creation_block_time TIMESTAMP DEFAULT 'epoch'::timestamp;
-- ALTER TABLE oracleconfig  ADD COLUMN feedSelection TEXT ;
-- ALTER TABLE oracleconfig  ADD COLUMN expired boolean default false ;
-- ALTER TABLE oracleconfig  ADD COLUMN expired_time TIMESTAMP DEFAULT 'epoch'::timestamp;






         



CREATE TABLE feederresource (
    id  SERIAL PRIMARY KEY,
    owner text NOT NULL,
    total numeric NOT NULL,
    UNIQUE (id),
    UNIQUE (owner)
);

CREATE TABLE asset_list (
    id SERIAL PRIMARY KEY,
    asset_name VARCHAR(255) NOT NULL,
    custom_name VARCHAR(255),
    symbol VARCHAR(50),
    methodology TEXT,
    list_name TEXT
    
);

CREATE TABLE exchange_list (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    asset_id INT REFERENCES asset_list(id) ON DELETE CASCADE
);

CREATE TABLE exchange_pairs (
    id SERIAL PRIMARY KEY,
    exchange_id INT REFERENCES exchange_list(id) ON DELETE CASCADE,
    pair VARCHAR(255) NOT NULL
);


