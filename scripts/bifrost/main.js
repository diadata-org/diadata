const { ApiPromise, WsProvider } = require('@polkadot/api');
const { options } = require('@bifrost-finance/api');

async function main() {
  const wsProvider = new WsProvider('wss://bifrost-rpc.liebi.com/ws');
  const api = await ApiPromise.create(options({
    provider: wsProvider
  }));

  await api.rpc.chain.subscribeNewHeads(async (header) => {
    console.log(`blockHeight: ${header.number}`);
    const blockHash = await api.rpc.chain.getBlockHash(header.number);
    const at = await api.at(blockHash);
    const events = await at.query.system.events();
    events.filter((record) => {
      return record.event.section === 'zenlinkProtocol' && event.method === 'AssetSwap';
    })
      .map(async (record) => {
        const { event, phase } = record;
        let asset_balance = event.data[3];
        let from = getTokenByPair(event.data[2][0]).token
        let to = getTokenByPair(event.data[2].pop()).token
        if (from == "vKSM" && to == "KSM" || from == "KSM" && to == "vKSM") {
          console.log(`Trade:${from}-${to}`, from, to, JSON.stringify(asset_balance[0]),
            JSON.stringify(asset_balance.pop()), `${header.number}-${phase.asApplyExtrinsic}`)
        }
      });
  });
}

main().catch((error) => {
  console.error(error);
  process.exit(-1);
});

const getTokenByPair = (value) => {
  let item = JSON.parse(value)
  switch (`${item.chainId}-${item.assetType}-${item.assetIndex}`) {
    case '2001-0-0':
      return { type: 'Native', token: 'BNC' };
    case '2001-2-519':
      return { type: 'Token', token: 'ZLK' };
    case '2001-2-1028':
      return { type: 'vsToken', token: 'vsKSM' };
    case '2001-2-1284':
      return { type: 'vsToken', token: 'vsBond(KSM)' };
    case '2001-2-1283':
      return { type: 'vsToken', token: 'vsBond(DOT)' };
    case '2001-2-260':
      return { type: 'vToken', token: 'vKSM' };
    case '2001-2-2050':
      return { type: 'Token', token: 'kBTC' };
    case '2001-2-266':
      return { type: 'vToken', token: 'vMOVR' };
    case '2001-2-264':
      return { type: 'vToken', token: 'vPHA' };
    case '2001-2-257':
      return { type: 'vToken', token: 'vBNC' };
    case '2001-2-516':
      return { type: 'Token', token: 'KSM' };
    case '2001-2-518':
      return { type: 'Token', token: 'KAR' };
    case '2001-2-770':
      return { type: 'Token', token: 'aUSD' };
    case '2001-2-521':
      return { type: 'Token', token: 'RMRK' };
    case '2001-2-522':
      return { type: 'Token', token: 'MOVR' };
    case '2001-2-520':
      return { type: 'Token', token: 'PHA' };
    case '2030-0-0':
      return { type: 'Token', token: 'BNC' };
    case '2030-2-2048':
      return { type: 'Token', token: 'DOT' };
    case '2030-2-2304':
      return { type: 'vToken', token: 'vDOT' };
    case '2030-2-2560':
      return { type: 'vsToken', token: 'vsDOT' };
    case '2030-2-2049':
      return { type: 'Token', token: 'GLMR' };
    case '2030-2-2052':
      return { type: 'Token', token: 'FIL' };
    case '2001-2-2048':
      return { type: 'Token', token: 'USDT' };
    case '2030-2-2305':
      return { type: 'vToken', token: 'vGLMR' };
    case '2030-2-2308':
      return { type: 'vToken', token: 'vFIL' };
    default:
      return null;
  }
}