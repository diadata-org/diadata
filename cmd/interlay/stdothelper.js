const Web3 = require("web3");
const web3 = new Web3(
  new Web3.providers.HttpProvider("https://rpc.api.moonbeam.network/")
);

let abi = [
  {
    "inputs": [],
    "name": "fundRaisedBalance",
    "outputs": [{
      "internalType": "uint256",
      "name": "",
      "type": "uint256"
    }],
    "stateMutability": "view",
    "type": "function"
  },{
    "inputs": [],
    "name": "totalSupply",
    "outputs": [{
      "internalType": "uint256",
      "name": "",
      "type": "uint256"
    }],
    "stateMutability": "view",
    "type": "function"
  },
];

const { createResponse, getPrice } = require("./utils");

const contractAddress = "0xfa36fe1da08c89ec72ea1f0143a35bfd5daea108";

const contract = new web3.eth.Contract(abi, contractAddress);

function fundRaisedBalance() {
  return new Promise(function (resolve, reject) {
    contract.methods.fundRaisedBalance().call((error, result) => {
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
  let dotPrice = await getPrice("DOT");
  let totalIssued = await  totalSupply();
  let totalBacked = await  fundRaisedBalance();

  let ratio = totalBacked / totalIssued;

  let stdotPrice =
    ratio > 1 ? dotPrice : (dotPrice * totalBacked) / totalIssued;

  totalIssued = totalIssued / 1e10;
  totalBacked = totalBacked / 1e10;

  return createResponse(totalIssued, totalBacked, stdotPrice, ratio);
}

function getStdot() {
  return new Promise(function (resolve, reject) {
    contract.methods.stDOTPrice().call((error, result) => {
      if (error) {
        reject(error);
      } else {
        resolve(result);
      }
    });
  });
}

module.exports = {
  getStdot: getStdot,
  getValues: getValues,
};
