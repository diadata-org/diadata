'use strict';

module.exports = {
  networks: {
    development: {
      host: 'localhost',
      port: 8545,
      gas: 5000000,
      gasPrice: 5e9,
      network_id: '*'
    }
  }
};
