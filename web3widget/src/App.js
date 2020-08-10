import React, { Component} from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';

import './App.css';

import YieldCalculator from './components/yield'

export default class App extends Component {

  render() {

    return (
      <div className="App">
        <YieldCalculator/>
      </div>
    );

  }
  
}
