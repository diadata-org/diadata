const Web3 = require("web3");

const {createResponse,getPrice} = require("./utils")

const web3 = new Web3(
  new Web3.providers.HttpProvider(
    "https://mainnet.infura.io/v3/2883d1b22e0e4d62b535592dd8075fee"
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

const contractAddress = "0xae7ab96520de3a18e5e111b5eaab095312d7fe84";

const contract = new web3.eth.Contract(abi, contractAddress);

function getTotalPooledETH() {
  return new Promise(function (resolve, reject) {
    contract.methods.getTotalPooledEther().call((error, result) => {
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
  let ethPrice = await getPrice("ETH")
  let totalIssued = await totalSupply();
  let totalBacked = await getTotalPooledETH();

 
  let ratio = totalBacked/totalIssued;

  totalIssued = totalIssued/1e18
  totalBacked = totalBacked/1e18

  let stETHprice =  ratio > 1
  ? ethPrice
  : ethPrice *  totalBacked/totalIssued

  return createResponse(totalIssued,totalBacked,stETHprice, ratio)
}




module.exports = {
  totalSupply: totalSupply,
  getTotalPooledETH: getTotalPooledETH,
  getValues: getValues,
};
