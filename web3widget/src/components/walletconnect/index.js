import Web3 from "web3";
import WalletConnectProvider from "@walletconnect/web3-provider";

//  Create WalletConnect Provider
const provider = new WalletConnectProvider({
  infuraId: "be1a3f5f45994142bb67759b9fef28c5" // Required
});


export const createConnector = async (handleWeb3, handleAccountsChange, handleWalletClose) => {

    //  Enable session (triggers QR Code modal)
    await provider.enable();

    //  Create Web3
    const web3 = new Web3(provider);



    // Subscribe to accounts change
    provider.on("accountsChanged", (accounts) => {
        console.log(accounts);

        handleAccountsChange(accounts);
    });
    
    // Subscribe to chainId change
    provider.on("chainChanged", (chainId) => {

        console.log(chainId);
        window.location.reload();
    });

    // Subscribe to networkId change
    provider.on("networkChanged", (networkId) => {
        console.log(networkId);
        window.location.reload();
    });

    // Subscribe to session connection/open
    provider.on("open", () => {
        console.log("open");

        handleWeb3(web3);
    });

    // Subscribe to session disconnection/close
    provider.on("close", (code, reason) => {
        console.log(code, reason);
        handleWalletClose(true);
    });

}

