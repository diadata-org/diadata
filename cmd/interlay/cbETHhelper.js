const {collaterlaratio} = require("./utils")
const Web3 = require("web3");
const web3 = new Web3(
  new Web3.providers.HttpProvider(
    "https://mainnet.infura.io/v3/2883d1b22e0e4d62b535592dd8075fee"
  )
);

let abi = [
  {
    inputs: [],
    name: "exchangeRate",
    outputs: [
      {
        internalType: "uint256",
        name: "_exchangeRate",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
  {
    inputs: [],
    name: "totalSupply",
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
];

const contractAddress = "0xBe9895146f7AF43049ca1c1AE358B0541Ea49704";

const contract = new web3.eth.Contract(abi, contractAddress);

function exchangeRate() {
  return new Promise(function (resolve, reject) {
    contract.methods.exchangeRate().call((error, result) => {
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

async function getValues() {
  let totalIssued = await totalSupply();
  // let totalBacked = await getTotalPooledETH();
  return collaterlaratio(totalIssued, "");
}

module.exports = {
  exchangeRate: exchangeRate,
  totalSupply: totalSupply,
  getValues: getValues,
};
