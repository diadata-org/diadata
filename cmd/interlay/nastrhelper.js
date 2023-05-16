const Web3 = require("web3");
const web3 = new Web3(
  new Web3.providers.HttpProvider("https://evm.astar.network/")
);
const {createResponse,getPrice} = require("./utils")


let abi = [
  {
    type: "function",
    stateMutability: "view",
    outputs: [
      {
        type: "uint256",
        name: "",
        internalType: "uint256",
      },
    ],
    name: "getAstrInStakingContract",
    inputs: [],
  },
  {
    type: "function",
    stateMutability: "view",
    outputs: [
      {
        type: "uint256",
        name: "",
        internalType: "uint256",
      },
      {
        type: "uint128",
        name: "",
        internalType: "uint128",
      },
    ],
    name: "getNAstrPriceFromCollateralRatio",
    inputs: [],
  },
  {
    type: "function",
    stateMutability: "view",
    outputs: [
      {
        type: "uint256",
        name: "",
        internalType: "uint256",
      },
    ],
    name: "getNAstrTotalSupply",
    inputs: [],
  },
  {
    type: "function",
    stateMutability: "view",
    outputs: [
      {
        type: "uint256",
        name: "",
        internalType: "uint256",
      },
    ],
    name: "values",
    inputs: [
      {
        type: "string",
        name: "",
        internalType: "string",
      },
    ],
  },
];

const contractAddress = "0xCb2158f97367596A68dd1670bC848dc6B2245111";

const contract = new web3.eth.Contract(abi, contractAddress);

function getAstrInStakingContract() {
  return new Promise(function (resolve, reject) {
    contract.methods.getAstrInStakingContract().call((error, result) => {
      if (error) {
        reject(error);
      } else {
        resolve(result);
      }
    });
  });
}

function getNAstrPriceFromCollateralRatio() {
  return new Promise(function (resolve, reject) {
    contract.methods
      .getNAstrPriceFromCollateralRatio()
      .call((error, result) => {
        if (error) {
          reject(error);
        } else {
          resolve(result);
        }
      });
  });
}

function getNAstrTotalSupply() {
  return new Promise(function (resolve, reject) {
    contract.methods.getNAstrTotalSupply().call((error, result) => {
      if (error) {
        reject(error);
      } else {
        resolve(result);
      }
    });
  });
}

async function getValues(){
  let astrPrice = await getPrice("ASTR")

   let totalIssued = await getNAstrTotalSupply();
  let totalBacked = await getAstrInStakingContract();

 
  let ratio = totalBacked/totalIssued;

  totalIssued = totalIssued/1e18
  totalBacked = totalBacked/1e18

  let nastrPrice =  ratio > 1
  ? astrPrice
  : astrPrice *  totalBacked/totalIssued



  return createResponse(totalIssued,totalBacked,nastrPrice, ratio)


}

module.exports = {
  getNAstrTotalSupply: getNAstrTotalSupply,
  getNAstrPriceFromCollateralRatio: getNAstrPriceFromCollateralRatio,
  getAstrInStakingContract: getAstrInStakingContract,
  getValues:getValues,
};
