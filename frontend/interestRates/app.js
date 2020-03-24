
var chart;
const maxPoints = 50,
	timeUpdate = 1000;

function initialFill() {
	// Get the last @maxPoints from the API in order to have a
	// full series when the chart is loaded
}

function getRate() {
	// Returns an object with information on the rate for labelling the plot
	rate = {
		name: 'Ethereum',
		symbol: 'ETH',
	}
	return rate
}

function getApi(type) {
	if(type == 'rate') {
		return 'https://api.diadata.org/v1/interestrate/SOFR/2020-03-23'
	} else if(type == 'quotation'){
		return 'https://api.diadata.org/v1/quotation/ETH'
	}
}

function requestData() {
    var xhr = $.ajax({
		url: getApi('quotation'),
		type: 'GET',

		// If GET request is successful, add data to the chart
        success: function(point) {

			// @timeseries points to the first chart's series.
			// @shift restricts the number of displayed points.
			var timeseries = chart.series[0],
				shift = timeseries.data.length > maxPoints;
			
			// convert time (string) to Unix time for plotting
			var date = Date.parse(point.Time);

			// Append a data point to the chart's series if its timestamp is new
			L = chart.series[0].xData.length;
			if(L == 0) {
				chart.series[0].addPoint([date, point.Price], true, shift);
				console.log("Initial fill: " + point.Time)
			} else if(L > 0 && date != chart.series[0].xData[L-1]){
				chart.series[0].addPoint([date, point.Price], true, shift);
				console.log("Updated point at: " + point.Time)
			}

            // Check for new data after @timeUpdate milliseconds
			setTimeout(requestData, timeUpdate); 
        },
		cache: false,
	});
}


document.addEventListener('DOMContentLoaded', function() {
    chart = Highcharts.chart('container', {
        chart: {
            type: 'spline',
            events: {
                load: requestData
            }
		},
		credits: {
			text: 'DIAdata',
			href: 'https://diadata.org'
		},
        title: {
            text: getRate().name,
        },
        xAxis: {
            type: 'datetime',
            tickPixelInterval: 150,
			maxZoom: 20 * 1000,
        },
        yAxis: {
            minPadding: 0.2,
            maxPadding: 0.2,
            title: {
                text: 'Value',
                margin: 80
            }
        },
        series: [
			{
				name: getRate().symbol,
            	data: []
			},
		]
	});	
});