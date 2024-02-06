const Web3 = require("web3");

const { createResponse, getPrice } = require("./utils");
const ethers = require("ethers");
const bignumber = ethers.BigNumber;

  
const web3 = new Web3(
  new Web3.providers.HttpProvider(
    process.env.ETHEREUM_NODE_URL || "https://mainnet.infura.io/v3/2883d1b22e0e4d62b535592dd8075fee"
  )
);
let abi = [
  {
    constant: true,
    inputs: [],
    name: "getTotalPooledEther",
    outputs: [
      {
        name: "",
        type: "uint256",
      },
    ],
    payable: false,
    stateMutability: "view",
    type: "function",
  },
  {
    constant: true,
    inputs: [],
    name: "totalSupply",
    outputs: [
      {
        name: "",
        type: "uint256",
      },
    ],
    payable: false,
    stateMutability: "view",
    type: "function",
  },
];

let poolabi = [
  {
    inputs: [],
    name: "getActiveMinipoolCount",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "uint256",
        name: "offset",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "limit",
        type: "uint256",
      },
    ],
    name: "getMinipoolCountPerStatus",
    outputs: [
      {
        internalType: "uint256",
        name: "initialisedCount",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "prelaunchCount",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "stakingCount",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "withdrawableCount",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "dissolvedCount",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
];

let rocketnetwork = [
  {
    "inputs": [],
    "name": "getTotalRETHSupply",
    "outputs": [{
      "internalType": "uint256",
      "name": "",
      "type": "uint256"
    }],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "getTotalETHBalance",
    "outputs": [{
      "internalType": "uint256",
      "name": "",
      "type": "uint256"
    }],
    "stateMutability": "view",
    "type": "function"
  },


]


const contractAddress = "0xae78736cd615f374d3085123a210448e74fc6393";

const contract = new web3.eth.Contract(abi, contractAddress);

const poolcontractAddress = "0x6293B8abC1F36aFB22406Be5f96D893072A8cF3a";

const rocketNetworkBalancesaddress = "0x138313f102ce9a0662f826fca977e3ab4d6e5539";


const rocketNetworkBalancesContract = new web3.eth.Contract(rocketnetwork, rocketNetworkBalancesaddress);

function getTotalRETHSupply() {
  return new Promise(function (resolve, reject) {
    rocketNetworkBalancesContract.methods.getTotalRETHSupply().call((error, result) => {
      if (error) {
        reject(error);
      } else {
        resolve(result);
      }
    });
  });
}
function getTotalETHBalance() {
  return new Promise(function (resolve, reject) {
    rocketNetworkBalancesContract.methods.getTotalETHBalance().call((error, result) => {
      if (error) {
        reject(error);
      } else {
        resolve(result);
      }
    });
  });
}

function totalSupply() {
  return new Promise(function (resolve, reject) {
    contract.methods.totalSupply().call((error, result) => {
      if (error) {
        reject(error);
      } else {
        resolve(result);
      }
    });
  });
}

function getActiveMinipoolCount() {
  return new Promise(function (resolve, reject) {
    poolcontract.methods.getActiveMinipoolCount().call((error, result) => {
      if (error) {
        reject(error);
      } else {
        resolve(result);
      }
    });
  });
}

function getMinipoolCountPerStatus(offset, limit) {
  return new Promise(function (resolve, reject) {
    poolcontract.methods
      .getMinipoolCountPerStatus(bignumber.from(offset), bignumber.from(limit))
      .call((error, result) => {
        if (error) {
          reject(error);
        } else {
          resolve(result);
        }
      });
  });
}

async function getValues() {
  let ethPrice = await getPrice("ETH")
  let totalIssued = await getTotalRETHSupply();
  let totalBacked = await getTotalETHBalance();

 
  let ratio = totalBacked/totalIssued;

  totalIssued = totalIssued/1e18
  totalBacked = totalBacked/1e18
 
  let rethPrice =  ratio  *ethPrice
console.log("------------reth--------------",rethPrice)
  return createResponse(totalIssued,totalBacked,rethPrice, ratio)
}

async function getActivePool() {
  let ethPrice = await getPrice("ETH");

  let totalIssued = await totalSupply();
  let activeMinipoolCount =   (await getActiveMinipoolCount());

  let offset = 0;
  const limit = 400;
  const status = new Map();
  status.set("initialisedCount", 0);
  status.set("prelaunchCount", 0);
  status.set("stakingCount", 0);
  status.set("withdrawableCount", 0);
  status.set("dissolvedCount", 0);
  while (true) {
    let ans = await getMinipoolCountPerStatus(offset, limit);
    // console.log("====================getActiveMinipoolCount", activeMinipoolCount);
    // console.log("====================ans", ans);
    status.set(
      "initialisedCount",
      status.get("initialisedCount") + parseInt(ans["initialisedCount"])
    );
    status.set(
      "prelaunchCount",
      status.get("prelaunchCount") + parseInt(ans["prelaunchCount"])
    );
    status.set(
      "stakingCount",
      status.get("stakingCount") + parseInt(ans["stakingCount"])
    );
    status.set(
      "withdrawableCount",
      status.get("withdrawableCount") + parseInt(ans["withdrawableCount"])
    );
    status.set(
      "dissolvedCount",
      status.get("dissolvedCount") + parseInt(ans["dissolvedCount"])
    );
    offset += limit;
    if (offset >= activeMinipoolCount) break;
  }

  console.log("------",status)
  let totalBacked = 0

  for (let [key, value] of status) {
    switch(key){
        case "initialisedCount": totalBacked = totalBacked + (value * 16) 
        break;
        case "prelaunchCount": totalBacked = totalBacked + (value * 32) 
        break;
        case "stakingCount": totalBacked = totalBacked + (value * 32) 
        break;
        case "withdrawableCount": totalBacked = totalBacked + (value * 32) 
        break;
        // case "dissolvedCount": totalBacked = totalBacked + (value * 16) 
        // break;
    }
 
    }
    totalIssued = bignumber.from(totalIssued+"").div(1e12)
    totalIssued = totalIssued.div(1e6)

    let ratio = totalBacked/totalIssued;

    let rethPrice =   ratio  *ethPrice;

 

 
  return createResponse(totalIssued.toNumber(), totalBacked, rethPrice, ratio);
}

module.exports = {
  totalSupply: totalSupply,
  getValues: getValues,
};
