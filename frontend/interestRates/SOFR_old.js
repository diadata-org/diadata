// ------------------------------------------------------------------------------------------------
// General variables and functions
// ------------------------------------------------------------------------------------------------

let today = new Date().toISOString().slice(0, 10);
let yesterday = new Date(Date.now() - 864e5).toISOString().slice(0,10);

function getHistoric(url, callback) {
    // getHistoric fetches historic data from the API
    
	// Instantiate request object
	var request = new XMLHttpRequest()
	request.open('GET', url, true)

	// Load data in GET request
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

function requestData(rate) {
	console.log("my object is: ", rate)
    $.ajax({
		cache: true,
		url: rate.urlActual,
		type: 'GET',

		// If GET request is successful, add data to the chart
        success: function(point) {

			// convert time (string) to Unix time for plotting
			var date = Date.parse(point.EffectiveDate);
            console.log("point is: ", point);
            
			// Append a data point to the chart's series if the timestamp is new
            // L = chartSOFR.series[0].xData.length;
			L = chartSOFR.series[0].xData.length;
			
			if(L == 0) {
				chartSOFR.series[0].addPoint([date, point.Value]);
                console.log("Initial fill: " + point.EffectiveDate);                
			} else if(L > 0 && date != window['chart' + rate.symbol].series[0].xData[L-1]){
				chartSOFR.series[0].addPoint([date, point.Value]);
				console.log("Updated point at: " + point.EffectiveDate);
			}

			// Check for new data after @timeUpdate milliseconds
			setTimeout(function(){requestData.call(this, rate)}, rate.timeUpdate); 
        },
	});
}


// ------------------------------------------------------------------------------------------------
// Rate specific information
// ------------------------------------------------------------------------------------------------

InfoSOFR = {
	name: 'Secured Overnight Financing Rate',
	symbol: 'SOFR',
	timeUpdate: 10000,
	urlHistoric: 'https://api.diadata.org/v1/interestrate/SOFR?dateInit=2018-11-01&dateFinal=' + yesterday,
	urlActual: 'https://api.diadata.org/v1/interestrate/SOFR/' + today,
}


getHistoric(InfoSOFR.urlHistoric, function(obj) {
    // Each entry of obj corresponds to rate information at a specific timestamp.
	prefillArray = []
    for(i = 0; i < obj.length; i++) {  
        prefillArray.push([Date.parse(obj[i].EffectiveDate), obj[i].Value]);
    }
    // Sort array by date ...
    prefillArray.sort()
    // ... and fill the chart.
	chartSOFR.series[0].setData(prefillArray)
});



document.addEventListener('DOMContentLoaded', function() {
	
    chartSOFR = Highcharts.stockChart('container', {
        chart: {
            type: 'spline',
            events: {
				load: function(){
					requestData.call(this, InfoSOFR)
				}
			}
		},
		credits: {
			text: 'DIADATA',
			href: 'https://diadata.org'
		},
        title: {
            text: InfoSOFR.name,
        },
        xAxis: {
            tickPixelInterval: 150,
			maxZoom: 20 * 1000,
			title: {
				text: 'Time',
				margin: 10,
			}
		},
        yAxis: {
            minPadding: 0.2,
            maxPadding: 0.2,
            title: {
                text: 'Index value',
                margin: 80
			}	
        },
        series: [
			{
				name: InfoSOFR.symbol,
				data: []
			},
		]
	});	
});

