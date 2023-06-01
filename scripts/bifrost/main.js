const { ApiPromise, WsProvider } = require('@polkadot/api');
const { options } = require('@bifrost-finance/api');
const { default: BigNumber } = require('bignumber.js');

async function main() {
  const wsProvider = new WsProvider('wss://bifrost-rpc.liebi.com/ws');
  const api = await ApiPromise.create(options({
    provider: wsProvider
  }));

  const pair_statuses_native = await api.query.zenlinkProtocol.pairStatuses.entries();
  const pair_statuses = pair_statuses_native.
    map((item) => {
      let out = { pair1: {}, pair2: {}, pairAccount: '' };
      let trade = JSON.parse(JSON.stringify(item[0].toHuman()));
      out.pair1.chainId = Number(trade[0][0].chainId.replace(',', ''))
      out.pair1.assetType = Number(trade[0][0].assetType.replace(',', ''))
      out.pair1.assetIndex = Number(trade[0][0].assetIndex.replace(',', ''))
      out.pair2.chainId = Number(trade[0][1].chainId.replace(',', ''))
      out.pair2.assetType = Number(trade[0][1].assetType.replace(',', ''))
      out.pair2.assetIndex = Number(trade[0][1].assetIndex.replace(',', ''))
      if (item[1].toJSON().trading != undefined) {
        out.pairAccount = item[1].toJSON().trading.pairAccount
      }
      return out
    }
    )

  const getPairAccount = (from, to) => {
    let pair1 = JSON.parse(from)
    let pair2 = JSON.parse(to)
    let result = pair_statuses.filter((item) => {
      if (item.pair1.chainId == pair1.chainId &&
        item.pair1.assetType == pair1.assetType &&
        item.pair1.assetIndex == pair1.assetIndex &&
        item.pair2.chainId == pair2.chainId &&
        item.pair2.assetType == pair2.assetType &&
        item.pair2.assetIndex == pair2.assetIndex) {
        return item.pairAccount
      }
    }
    ).map((item) => {
      return item.pairAccount
    })
    return result[0]
  }

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
        let to_native = event.data[2][0];
        let from_native = event.data[2].pop();
        let from = getTokenByPair(from_native).token;
        let to = getTokenByPair(to_native).token;
        let fromDecimals = getTokenByPair(from_native).decimals;
        let toDecimals = getTokenByPair(to_native).decimals;
        let out = `${to}-${from} ${to} ${from}`
        out += ` ${BigNumber(asset_balance[0]).dividedBy(10 ** toDecimals)} ${BigNumber(asset_balance.pop()).dividedBy(10 ** fromDecimals)}`;
        out += ` ${JSON.parse(to_native).assetIndex} ${JSON.parse(from_native).assetIndex}`;
        out += ` ${header.number}-${phase.asApplyExtrinsic}`;
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
      return { type: 'Native', token: 'BNC', decimals: 12 };
    case '2001-2-519':
      return { type: 'Token', token: 'ZLK', decimals: 18 };
    case '2001-2-1028':
      return { type: 'vsToken', token: 'vsKSM', decimals: 12 };
    case '2001-2-1284':
      return { type: 'vsToken', token: 'vsBond(KSM)', decimals: 12 };
    case '2001-2-1283':
      return { type: 'vsToken', token: 'vsBond(DOT)', decimals: 10 };
    case '2001-2-260':
      return { type: 'vToken', token: 'vKSM', decimals: 12 };
    case '2001-2-2050':
      return { type: 'Token', token: 'kBTC' };
    case '2001-2-266':
      return { type: 'vToken', token: 'vMOVR', decimals: 18 };
    case '2001-2-264':
      return { type: 'vToken', token: 'vPHA', decimals: 12 };
    case '2001-2-257':
      return { type: 'vToken', token: 'vBNC', decimals: 12 };
    case '2001-2-516':
      return { type: 'Token', token: 'KSM', decimals: 12 };
    case '2001-2-518':
      return { type: 'Token', token: 'KAR', decimals: 12 };
    case '2001-2-770':
      return { type: 'Token', token: 'aUSD' };
    case '2001-2-521':
      return { type: 'Token', token: 'RMRK', decimals: 10 };
    case '2001-2-522':
      return { type: 'Token', token: 'MOVR', decimals: 18 };
    case '2001-2-520':
      return { type: 'Token', token: 'PHA', decimals: 12 };
    case '2030-0-0':
      return { type: 'Token', token: 'BNC', decimals: 12 };
    case '2030-2-2048':
      return { type: 'Token', token: 'DOT', decimals: 10 };
    case '2030-2-2304':
      return { type: 'vToken', token: 'vDOT', decimals: 10 };
    case '2030-2-2560':
      return { type: 'vsToken', token: 'vsDOT', decimals: 10 };
    case '2030-2-2049':
      return { type: 'Token', token: 'GLMR', decimals: 18 };
    case '2030-2-2052':
      return { type: 'Token', token: 'FIL' };
    case '2001-2-2048':
      return { type: 'Token', token: 'USDT', decimals: 6 };
    case '2030-2-2305':
      return { type: 'vToken', token: 'vGLMR', decimals: 18 };
    case '2030-2-2308':
      return { type: 'vToken', token: 'vFIL' };
    default:
      return null;
  }
}
