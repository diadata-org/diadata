# GSRPC Registry
The GSRPC Registry can parse target metadata information into an in-memory registry of complex structures. 

By leveraging the on-chain metadata, GSRPC is more robust to changes on types, allowing clients to only keep updated the types that are relevant to their business operation.

This registry can be used afterwards to decode data read from live chains (events & extrinsics).

## How to parse events and their fields
First we instantiate the API with the client node and open a connection: 
```go
testURL := "wss://fullnode.parachain.centrifuge.io" // Your endpoint
api, err := gsrpc.NewSubstrateAPI(testURL)

if err != nil {
    log.Printf("Couldn't connect to '%s': %s\n", testURL, err)
    return
}
```
Then we instantiate the Event Retriever logic which internally creates a new EventRegistry reading from the target metadata of the connected chain. We pass as well the state RPC so the storage API is available: 
```go
retriever, err := NewDefaultEventRetriever(state.NewEventProvider(api.RPC.State), api.RPC.State)

if err != nil {
    log.Printf("Couldn't create event retriever: %s", err)
    return
}
```
At this point what we need is a block hash to read the events within. In this example we get the latest block header and the correspondent block hash out of the block number:
```go
header, err := api.RPC.Chain.GetHeaderLatest()

if err != nil {
    log.Printf("Couldn't get latest header for '%s': %s\n", testURL, err)
    return
}

blockHash, err := api.RPC.Chain.GetBlockHash(uint64(header.Number))

if err != nil {
    log.Printf("Couldn't retrieve blockHash for '%s', block number %d: %s\n", testURL, header.Number, err)
    return
}
```
Finally, we just use the retriever function to read all the events in that block based on the chain metadata loaded in the event registry: 
```go
events, err := retriever.GetEvents(blockHash)

if err != nil {
    log.Printf("Couldn't retrieve events for '%s', block number %d: %s\n", testURL, header.Number, err)
    return
}

log.Printf("Found %d events for '%s', at block number %d.\n", len(events), testURL, header.Number)

// Example of the events returned structure
for _, event := range events {
    log.Printf("Event ID: %x \n", event.EventID)
    log.Printf("Event Name: %s \n", event.Name)
    log.Printf("Event Fields Count: %d \n", len(event.Fields))
    for k, v := range event.Fields {
        log.Printf("Field Name: %s \n", k)
        log.Printf("Field Type: %v \n", reflect.TypeOf(v))
        log.Printf("Field Value: %v \n", v)
    }
}

```

## Extended Usage
Since docs get outdated fairly quick, here are links to tests that will always be up-to-date.
### Populate Call, Error & Events Registries, Extrinsic Decoder
[Factory tests](factory_test.go)
[Decoder tests](decoder_test.go)

### Event retriever
[TestLive_EventRetriever_GetEvents](retriever/event_retriever_live_test.go)

### Extrinsic retriever
[TestLive_ExtrinsicRetriever_GetExtrinsics](retriever/extrinsic_retriever_live_test.go)