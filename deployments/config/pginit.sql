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

-- polling table stores data - required for HTTP polling
CREATE TABLE polling (
    polling_id UUID DEFAULT gen_random_uuid(),
    blockchain text,
    contract_address text NOT NULL,
    page numeric DEFAULT 1,
    UNIQUE(blockchain, contract_address)
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

ALTER TABLE oracleconfig  ADD COLUMN name VARCHAR(255);
ALTER TABLE oracleconfig  ADD COLUMN draft boolean DEFAULT true;
ALTER TABLE oracleconfig  ADD COLUMN customer_id int ;
ALTER TABLE oracleconfig  ADD COLUMN billable boolean DEFAULT false ;


ALTER TABLE oracleconfig
ADD CONSTRAINT unique_customer_chainid_address 
UNIQUE (customer_id, chainid, address);



ALTER TABLE oracleconfig
ADD CONSTRAINT unique_feeder_id UNIQUE (feeder_id);


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







 CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    account_creation_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    customer_plan INTEGER REFERENCES plans(plan_id) ON DELETE SET NULL,
    deployed_oracles INTEGER DEFAULT 0,
    payment_status VARCHAR(50),
    last_payment TIMESTAMP,
    payment_source VARCHAR(255),
    number_of_data_feeds INTEGER DEFAULT 0,
    active BOOLEAN DEFAULT TRUE
);

CREATE TABLE wallet_public_keys (
    key_id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES customers(customer_id) ON DELETE CASCADE,
    public_key TEXT NOT NULL,
    access_level VARCHAR(50) NOT NULL DEFAULT 'read_write',
    UNIQUE (public_key),
    CONSTRAINT check_access_level CHECK (access_level IN ('read', 'read_write'))

);



ALTER TABLE wallet_public_keys ADD COLUMN username VARCHAR(255) UNIQUE;

CREATE TABLE wallet_public_keys_temp (
    key_id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES customers(customer_id) ON DELETE CASCADE,
    public_key TEXT NOT NULL,
    access_level VARCHAR(50) NOT NULL DEFAULT 'read_write',
    username VARCHAR(255)
 );

 ALTER TABLE wallet_public_keys_temp
ADD CONSTRAINT unique_customer_public_key UNIQUE (customer_id, public_key);


CREATE TABLE transfer_created (
    event VARCHAR(50),
    transaction VARCHAR(50),
    network_id INT,
    network_name VARCHAR(50),
    contract_address VARCHAR(42),
    email VARCHAR(255),
    company TEXT,
    parent VARCHAR(50),
    id UUID,
    invoice_id VARCHAR(50),
    bill_date TIMESTAMP,
    to_address VARCHAR(42),
    from_address VARCHAR(42),
    token_symbol VARCHAR(10),
    token_address VARCHAR(42),
    payment_type VARCHAR(50),
    usd BOOLEAN,
    amount NUMERIC(10, 2),
    item VARCHAR(255),
    item_id INT,
    source VARCHAR(50),
    batch_id UUID,
    transfer_id UUID,
    ref_id VARCHAR(255),
    agreement_id UUID
);

CREATE TABLE loop_payment_transfer_processed (
    event               VARCHAR(255) NOT NULL,
    transaction         VARCHAR(255) NOT NULL,
    network_id          INTEGER NOT NULL,
    network_name        VARCHAR(255) NOT NULL,
    contract_address    VARCHAR(255) NOT NULL,
    email               VARCHAR(255),
    company             VARCHAR(255),
    parent              VARCHAR(255),
    transfer_id         VARCHAR(255) NOT NULL,
    success             BOOLEAN NOT NULL,
    payment_token_address VARCHAR(255),
    payment_token_symbol VARCHAR(255),
    end_user            VARCHAR(255),
    reason              VARCHAR(255),
    invoice_id          VARCHAR(255),
    amount_paid         DOUBLE PRECISION,
    agreement_id        VARCHAR(255),
    ref_id              VARCHAR(255),
    batch_id            VARCHAR(255),
    usd_amount          VARCHAR(255)
);

CREATE TABLE loop_payment_responses (
    id SERIAL PRIMARY KEY,
    event TEXT,
    transaction TEXT,
    network_id INT,
    network_name TEXT,
    contract_address TEXT,
    email TEXT,
    company TEXT,
    parent TEXT,
    subscriber TEXT,
    item TEXT,
    item_id TEXT,
    agreement_id TEXT,
    agreement_amount TEXT,
    frequency_number INT,
    frequency_unit TEXT,
    add_on_agreements TEXT,
    add_on_items TEXT,
    add_on_item_ids TEXT,
    add_on_total_amount TEXT,
    payment_token_symbol TEXT,
    payment_token_address TEXT,
    event_date INT,
    ref_id TEXT,
    invoice_id TEXT,
    metadata JSONB DEFAULT '{}'::jsonb
);


CREATE TABLE plans (
    plan_id SERIAL PRIMARY KEY,
    plan_name VARCHAR(50) NOT NULL UNIQUE,
    plan_description TEXT,
    plan_price NUMERIC(10, 2) NOT NULL,
    plan_features TEXT
);

ALTER TABLE plans ADD COLUMN total_feeds integer default 3;
ALTER TABLE plans ADD COLUMN total_oracles integer default 3;





INSERT INTO  "plans"("plan_id","plan_name","plan_description","plan_price","plan_features","total_feeds")
VALUES
(1,E'Plan 2',E'default',0,E'desc',10);

INSERT INTO  "plans"("plan_id","plan_name","plan_description","plan_price","plan_features","total_feeds")
VALUES
(2,E'Plan 1',E'default',0,E'desc',3);

