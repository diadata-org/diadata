import WalletConnect from "@walletconnect/client";
import QRCodeModal from "@walletconnect/qrcode-modal";

export const createConnector = (handleWc) => {
    
    // Create a connector
    const connector = new WalletConnect({
        bridge: "https://bridge.walletconnect.org", // Required
        qrcodeModal: QRCodeModal,
    });
  
    // Check if connection is already established
    if (!connector.connected) {
        // create new session
        connector.createSession();
    } else {
        // kill any existing sessions
        connector.killSession();

        // and create a new one
        connector.createSession();
    }
  
    // Subscribe to connection events
    connector.on("connect", (error, payload) => {
        if (error) {
            console.log(console.error);
            throw error;
        }

        console.log(payload);

        handleWc({ web3Connected: true });
    
        // Get provided accounts and chainId
        const { accounts, chainId } = payload.params[0];

        const userAccount = accounts[0];
    });
  
    connector.on("session_update", (error, payload) => {
        if (error) {
            console.log(console.error);
            throw error;
        }


        console.log(payload);
  
        // Get updated accounts and chainId
        const { accounts, chainId } = payload.params[0];

        const userAccount = accounts[0];
    });
  
    connector.on("disconnect", (error, payload) => {
        if (error) {
            console.log(console.error);
            throw error;
        }


        console.log(payload);
    
        // Delete connector
    });
  
};
