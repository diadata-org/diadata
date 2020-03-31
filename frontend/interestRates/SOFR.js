
var chart;
const maxPoints = 1000,
	timeUpdate = 1000;


function getRateInfo() {
	// Returns an object with information on the rate for labelling the plot
	rate = {
		name: 'Secured Overnight Financing Rate',
		symbol: 'SOFR',
	}
	return rate
}

function getApi(type) {
	if(type == 'historic') {
		return 'https://api.diadata.org/v1/interestrate/SOFR?dateInit=2018-11-01&dateFinal=2020-03-29'
	} else if(type == 'actual') {
        return 'https://api.diadata.org/v1/interestrate/SOFR/2020-03-30'
    }
}

// ------------------------------------------------------------------------------------------------
// Prefill with historic data
// ------------------------------------------------------------------------------------------------

function getHistoric(url, callback) {
    // getHistoric fetches historic data from the API
    
	// Instantiate request object
	var request = new XMLHttpRequest()
	request.open('GET', url, true)

	// Load data from GET request
	request.onload = function() {
		var data = JSON.parse(this.response)
		if(this.status == 200) {
			if (typeof callback === "function") {
				callback(data)
			}
		} else if(this.status == 404) {
			console.log('Not found error')
		}
	}
	request.onerror = function() {
		console.log('Request error.')
    }
    request.send()
}

getHistoric(getApi('historic'), function(obj) {
    // Each entry of obj corresponds to rate information at a specific timestamp.
	prefillArray = []
    for(i = 0; i < obj.length; i++) {  
        prefillArray.push([Date.parse(obj[i].Time), obj[i].Value]);
    }
    // Sort array by date ...
    prefillArray.sort()
    // ... and fill the chart.
    chart.series[0].setData(prefillArray)
});

// ------------------------------------------------------------------------------------------------
// Update asynchronously
// ------------------------------------------------------------------------------------------------

function requestData() {
    var xhr = $.ajax({
		url: getApi('actual'),
		type: 'GET',

		// If GET request is successful, add data to the chart
        success: function(point) {
			// @timeseries points to the first chart's series.
			// @shift restricts the number of displayed points.
			var timeseries = chart.series[0],
				shift = timeseries.data.length > maxPoints;
			
			// convert time (string) to Unix time for plotting
			var date = Date.parse(point.Time);
            console.log(point)
            
			// Append a data point to the chart's series if the timestamp is new
            L = chart.series[0].xData.length;
            
			if(L == 0) {
				chart.series[0].addPoint([date, point.Value], true, shift);
                console.log("Initial fill: " + point.Time)                
			} else if(L > 0 && date != chart.series[0].xData[L-1]){
				chart.series[0].addPoint([date, point.Value], true, shift);
				console.log("Length of array is: ", L)
				console.log("Updated point at: " + point.Time)
			}

			// Check for new data after @timeUpdate milliseconds
			setTimeout(requestData, timeUpdate); 
        },
		cache: false,
	});
}


// ------------------------------------------------------------------------------------------------
// Draw chart
// ------------------------------------------------------------------------------------------------

document.addEventListener('DOMContentLoaded', function() {
    
    chart = Highcharts.stockChart('container', {
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
            text: getRateInfo().name,
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
				name: getRateInfo().symbol,
				data: []
			},
		]
	});	
});

