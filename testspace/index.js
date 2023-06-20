const { ApiPromise, WsProvider } = require("@polkadot/api");

async function queryInfo(api) {
  const assetMetadata = await api.query.assetRegistry.assetMetadatas.entries();

  assetMetadata.map((asset) => {
    let h = asset[1].toHuman();
    console.log(JSON.stringify(h));
  });
}

async function setup() {
  const wsProvider = new WsProvider("wss://acala-polkadot.api.onfinality.io/public-ws");
  const api = await ApiPromise.create({
    provider: wsProvider,
  });

  await queryInfo(api);
  await api.disconnect();
}

setup();