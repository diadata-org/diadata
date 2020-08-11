import React, { Component} from 'react';

// css
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';

// components
import YieldCalculator from './components/yield';

// web3
import getWeb3 from './helpers/web3';
import { getDiaRate } from './helpers/api';

export default class App extends Component {

  render() {

    return (
      <div className="App">
        <YieldCalculator getWeb3={getWeb3} getDiaRate={getDiaRate}/>
      </div>
    );
  }
  
}
