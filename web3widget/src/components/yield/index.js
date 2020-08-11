import React, { Component} from 'react';

// react
import { Container, Row, Col, Button, Form } from 'react-bootstrap';
import Slider from 'react-rangeslider';
import { SwappingSquaresSpinner } from 'react-epic-spinners';

// web3
import { config } from "../../helpers/web3/config.js";
import { getYieldDetails } from "../../helpers/web3/dia.js";


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
            approveDiaForStaking: true,
            networkConfig: {},
            userccount: ""
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
    }

    async componentDidMount(){
       this.setState({ networkConfig: this.props.networkConfig });
    }

    handleOnAccountChange(event) {
        const vm = this;
        const userAccount = event.target.value;
        vm.setState({userAccount});
    }

    handleOnStakeChange(event) {
        const vm = this;
        const userStake = event.target.value;
        vm.setState({userStake}, ()=> {
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
            console.log(userAccount);

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

            // get the yield details from the contract
            const defaultNetwork = this.state.networkConfig.defaultNetwork;
            const contractAddress = this.state.networkConfig[defaultNetwork].yieldContractAddress;
            const contractAbi = JSON.parse(config.yieldContract.abi);
            const web3 = this.state.web3;
            const yieldResults = await getYieldDetails(web3,contractAbi,contractAddress);


            // set the yield & api rates

            const yieldRates = {3: Number(yieldResults.threeMonthPromille) / 10, 
                                6: Number(yieldResults.sixMonthPromille ) / 10, 
                                9: Number(yieldResults.nineMonthPromille) / 10 }

            this.setState({fetchingRates: false, yieldRates: yieldRates, diaRate:diaRate, yieldPercentage: yieldRates[this.state.yieldDuration]  });

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

                if(i === 0){
                    const diaAmount = Math.round(yieldRates[3] / 100 * userStake);
                    const diaAmountFormatted = diaAmount.toLocaleString() + "  " + diaRate.Symbol;
                 

                    const diaUsdAmount = Math.round(diaAmount * diaRate.Price);
                    const diaUsdAmountFormatted = "$ " + diaUsdAmount.toLocaleString();

                    yieldResults.push({id: (i+1), period: "3m",  diaAmount: diaAmountFormatted, usdAmount: diaUsdAmountFormatted});
                }

                if(i === 1){
                    const diaAmount = Math.round(yieldRates[6] / 100 * userStake);
                    const diaAmountFormatted = diaAmount.toLocaleString() + "  " + diaRate.Symbol;
                 

                    const diaUsdAmount = Math.round(diaAmount * diaRate.Price);
                    const diaUsdAmountFormatted = "$ " + diaUsdAmount.toLocaleString();

                    yieldResults.push({id: (i+1), period: "6m",  diaAmount: diaAmountFormatted, usdAmount: diaUsdAmountFormatted});
                }

                if(i === 2){
                    const diaAmount = Math.round(yieldRates[9] / 100 * userStake);
                    const diaAmountFormatted = diaAmount.toLocaleString() + "  " + diaRate.Symbol;
                 

                    const diaUsdAmount = Math.round(diaAmount * diaRate.Price);
                    const diaUsdAmountFormatted = "$ " + diaUsdAmount.toLocaleString();

                    yieldResults.push({id: (i+1), period: "9m",  diaAmount: diaAmountFormatted, usdAmount: diaUsdAmountFormatted});
                }
               
            }
            this.setState({yieldResults});

        } else{
            this.setState({yieldResults: []});
        }
    }

    stakeDia() {

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
                      <SwappingSquaresSpinner color='#133aab' size={200} />
                    </Col>
                </Row>
            </Container>
        )

    }

    renderStakingFormBody() {

        const { yieldDuration, yieldPercentage, yieldResults, userStake, approveDiaForStaking, userAccount } = this.state;

        return (
            <Container fluid className="yield-form-container">
                <hr id="top-line"/>
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
                   { approveDiaForStaking ? <Button id="staking-btn" disabled={userStake <= 0}>PREPARE STAKING</Button> : <Button id="staking-btn" disabled={userStake <= 0}>START STAKING</Button>}
                </Row>
            </Container>
        )

    }

    renderStakingForm() {

        const {  web3Connected } = this.state;

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