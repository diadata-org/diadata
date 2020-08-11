import React, { Component} from 'react';

import { Container, Row, Col, Button, Form } from 'react-bootstrap';
import Slider from 'react-rangeslider';

import './index.css';

export default class YieldCalculator extends Component {

    constructor(props, context) {
        super(props, context)

        this.state = {
            yieldDuration: 3,
            yieldResults: [{
                id:1,
                period: "3m",
                diaAmount: "125,015 DIA",
                usdAmount: "$263,781"
            },
            {
                id:2,
                period: "6m",
                diaAmount: "150,018 DIA",
                usdAmount: "$316,538"
            },
            {
                id:3,
                period: "9m",
                diaAmount: "200,025 DIA",
                usdAmount: "$422,053"
            }],
            userStake: 0,
            web3Connected: true,
            fetchedRate: false,
            web3: undefined,
            diaRate: undefined

        }

        // bind methods
        this.handleOnChange = this.handleOnChange.bind(this);
        this.renderCalculator = this.renderCalculator.bind(this);
        this.renderStakingForm = this.renderStakingForm.bind(this);
        this.renderApprovalForm = this.renderApprovalForm.bind(this);
        this.connectToWeb3 = this.connectToWeb3.bind(this);
        this.fetchYieldRates = this.fetchYieldRates.bind(this);
    }

    async componentDidMount(){
       console.log(this.props)

    }

    handleOnChange = (yieldDuration) => {
        this.setState({yieldDuration})
    }

    connectToWeb3() {

        try {
            const web3 = this.props.getWeb3();
            // set the state and also fetch the yield rates from api & contract
            const vm = this;
            vm.setState({ web3: web3, web3Connected: true }, ()=> {
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

                console.log(diaRate);
        }
        catch(error){
            console.log(error);
        }
    }

    stakeDia() {

    }

    renderApprovalForm() {

        const { web3Connected, userStake } = this.state;

        return (
            web3Connected && userStake <= 0 ? 
            <Form className="yield-form-container">
                <Form.Group as={Row} controlId="formRecevingAddress">
                    <Form.Label column sm="12" className="form-lbl">
                        Receiving Address
                    </Form.Label>
                    <Col sm="12">
                        <Form.Control type="text" placeholder="i.e 0x69826De274580f49ef8e6e89f84b8eAb6CC9C20B" />
                    </Col>
                </Form.Group>
                <Form.Group as={Row} controlId="formStakingAmount">
                    <Form.Label column sm="8" className="form-lbl">
                        Staking Amount
                    </Form.Label>
                    <Col sm="12">
                        <Form.Control type="number" placeholder="1000" />
                    </Col>
                </Form.Group>

                <Form.Group as={Row} controlId="formApprove">
                    <Col sm="12" className="footer">
                        <Button id="approve-btn">Continue</Button>
                    </Col>
                </Form.Group>
                     
            </Form>
            :
            null
        )
    }

    renderStakingForm() {

        const { yieldDuration, yieldResults, userStake } = this.state;

        return (
            <Container fluid className="yield-form-container">
                <hr id="top-line"/>
                <Row>
                    <Col id="user-stake-heading" className="headings">Your Dia Stake</Col>
                    <Col id="user-returns-heading" className="headings">Return</Col>
                </Row>
                <Row>
                    <Col id="user-stake" className="returns">{ userStake }</Col>
                    <Col id="user-returns" className="returns">+15%</Col>
                </Row>
                <Row>
                    <Col xs={12}>
                        <Slider
                            value={yieldDuration}
                            onChange={this.handleOnChange}
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
                    <Button id="staking-btn">START STAKING</Button>
                </Row>
            </Container>
        )

    }
    renderCalculator() {
        const { userStake, web3Connected, fetchedRate } = this.state;

        return(
            <div className="yield-calc">
                <header className="header">
                    Yield Calculator
                    {
                        web3Connected ? <Button variant="outline-success" id="connected-btn" disabled>Connected</Button>
                        :  <Button id="web3-connect">Connect To Metamask</Button>
                    }
                </header>
                
                
                { userStake > 0 && web3Connected && fetchedRate  ? this.renderStakingForm() : this.renderApprovalForm() }
            </div>
        )

    }
    
    render = () => this.renderCalculator() 
    

}