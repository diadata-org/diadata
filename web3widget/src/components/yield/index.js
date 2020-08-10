import React, { Component} from 'react';

import { Container, Row, Col } from 'react-bootstrap';
import Slider from 'react-rangeslider';

import './index.css';

export default class YieldCalculator extends Component {

    constructor(props, context) {
        super(props, context)
        this.state = {
          volume: 0
        }
    }

    handleOnChange = (value) => {
        this.setState({
          volume: value
        })
    }
    

    render() {

        let { volume } = this.state;

        return(
            <div className="yield-calc">
                <header className="header">Yield Calculator</header>
                <div className="body">
                    <hr/>
                    <Container fluid>
                        <Row>
                            <Col>Your Dia Stake</Col>
                            <Col>Return</Col>
                        </Row>
                        <Row>
                            <Col>1000123</Col>
                            <Col>+15%</Col>
                        </Row>
                        <Row>
                            <Col>
                                <Slider
                                    value={volume}
                                    onChange={this.handleOnChange}
                                />
                            </Col>
                        </Row>
                    </Container>
                </div>

            </div>
        )
    }

}