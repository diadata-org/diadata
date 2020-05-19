// ------------------------------------------------------------------------------------------------
// General variables and functions
// ------------------------------------------------------------------------------------------------

let today = new Date().toISOString().slice(0, 10);
let yesterday = new Date(Date.now() - 864e5).toISOString().slice(0,10);

// getHistoric fetches historic data from our API with address @url
function getHistoric(url, callback) {
    
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

// requestData(rate) asynchrononously loads new data points into the chart.
function requestData(rate) {
    $.ajax({
		cache: true,
		url: rate.urlActual,
		type: 'GET',

		// If GET request is successful, add data to the chart
        success: function(point) {

			// convert time (string) to Unix time for plotting
			var date = Date.parse(point.EffectiveDate);
            
			// Append a data point to the chart's series if the timestamp is new
			L = window['chart' + rate.symbol].series[0].xData.length;
			
			if(L == 0) {
				window['chart' + rate.symbol].series[0].addPoint([date, point.Value]);
			} else if(L > 0 && date != window['chart' + rate.symbol].series[0].xData[L-1]){
				window['chart' + rate.symbol].series[0].addPoint([date, point.Value]);
			}
			console.log("Updated " + rate.symbol);
			window['chart' + rate.symbol].redraw();
			setTimeout(function(){requestData.call(this, rate)}, rate.timeUpdate); 
        },
	});
}

function makechart(rate) {
	window['chart' + rate.symbol] = Highcharts.stockChart(rate.container, {
		rangeSelector: {
			buttonTheme: {
				   width: 20,
			},
			inputBoxWidth: 75,
	   	}, 
		chart: {
            type: 'spline',
            events: {
				load: function(){
					requestData.call(this, rate)
				}
			}
		},
		credits: {
			text: 'DIADATA',
			href: 'https://diadata.org'
		},
        title: {
			text: rate.name,
			style: {
				fontSize: '20px',
			},
        },
        xAxis: {
            tickPixelInterval: 150,
			maxZoom: 20 * 1000,
			title: {
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
				name: rate.symbol,
				data: []
			},
		]
	});	
}



// ------------------------------------------------------------------------------------------------
// Rate specific information
// ------------------------------------------------------------------------------------------------

InfoSOFR = {
	name: 'Secured Overnight Financing Rate',
	symbol: 'SOFR',
	container: 'containerSOFR',
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
	window['chart' + InfoSOFR.symbol].series[0].setData(prefillArray)
	// Redraw for correct display of boxes
	window['chart' + InfoSOFR.symbol].redraw();
});

document.addEventListener('DOMContentLoaded', makechart(InfoSOFR));




InfoESTR = {
	name: 'Euro Short-Term Rate',
	symbol: '€STR',
	container: 'containerESTR',
	timeUpdate: 10000,
	urlHistoric: 'https://api.diadata.org/v1/interestrate/ESTER?dateInit=2018-10-01&dateFinal=' + yesterday,
	urlActual: 'https://api.diadata.org/v1/interestrate/ESTER/' + today,
}


getHistoric(InfoESTR.urlHistoric, function(obj) {
	prefillArray = []
    for(i = 0; i < obj.length; i++) {  
        prefillArray.push([Date.parse(obj[i].EffectiveDate), obj[i].Value]);
    }
    prefillArray.sort()
	window['chart' + InfoESTR.symbol].series[0].setData(prefillArray)
	window['chart' + InfoESTR.symbol].redraw();
});

document.addEventListener('DOMContentLoaded', makechart(InfoESTR));