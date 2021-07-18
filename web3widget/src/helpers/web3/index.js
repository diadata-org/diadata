import Web3 from "web3";

const getWeb3 = () => {
  return new Promise(async (resolve, reject) => {
      // Modern dapp browsers...
      if (window.ethereum) {
        const web3 = new Web3(window.ethereum);
        try {
          // Request account access if needed
          await window.ethereum.enable();
          // Acccounts now exposed
          resolve(web3);
        } catch (error) {
          reject(error);
        }
      }
      // Legacy dapp browsers...
      if (window.web3) {
        // Use Mist/MetaMask's provider.
        const web3 = window.web3
        console.log("Injected web3 detected.")
        // Acccounts now exposed
        resolve(web3);
      }
      
  });

}
  
export default getWeb3