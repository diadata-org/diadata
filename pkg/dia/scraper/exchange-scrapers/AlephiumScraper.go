package scrapers

// TODO get swap contract list
//	TODO foreach swap contract
//		TODO - get swap tokens pair - from DB
//			TODO - not exists - create one
//		TODO - get latest event startFrom (default 0)
//		TODO - get events for contract from startFrom, limit 100
//		TODO - no events - skip
//		TODO - filter with eventIndex=2 only
//		TODO - no events - skip
//		TODO - get fileds[1]fields[4], fields[2]fields[3] amounts
//		TODO - get token pairs foreach amount
//		TODO - write trade result to channel
//		TODO - store nextStart from event list to storage

// CREATE TABLE swap_event_request (
//     event_request_id UUID DEFAULT gen_random_uuid(),
//     blockchain text,
//     swap_contract_address text NOT NULL,
//     swap_token0_address text NOT NULL,
//     swap_token1_address text NOT NULL,
//     next_start numeric,
//     UNIQUE(blockchain),
//     UNIQUE(swap_contract_address)
// );

// where is migration logic? - I just need to add CREATE TABLE to deployments/config/pginit.sql ?
