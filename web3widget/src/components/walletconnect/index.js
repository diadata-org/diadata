import WalletConnect from "@walletconnect/client";
import QRCodeModal from "@walletconnect/qrcode-modal";


// web3
import { config } from "../../helpers/web3/config.js";
import { getYieldDetails, getDiaBalance, getDiaAllowance, approveDiaForStaking, stakeDia  } from "../../helpers/web3/dia.js";


export const createConnector = () => {

    // Create a connector
    const connector = new WalletConnect({
        bridge: "https://bridge.walletconnect.org", // Required
        qrcodeModal: QRCodeModal,
    });
  
    // Check if connection is already established
    if (!connector.connected) {
        // create new session
        connector.createSession();
    }
  
    // Subscribe to connection events
    connector.on("connect", (error, payload) => {
        if (error) {
            throw error;
        }
    
        // Get provided accounts and chainId
        const { accounts, chainId } = payload.params[0];

        const userAccount = accounts[0];
    });
  
    connector.on("session_update", (error, payload) => {
        if (error) {
            throw error;
        }
  
        // Get updated accounts and chainId
        const { accounts, chainId } = payload.params[0];

        const userAccount = accounts[0];
    });
  
    connector.on("disconnect", (error, payload) => {
        if (error) {
            throw error;
        }
    
        // Delete connector
    });
  
};
