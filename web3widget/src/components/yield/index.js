import React, { Component} from 'react';

// react
import { Container, Row, Col, Button, Form, Alert, Spinner } from 'react-bootstrap';
import Slider from 'react-rangeslider';
import { SwappingSquaresSpinner } from 'react-epic-spinners';

// web3
import { config } from "../../helpers/web3/config.js";
import { getYieldDetails, getDiaBalance, getDiaAllowance, approveDiaForStaking, stakeDia  } from "../../helpers/web3/dia.js";

// css
import './index.css';

export default class YieldCalculator extends Component {

    constructor(props, context) {
        super(props, context)

        this.state = {
            yieldDuration: 3,
            yieldPercentage: 15,
            yieldRates:{},
            yieldResults: [],
            userStake: 0,
            web3Connected: false,
            fetchingRates: false,
            web3: undefined,
            diaRate: undefined,
            currentNetwork: "",
            diaNotApprovedForStaking: true,
            networkConfig: {},
            userAccount: "",
            userDiaBalance: 0,
            userDiaAllowance: 0,
            errorMsg: "",
            processingTx: false,
            viewTx: false,
            txHash: ""
        }

        // bind methods
        this.handleOnRangeChange  = this.handleOnRangeChange.bind(this);
        this.renderCalculator = this.renderCalculator.bind(this);
        this.renderStakingForm = this.renderStakingForm.bind(this);
        this.connectToWeb3 = this.connectToWeb3.bind(this);
        this.fetchYieldRates = this.fetchYieldRates.bind(this);
        this.renderFetchingRates = this.renderFetchingRates.bind(this);
        this.handleOnStakeChange = this.handleOnStakeChange.bind(this);
        this.handleOnAccountChange = this.handleOnAccountChange.bind(this);
        this.renderStakingFormBody = this.renderStakingFormBody.bind(this);
        this.renderErrorMsg = this.renderErrorMsg.bind(this);
        this.approveDia = this.approveDia.bind(this);
        this.getTXLink = this.getTXLink.bind(this);
        this.resetViewTx = this.resetViewTx.bind(this);
        this.stakeUserDia = this.stakeUserDia.bind(this);
    
    }

    async componentDidMount(){
       this.setState({ networkConfig: this.props.networkConfig });
    }

    resetViewTx() {
        // remove the view tx view
        const vm = this;
        const id = setTimeout(function() {
            if (vm.state.viewTx) {
              vm.setState({ viewTx: false, txHash: "" });

              clearTimeout(id);
            }
          }, 6000);
    }

    async approveDia() {
        try {
            this.setState({ processingTx: true });

            // web3
            const web3 = this.state.web3;

            // get the default network
            const defaultNetwork = this.state.networkConfig.defaultNetwork;
        
            // get the yield rates
            const yieldContractAddress = this.state.networkConfig[defaultNetwork].yieldContractAddress;
        
            // get the user DIA balance
            const erc20contractAddress = this.state.networkConfig[defaultNetwork].erc20ContractAddress;
            const erc20contractAbi = JSON.parse(config.erc20Contract.abi);

            // approve the DIA for staking
            const diaApproved = await approveDiaForStaking(web3, erc20contractAbi, erc20contractAddress, yieldContractAddress, String(this.state.userStake));

            const txHash = diaApproved.transactionHash;

            console.log(txHash);

            // update the user allowance to the allowed amount
            this.setState({ userDiaAllowance: Number(this.state.userStake), 
                            diaNotApprovedForStaking: false, 
                            processingTx: false,
                            viewTx: true,
                            txHash
                        });

            this.resetViewTx();

        }
        catch(error) {

            this.setState({  diaNotApprovedForStaking: true, processingTx: false });
            console.log(error);
        }
    }

    async stakeUserDia() {
        try {

            this.setState({ processingTx: true });

            // web3
            const { web3, yieldDuration, userAccount, userStake, userDiaAllowance } = this.state;

            // get the default network
            const defaultNetwork = this.state.networkConfig.defaultNetwork;
        
            // get the yield rates
            const yieldContractAddress = this.state.networkConfig[defaultNetwork].yieldContractAddress;
            const yieldContractAbi = JSON.parse(config.yieldContract.abi);
        
           
            // stake the dia
            const diaStaked= await stakeDia(String(yieldDuration)+"m", web3, yieldContractAbi, yieldContractAddress, userAccount, String(userStake));

            const txHash = diaStaked.transactionHash;

            console.log(txHash);

            // update the user allowance to the allowed amount
            this.setState({ userDiaAllowance: Number(userDiaAllowance) - Number(userStake), 
                            processingTx: false,
                            viewTx: true,
                            txHash
                        });

            this.resetViewTx();

        }
        catch(error) {

            this.setState({  processingTx: false, viewTx: false });
            console.log(error);
        }
    }

    handleOnAccountChange(event) {
        const vm = this;
        const userAccount = event.target.value;
        vm.setState({userAccount});
    }

    handleOnStakeChange(event) {
        const vm = this;
        const userStake = event.target.value;
        let errorMsg = "";

        let diaNotApprovedForStaking = true;

        if (Number(userStake) > Number(this.state.userDiaBalance)) {
            errorMsg = `Staking amount, is greater than your available DIA balance of ${Number(this.state.userDiaBalance).toLocaleString()}!`;
        }

        if ( Number(userStake) > 0 && (Number(userStake) <= Number(this.state.userDiaAllowance))){
            diaNotApprovedForStaking = false;
        }

        vm.setState({ userStake, errorMsg, diaNotApprovedForStaking }, ()=> {
            vm.calculateYieldResults()
        });
    }

    handleOnRangeChange (yieldDuration) {
        const vm = this;
        vm.setState({yieldDuration, yieldPercentage: this.state.yieldRates[yieldDuration]})
    }

    async connectToWeb3() {

        try {
            const web3 = await this.props.getWeb3();

            // set the current network connected to
            const currentNetwork = await web3.eth.net.getNetworkType();
          
            // set the state and also fetch the yield rates from api & contract
            const vm = this;
            const userAccount = (await web3.eth.getAccounts())[0];

            vm.setState({ web3: web3, web3Connected: true, currentNetwork, userAccount }, ()=> {
                vm.setState({fetchingRates: true });
                vm.fetchYieldRates();
            })
        }
        catch(error) {
            console.log(error);
        }
    }

    async fetchYieldRates() {
        try {
            
            const diaRate = await this.props.getDiaRate();

            // get the default network
            const defaultNetwork = this.state.networkConfig.defaultNetwork;
           
            const web3 = this.state.web3;

            // get the yield rates
            const yieldContractAddress = this.state.networkConfig[defaultNetwork].yieldContractAddress;
            const yieldContractAbi = JSON.parse(config.yieldContract.abi);
            const yieldResults = await getYieldDetails(web3, yieldContractAbi, yieldContractAddress);

            // get the user DIA balance
            const erc20contractAddress = this.state.networkConfig[defaultNetwork].erc20ContractAddress;
            const erc20contractAbi = JSON.parse(config.erc20Contract.abi);
            const userDiaBalance = await getDiaBalance(web3, erc20contractAbi, erc20contractAddress);
            const userDiaAllowance = await getDiaAllowance(web3, erc20contractAbi, erc20contractAddress, yieldContractAddress);

            console.log("user DIA balance is ", userDiaBalance.toLocaleString() );
            console.log("user DIA balance approved for staking is ", userDiaAllowance.toLocaleString() );


            // set the yield & api rates
            const yieldRates = {3: Number(yieldResults.threeMonthPromille) / 10, 
                                6: Number(yieldResults.sixMonthPromille ) / 10, 
                                9: Number(yieldResults.nineMonthPromille) / 10 }

            this.setState({ fetchingRates: false, 
                            yieldRates: yieldRates, 
                            diaRate:diaRate, 
                            yieldPercentage: yieldRates[this.state.yieldDuration], 
                            userDiaBalance,
                            userDiaAllowance  });

        }
        catch(error){
            this.setState({fetchingRates: false});
            console.log(error);
        }
    }

    async calculateYieldResults() {

        const { userStake, yieldRates, diaRate } = this.state;

        if( userStake > 0) {

            let yieldResults = [];

            for (let i = 0; i < 3; i++) {

                let period = ""
                let diaAmount = 0;

                if(i === 0){
                   diaAmount = Math.round(yieldRates[3] / 100 * userStake);
                   period = "3m"
                }

                if(i === 1){
                    diaAmount = Math.round(yieldRates[6] / 100 * userStake);
                    period = "6m"
                }

                if(i === 2){
                    diaAmount = Math.round(yieldRates[9] / 100 * userStake);
                    period = "9m"
                }

                const diaAmountFormatted = diaAmount.toLocaleString() + "  " + diaRate.Symbol;
                const diaUsdAmount = Math.round(diaAmount * diaRate.Price);
                const diaUsdAmountFormatted = "$" + diaUsdAmount.toLocaleString();

                yieldResults.push({id: (i+1), period,  diaAmount: diaAmountFormatted, usdAmount: diaUsdAmountFormatted});
               
            }
            this.setState({ yieldResults });

        } else{
            this.setState({yieldResults: []});
        }
    }

    renderFetchingRates() {

        return (
            <Container fluid className="yield-form-container">
                <hr id="top-line"/>

                <Row>
                    <Col xs={12} className="loading-rates"> 
                         Fetching Rates..
                    </Col>
                    <Col xs={12} className="loading-rates"> 
                      <SwappingSquaresSpinner color='#fff' size={200} />
                    </Col>
                </Row>
            </Container>
        );

    }

    renderErrorMsg = (msg) =>  <Alert variant='danger' className="error-msg">{msg}</Alert>

    renderStakingFormBody() {

        const { yieldDuration, yieldPercentage, yieldResults, userStake, diaNotApprovedForStaking, userAccount, errorMsg, processingTx, viewTx  } = this.state;

        return (
            <Container fluid className="yield-form-container">
                <hr id="top-line"/>
                { errorMsg !== "" ? this.renderErrorMsg(errorMsg): null }
               
                <Row>
                    <Col id="user-stake-heading" className="headings">Your Dia Stake</Col>
                    <Col id="user-returns-heading" className="headings">Return</Col>
                </Row>
                <Row>
                    <Col>
                         <Form.Control 
                            type="number" 
                            placeholder={userStake}  
                            id="user-stake" 
                            className="returns" 
                            value={userStake} 
                            onChange={this.handleOnStakeChange}
                            />
                    </Col>
                     <Col id="user-returns" className="returns">+{yieldPercentage}%</Col>
                </Row>
                <Row>
                    <Col xs={12}>
                        <Slider
                            value={yieldDuration}
                            onChange={this.handleOnRangeChange}
                            min={3}
                            max={9}
                            step={3}
                            tooltip={false}
                        />
                        <div className="pointer-lines-ranges">
                            <div className="pointer-lines" id="first-line-margin"></div>
                            <div className="pointer-lines" id="second-line-margin"></div>
                            <div className="pointer-lines" id="third-line-margin"></div>
                        </div>
                        <div className="yield-ranges">
                            <span>3 months</span>
                            <span>6 months</span>
                            <span>9 months</span>
                        </div>
                    </Col>
                </Row>
                
                { /* The staking data */ }
                <div id="yield-results">

                    {yieldResults.map(res => {
                        return (
                            <Row key={res.id} className="row-res row-res-border">
                                <Col xs={4} className="yield-period">{res.period}</Col>
                                <Col xs={4} className="yield-amount yield-dia">{res.diaAmount}</Col>
                                <Col xs={4} className="yield-amount">{res.usdAmount}</Col>
                            </Row>
                        )
                    })}
                    
                </div>
                <Row className="footer">
                    <Form.Label column sm="12" className="form-lbl">
                        Receiving Address
                    </Form.Label>
                    <Col sm="12">
                        <Form.Control 
                            type="text" 
                            placeholder="i.e 0x69826De274580f49ef8e6e89f84b8eAb6CC9C20B" 
                            value={userAccount} 
                            onChange={this.handleOnAccountChange}/>
                    </Col>
                    { viewTx ? <Col sm="12" className="view-tx">{this.getTXLink()}</Col> : null }
                    { diaNotApprovedForStaking ? 
                        processingTx ? <Spinner animation="border" className="processing-tx" /> :
                        <Button id="staking-btn" disabled={userStake <= 0 || errorMsg !== ""} onClick={this.approveDia}>PREPARE STAKING</Button> : 
                        processingTx ? <Spinner animation="border" className="processing-tx" /> :
                        <Button id="staking-btn" disabled={userStake <= 0 || errorMsg !== ""} onClick={this.stakeUserDia}>START STAKING</Button> }

                
                </Row>
            </Container>
        )

    }

    getTXLink() {
        const { txHash, currentNetwork } = this.state;
        const etherscanUrl = `https://${currentNetwork}.etherscan.io/tx/${txHash}`;
        return <a href={etherscanUrl} target="_blank" rel="noopener noreferrer">View Tx On Etherscan</a>
    }

    renderStakingForm() {
        const { web3Connected } = this.state;
        return web3Connected ? this.renderStakingFormBody() : null 
    }

    renderCalculator() {
        const { web3Connected, currentNetwork, fetchingRates } = this.state;

        return(
            <div className="yield-calc">
                <header className="header">
                    Calculate Yield
                    {
                        web3Connected ? <Button variant="outline-success" id="connected-btn" disabled>Connected to {currentNetwork}</Button>
                        :  <Button id="web3-connect" onClick={()=> this.connectToWeb3()}>Connect To Metamask</Button>
                    }
                </header>
                
                
                { fetchingRates ? this.renderFetchingRates() : this.renderStakingForm()  }
            </div>
        )

    }
    
    render = () => this.renderCalculator() 
    

}