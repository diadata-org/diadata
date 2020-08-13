import React, { Component} from 'react';

import { Modal, Button, } from 'react-bootstrap';

// wallet connect
import { createConnector } from '../walletconnect';

// web3
import getWeb3 from "../../helpers/web3";

export default class WalletSelector extends Component {

    constructor(props, context) {
        super(props, context);

        this.state = {
            show: true
        };

        this.connectToWeb3 = this.connectToWeb3.bind(this);
    }

    async connectToWeb3() {

        let web3 = undefined;

        try {
            web3 = await getWeb3();
        }
        catch(error) {
            console.log(error);
        }

        return web3;
    }

    async connectToWallectConnect() {
        return createConnector();
    }


    render() {
        return (
            <Modal
              show={this.state.show}
              onHide={() => this.setState({show : false})}
              backdrop="static"
              keyboard={false}
            >
              <Modal.Header closeButton>
                <Modal.Title>Please select wallet</Modal.Title>
              </Modal.Header>
              <Modal.Body>
                <Button variant="primary" onClick={() => this.connectToWeb3() }>
                  Metamask
                </Button>
                <Button variant="primary" onClick={() => this.connectToWallectConnect()}>
                  Wallet Connect
                </Button>
              </Modal.Body>
              <Modal.Footer>
                <Button variant="secondary" onClick={() => this.setState({show : false})}>
                  Close
                </Button>
              </Modal.Footer>
            </Modal>
        );

    }

}