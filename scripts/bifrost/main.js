const { ApiPromise, WsProvider } = require('@polkadot/api');
const { options } = require('@bifrost-finance/api');
const { default: BigNumber } = require('bignumber.js');

async function main() {
  const wsProvider = new WsProvider('wss://bifrost-rpc.liebi.com/ws');
  const api = await ApiPromise.create(options({
    provider: wsProvider
  }));

  await api.rpc.chain.subscribeNewHeads(async (header) => {
    console.log(`blockHeight: ${header.number}`);
    const blockHash = await api.rpc.chain.getBlockHash(header.number);
    const at = await api.at(blockHash);
    // const at = await api.at('0xb1cbc766d73b63f93accf36dfb96067bbed658482ac7a01015aab6c28927e7d9'); // for testing
    const events = await at.query.system.events();
    events.filter((record) => {
      return record.event.section === 'zenlinkProtocol' && record.event.method === 'AssetSwap';
    })
      .map(async (record) => {
        const { event, phase } = record;
        // We collect data as string with
        // pair toSymbol fromSymbol toAmount fromAmount toAssetIndex fromAssetIndex foreignTradeID
        let asset_balance = event.data[3];
        let assets = event.data[2];
        let to_native = JSON.parse(assets[0]);
        let from_native = JSON.parse(assets[assets.length - 1]);
        let from_asset_id = `${from_native.chainId}-${from_native.assetType}-${from_native.assetIndex}`;
        let to_asset_id = `${to_native.chainId}-${to_native.assetType}-${to_native.assetIndex}`;
        let from = getTokenByPair(from_asset_id).token;
        let to = getTokenByPair(to_asset_id).token;
        let fromDecimals = getDecimals(from);
        let toDecimals = getDecimals(to);
        let out = `${to}-${from} ${to} ${from}`
        out += ` ${BigNumber(asset_balance[0]).dividedBy(10 ** toDecimals)} ${BigNumber(asset_balance[asset_balance.length - 1]).dividedBy(10 ** fromDecimals)}`;
        out += ` ${to_asset_id} ${from_asset_id}`;
        out += ` ${header.number}-${phase.asApplyExtrinsic}`;
        console.log(`Trade:${out}`)
      });
  });
}

main().catch((error) => {
  console.error(error);
  process.exit(-1);
});

const getTokenByPair = (value) => {
  switch (value) {
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
};

const getDecimals = (type) => {
  switch (type) {
    case 'USDT':
      return 6;
    case 'kBTC':
      return 8;
    case 'ETH':
    case 'vETH':
    case 'ZLK':
    case 'FIL':
    case 'vFIL':
    case 'MOVR':
    case 'vMOVR':
    case 'GLMR':
    case 'vGLMR':
    case 'CRAB':
    case 'ASTR':
    case 'SDN':
    case 'MGX':
      return 18;

    case 'DOT':
    case 'RMRK':
    case 'vDOT':
    case 'vsDOT':
    case 'TUR':
      return 10;

    case 'BNCZLKLP':
    case 'KARZLKLP':
    case 'vKSMKSMLP':
    case 'vsKSMKSMLP':
    case 'BNCKSMLP':
    case 'KSMkUSDLP':
    case 'KSMKBTCLP':
    case 'PHA':
    case 'CSM':
    case 'KINT':
      return 12;

    default:
      return 12;
  }
};