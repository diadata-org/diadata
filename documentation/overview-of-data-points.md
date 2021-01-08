# Overview of Data Points

## Digital assets

<table>
  <thead>
    <tr>
      <th style="text-align:left">Data Type</th>
      <th style="text-align:left">Description</th>
      <th style="text-align:left">Methodology</th>
      <th style="text-align:left">API Documentation</th>
      <th style="text-align:left">Oracle Link</th>
      <th style="text-align:left">API Update Period</th>
      <th style="text-align:left">Oracle Update Period</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td style="text-align:left">Symbol</td>
      <td style="text-align:left">Symbol for cryptocurrency</td>
      <td style="text-align:left">Retrieval as part of a trading pair</td>
      <td style="text-align:left"><a href="https://api.diadata.org/v1/symbols">Symbols</a>
      </td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/oracle-documentation/crypto-assets">Symbol Oracle</a>
      </td>
      <td style="text-align:left">1 day</td>
      <td style="text-align:left">1 day</td>
    </tr>
    <tr>
      <td style="text-align:left">Quotation</td>
      <td style="text-align:left">Most recent price of a cryptocurrency in USD</td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/methodology/digital-assets/exchangeprices">Pricing Methodology</a>
      </td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/api-1/api-endpoints#quotation">Quotation</a>
      </td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/oracle-documentation/crypto-assets">Price Oracle</a>
      </td>
      <td style="text-align:left">2 min.</td>
      <td style="text-align:left">1 day</td>
    </tr>
    <tr>
      <td style="text-align:left">Exchange</td>
      <td style="text-align:left">Cryptocurrency exchange name</td>
      <td style="text-align:left">-</td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/api-1/api-endpoints#exchanges">Exchanges</a>
      </td>
      <td style="text-align:left">-</td>
      <td style="text-align:left">Depending on assignments or fundings on our platform</td>
      <td style="text-align:left">-</td>
    </tr>
    <tr>
      <td style="text-align:left">Chart Point</td>
      <td style="text-align:left">Coin Prices as processed by our pricing backend, organized by exchange
        and filter</td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/methodology/digital-assets/exchangeprices">Pricing Methodology</a>
      </td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/api-1/api-endpoints#chart-points">Chart Points</a>
      </td>
      <td style="text-align:left">-</td>
      <td style="text-align:left">2 min.</td>
      <td style="text-align:left">-</td>
    </tr>
    <tr>
      <td style="text-align:left">Supply</td>
      <td style="text-align:left">Circulating supply of cryptocurrencies</td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/methodology/digital-assets/supplynumbers">Supply Methodology</a>
      </td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/api-1/api-endpoints#supply">Supply</a>
      </td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/oracle-documentation/crypto-assets">Supply Oracle</a>
      </td>
      <td style="text-align:left">1 day</td>
      <td style="text-align:left">1 day</td>
    </tr>
    <tr>
      <td style="text-align:left">Trade Volume</td>
      <td style="text-align:left">Trading volume of a cryptocurrency</td>
      <td style="text-align:left">Added volumes of raw trading data after filtering as described <a href="https://docs.diadata.org/documentation/methodology/digital-assets/exchangeprices">here</a>
      </td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/api-1/api-endpoints#trade-volume">Trade Volumes</a>
      </td>
      <td style="text-align:left">-</td>
      <td style="text-align:left">1 day</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">Raw Crypto Trade</td>
      <td style="text-align:left">Trading data</td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/methodology/digital-assets/cryptocurrency-trading-data">Trading Data Methodology</a>
      </td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/api-1/api-endpoints#raw-crypto-trades">Raw Crypto Trades</a>
      </td>
      <td style="text-align:left">-</td>
      <td style="text-align:left">2 min.</td>
      <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">CVI Index</td>
      <td style="text-align:left">Crypto volatility index measuring volatility in cryptocurrency markets</td>
      <td
      style="text-align:left"><a href="https://docs.diadata.org/documentation/methodology/digital-assets/cvi">CVI Methodology</a>
        </td>
        <td style="text-align:left"><a href="https://docs.diadata.org/documentation/api-1/api-endpoints#cvi-index">CVI Index</a>
        </td>
        <td style="text-align:left"><em>soon</em>
        </td>
        <td style="text-align:left">5 min.</td>
        <td style="text-align:left"></td>
    </tr>
    <tr>
      <td style="text-align:left">DeFi Interest Rate</td>
      <td style="text-align:left">Lending/Borrowing rates on DeFi protocols</td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/methodology/digital-assets/cryptocurrency-trading-data#lending-borrowing-data">Lending Protocol Methodology</a>
      </td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/api-1/api-endpoints#defi-interest-rate">DeFi Lending Rate</a> 
      </td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/oracle-documentation/defi-protocol-rates-and-states">DeFi Lending Oracle</a>
      </td>
      <td style="text-align:left">1 min.</td>
      <td style="text-align:left">1 day</td>
    </tr>
    <tr>
      <td style="text-align:left">DeFi Lending Protocol State</td>
      <td style="text-align:left">Total value locked in a lending/borrowing protocol</td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/methodology/digital-assets/cryptocurrency-trading-data#lending-borrowing-data">Lending Protocol Methodology</a>
      </td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/api-1/api-endpoints#defi-lending-state">DeFi Lending Protocol</a>
      </td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/oracle-documentation/defi-protocol-rates-and-states">DeFi  Lending Oracle</a>
      </td>
      <td style="text-align:left">1 min.</td>
      <td style="text-align:left">1 day</td>
    </tr>
    <tr>
      <td style="text-align:left">Guest Quotation</td>
      <td style="text-align:left">Quotations accumulated with methodology of partner sites</td>
      <td style="text-align:left">-</td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/api-1/api-endpoints#guest-quotation">Guest Quotation</a>
      </td>
      <td style="text-align:left">
        <p><a href="https://etherscan.io/address/0x48760771feda4be44a6ed3bff13ecbc445159b1d">CoinMarketCap</a>
        </p>
        <p><a href="https://etherscan.io/address/0x48760771feda4be44a6ed3bff13ecbc445159b1d">Oracle</a>
        </p>
        <p><a href="https://docs.diadata.org/documentation/oracle-documentation/guest-quotations">Coingecko Oracle</a>
        </p>
      </td>
      <td style="text-align:left">2 min.</td>
      <td style="text-align:left">1 day</td>
    </tr>
    <tr>
      <td style="text-align:left">Farming Pool</td>
      <td style="text-align:left">Rates and balances of crypto farming pools</td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/methodology/digital-assets/return-rates-in-crypto-farming">Pool Rates</a>
      </td>
      <td style="text-align:left">
        <p><a href="https://docs.diadata.org/documentation/api-1/api-endpoints#farming-pools">List of Pools</a>
          <br
          />--</p>
        <p><a href="https://docs.diadata.org/documentation/api-1/api-endpoints#farming-pool-data">Pool Data</a>
        </p>
      </td>
      <td style="text-align:left"><a href="https://docs.diadata.org/documentation/oracle-documentation/farming-pools">Farming Pool Oracle</a>
      </td>
      <td style="text-align:left">2 min.</td>
      <td style="text-align:left">1day</td>
    </tr>
  </tbody>
</table>

## Traditional Assets

| Data Type | Description | Methodology | API Documentation | Oracle Link | Update Period |
| :--- | :--- | :--- | :--- | :--- | :--- |
| Interest Rates | Interest rate types from  central banks such as ECB  | [Rates methodology](https://docs.diadata.org/documentation/methodology/traditional-assets/overnight-rates) | [Interest Rates](https://docs.diadata.org/documentation/api-1/api-endpoints#interest-rates) | _-_ | 1 day |
| Interest Rate | Particular value of an interest rate at a given time | [Rates methodology](https://docs.diadata.org/documentation/methodology/traditional-assets/overnight-rates) | [Interest Rate](https://docs.diadata.org/documentation/api-1/api-endpoints#interest-rate) | \_\_[_soon_](https://docs.diadata.org/documentation/oracle-documentation/interest-rates)\_\_ | 1 day |
| Compounded Index | Interest rate compounded since first publication date | [Compounding methodology](https://docs.diadata.org/documentation/methodology/traditional-assets/compounded-rates#standard-methodology) | [Compounded Index](https://docs.diadata.org/documentation/api-1/api-endpoints#compounded-index) | _soon_ | 1 day |
| Compounded Average | Average interest rate compounded over a variable time range | [Compounding methodology](https://docs.diadata.org/documentation/methodology/traditional-assets/compounded-rates#standard-methodology) | [Compounded Average](https://docs.diadata.org/documentation/api-1/api-endpoints#compounded-average) | - | 1 day |
| Compounded DIA Average | Average interest rate compounded over a variable time range, using DIA methodology | [DIA Compounding methodology](https://docs.diadata.org/documentation/methodology/traditional-assets/compounded-rates#dia-methodology) | [Compounded DIA Average](https://docs.diadata.org/documentation/api-1/api-endpoints#compounded-average-using-dia-method) | - |  1 day |
| Fiat Currency Exchange | Exchange rate for fiat currency | Unprocessed retrieval from ECB API | [Fiat currency Exchange](https://docs.diadata.org/documentation/api-1/api-endpoints#fiat-currency-exchange-rates) | [Fiat Price Oracle](https://docs.diadata.org/documentation/oracle-documentation/fiat-prices) | 1 day |

