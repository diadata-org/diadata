'use strict';

module.exports = {
  networks: {
    development: {
      host: 'localhost',
      port: 8545,
      gas: 5000000,
      gasPrice: 5e9,
      network_id: '*'
    },
    rinkeby: {
      host: "localhost",
      port: 8545,
      network_id: "4", // Rinkeby ID 4
      // from: account from which to deploy should be specified in deploymentScript
      gas: 6712390
    }
  },
  compilers: {
    solc: {
      version: '0.4.25'
    }
  },
}
