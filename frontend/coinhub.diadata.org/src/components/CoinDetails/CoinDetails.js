import axios from 'axios';
import router from '@/router';
import { AtomSpinner } from 'epic-spinners';
import sortBy from 'lodash/sortBy';
import shared from  '@/shared/shared';
import moment from 'moment';
import getSymbolFromCurrency from 'currency-symbol-map';
import Highcharts from 'highcharts';
import stockInit from 'highcharts/modules/stock';
import numeral from 'numeral';

stockInit(Highcharts)
export default {
	components: {
		AtomSpinner
	},
	name: 'CoinDetails',
	props: {},
	data() {
		return {
			exchange_fields: [
				{ key: 'Name', label: 'Exchange', sortable: true },
				{ key: 'PriceFormatted', label: 'Price', sortable: true },
				{ key: 'Volume24', label: 'Volume (24h)', sortable: true },
				{ key: 'TimeFormatted', label: 'Last Updated', sortable: true },
				{ key: 'show_trades', label: 'Trades', sortable: false },
			],
			exchanges: [],
			last_trade_fields: [
				{ key: 'Pair', label: 'Pair', sortable: true },
				{ key: 'EstimatedPrice', label: 'Price', sortable: true },
				{ key: 'Order', label: 'Order', sortable: true },
				{ key: 'Volume', label: 'Volume', sortable: true },
				{ key: 'TimeFormatted', label: 'Time', sortable: true },
			],
			loading: true,
			errored: false,
			coinDetails:{},
			coinSymbol: '',
			coindata: null,
			selectedCurrency:'',
			selectedAlgorithm: '',
			selectedAlgorithmName: '',
			selectedExchange: '',
			chartAllOptions: {},
			chartSimexOptions: {},
			rateArray: [],
			error:'',
			showAllCharts: false,
			algorithms: [],
			exchangeNames: [],
			currencies: [],
			requestURL: ""
        };
	},
	created() {
		this.coinSymbol = this.$route.params.coinSymbol;
	},
	mounted() {
		if(this.$route.params.coinRank) {
			localStorage.rank = this.$route.params.coinRank;
		}
		else {
			localStorage.rank = 'N/A';
		}

		// fetch the coin details
		this.fetchCoinDetails();
	},
	methods: {
		formatPairData() {

			if(localStorage.selectedCurrency) {
				this.selectedCurrency = localStorage.selectedCurrency;
			}
			else{
				this.selectedCurrency = "USD";
			}

            if(localStorage.selectedAlgorithm) {
                this.selectedAlgorithm = localStorage.selectedAlgorithm;
            }
            else{
                this.selectedAlgorithm = "MA120";
                this.selectedAlgorithmName = "Moving Avg, 120s";
            }

            if(localStorage.selectedExchange && localStorage.selectedExchange !== "All") {
                this.selectedExchange = localStorage.selectedExchange;
                this.requestURL = `/v1/chartPoints/${this.selectedAlgorithm}/${this.selectedExchange}/${this.coinSymbol.toUpperCase()}`;
            }
            else{
                this.selectedExchange = "All";
                this.requestURL = `/v1/chartPointsAllExchanges/${this.selectedAlgorithm}/${this.coinSymbol.toUpperCase()}`;
            }

			let {Change, Coin, Rank, Exchanges} = this.coindata;

			this.rateArray = Change.USD;
			this.currencies = shared.getCurrencies(this.rateArray);
			this.algorithms = [{
				displayName: "Moving Avg",
                urlString: "MA120"
            },{
				displayName: "Outlier Cleaned (IQR acceptable range)",
                urlString: "MEDIR120"
            }];

			// format the coin details
			const coinPrice = shared.calculateCurrencyFromRate(Coin.Price,this.rateArray,this.selectedCurrency,"today");
			const coinPriceFormatted = shared.formatCurrency(coinPrice,this.selectedCurrency);
			const coinPriceYesterday = shared.calculateCurrencyFromRate(Coin.PriceYesterday,this.rateArray,this.selectedCurrency,"yesterday");
			const change24 = (coinPrice  - coinPriceYesterday) / coinPriceYesterday * 100;
			const change24Formatted = shared.formatChange24(change24);
			const volume24Formatted = shared.formatMarketCapAndVolume24(shared.calculateCurrencyFromRate(Coin.VolumeYesterdayUSD,this.rateArray,this.selectedCurrency,"yesterday"),this.selectedCurrency);
			const circulatingSupplyFormattedWithoutSymbol = shared.formatCirculatingSupply(Coin.CirculatingSupply, undefined);
			this.coinDetails = {
				coinName: Coin.Name,
				coinSymbol: Coin.Symbol,
				coinPriceFormatted,
				change24,
				change24Formatted,
				rank: Rank, // localStorage.rank,
				volume24Formatted,
				circulatingSupplyFormattedWithoutSymbol,
			};

			// format the exchanges
			Exchanges.forEach((exchange)=>{
				exchange.PriceFormatted = shared.formatCurrency(shared.calculateCurrencyFromRate(exchange.Price,this.rateArray,this.selectedCurrency,"today"),this.selectedCurrency),
				exchange.Volume24 = shared.formatMarketCapAndVolume24(shared.calculateCurrencyFromRate(exchange.VolumeYesterdayUSD,this.rateArray,this.selectedCurrency,"yesterday"),this.selectedCurrency),
				exchange.TimeFormatted = shared.formatDateTime(exchange.Time,"dddd, MMMM Do YYYY, h:mm:ss a");

				// format the last trades too
				exchange.LastTrades.forEach((lastTrade) => {
						lastTrade.EstimatedPrice = shared.formatCurrency(shared.calculateCurrencyFromRate(lastTrade.EstimatedUSDPrice,this.rateArray,this.selectedCurrency,"today"),this.selectedCurrency);
						lastTrade.TimeFormatted = shared.formatDateTime(lastTrade.Time,"h:mm:ss a");
				});
			});

			this.exchangeNames = Exchanges.map((elem) => {
                return elem.Name;
			} );

			Exchanges = sortBy(Exchanges, 'VolumeYesterdayUSD').reverse();
			this.exchanges = Exchanges;



			// finally fetch the chart details
			this.fetchCoinChartDetails();
		},
		async fetchCoinDetails() {
			try {
				const response = await axios.get(shared.getApi()+`/v1/symbol/${this.coinSymbol.toUpperCase()}`);
				this.coindata = response.data;
				this.formatPairData();
			}
			catch (error) {
				console.log(error);
				this.errored = true;
			}
		},
		async fetchCoinChartDetails() {

			try {
				let response1 = await axios.get(shared.getApi()+this.requestURL);

				const price = 'Price (' + this.selectedCurrency + ')';
				const currencySymbol  = getSymbolFromCurrency(this.selectedCurrency);

				if(response1.data !== undefined) {
					const MA120AllArray = this.formatChartValues(response1.data.DataPoints[0].Series[0].values);

					this.chartAllOptions = {
						chart: {
							zoomType: 'x'
						},
						rangeSelector: {
							buttons: [{
								type: 'ytd',
								text: 'YTD'
							}, {
								type: 'month',
								count: 1,
								text: '1M'
							},{
								type: 'day',
								count: 7,
								text: '7D'
							},{
								type: 'day',
								count: 1,
								text: '1D'
							}],
							selected: 2
						},
						xAxis: {
							type: 'datetime',
							title: {
								text: 'Time'
							},
							ordinal: false,
						},
						yAxis: {
							title: {
								text: price
							},
							plotLines: [{
									value: 0,
									width: 2,
									color: 'silver'
							}]
						},
						tooltip: {
							formatter: function () {
								var s = '';
								var sc = this.selectedCurrency;
								this.points.forEach(function (element) {
									s += currencySymbol + " ";
									s += (element.y < 1) ? numeral(element.y).format('0.0[0000]') : numeral(element.y).format('0,0.00');
								});

								return s;
							},
						},
						series: [{
							name: "2 Minute MA",
							data: MA120AllArray,
						}]
					};

					// simex
					this.showAllCharts = true;

					this.chartSimexOptions = {
					chart: {
						zoomType: 'x'
					},
					rangeSelector: {
						buttons: [{
							type: 'ytd',
							text: 'YTD'
						}, {
							type: 'month',
							count: 1,
							text: '1M'
						},{
							type: 'day',
							count: 7,
							text: '7D'
						},{
							type: 'day',
							count: 1,
							text: '1D'
						}],
						selected: 3
					},
					title: {
						text: 'Simex'
					},
					subtitle: {
						text: ''
					},
					xAxis: {
						type: 'datetime',
						dateTimeLabelFormats: { // don't display the dummy year
							month: '%e. %b',
							year: '%b'
						},
						title: {
							text: 'Time'
						}
					},
					yAxis: {
						title: {
							text: price
						},
						min: 0
					},
					tooltip: {
						headerFormat: '<b>{series.name}</b><br>',
						pointFormat: `{point.x:%e. %b}: ${currencySymbol }{point.y:.2f} `
					},
					series: [{
						name: "MA120",
						data: []
					},{
						name: "2 Min. MA",
						data: []
					}]};
				}
				this.loading = false;
			} catch(error) {
				console.log(error);
				this.loading = false;
				this.errored = true;
			}
		},
		formatChartValues(chartValues) {
			let formattedValues = [];
			chartValues = chartValues.reverse();

			chartValues.forEach((chartValue) => {
				const UTCDate = new Date(chartValue[0]).valueOf();
				const price = parseFloat(shared.calculateCurrencyFromRate(chartValue[4],this.rateArray,this.selectedCurrency,"today"));
				formattedValues.push([UTCDate,price]);
			});

			return formattedValues;
		},
		switchCurrencies : function(selectedCurrency){
			this.selectedCurrency = selectedCurrency;
			localStorage.selectedCurrency = selectedCurrency;
			this.formatPairData();
		},
		switchAlgorithm : function(selectedAlgorithm){
			this.selectedAlgorithm = selectedAlgorithm.urlString;
			this.selectedAlgorithmName = selectedAlgorithm.displayName;
			localStorage.selectedAlgorithm = selectedAlgorithm.urlString;
			this.formatPairData();
		},
		switchExchange : function(selectedExchange){
			this.selectedExchange = selectedExchange;
			localStorage.selectedExchange = selectedExchange;
			this.formatPairData();
		},
	},
};
