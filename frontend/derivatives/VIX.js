var chart;

IndexInfo = {
	name: 'Volatility Index',
	symbol: 'VIX',
	timeInterval: 100000,
	timeUpdate: 10000
}

let now = Math.round(Date.now()/1000);
let beforeNow = (now - IndexInfo.timeInterval).toString();
now = now.toString()

getApi =  {
	prefill: 'https://api.diadata.org/v1/cviIndex?starttime=1&endtime=1586085310',
	postfill: 'https://api.diadata.org/v1/cviIndex?starttime=1586085311&endtime=' + beforeNow,
	update: 'https://api.diadata.org/v1/cviIndex?starttime=' + beforeNow + '&endtime=' + now,
}

// ------------------------------------------------------------------------------------------------
// Prefill with data
// ------------------------------------------------------------------------------------------------

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

// Fill chart with data upon loading screen
getHistoric(getApi.prefill, function(obj) {
    // Each entry of obj corresponds to rate information at a specific timestamp
	prefillArray = []
    for(i = 0; i < obj.length; i++) {  
        prefillArray.push([Date.parse(obj[i].Timestamp), obj[i].Value]);
	}
    // Fill the chart
    chart.series[0].setData(prefillArray)
});


// ------------------------------------------------------------------------------------------------
// As soon as chart is loaded add all data and then update asynchonously
// ------------------------------------------------------------------------------------------------

function requestData() {

	// Load older data
	getHistoric(getApi.postfill, function(obj) {
		postfillArray = []
		for(i = 0; i < obj.length; i++) {  
			postfillArray.push([Date.parse(obj[i].Timestamp), obj[i].Value]);
		}
		chart.series[0].setData(postfillArray)
	});
		
    $.ajax({
		cache: true,
		url: getApi.update,
		type: 'GET',

		// If GET request is successful, add data to the chart
        success: function(points) {

			console.log("success")
			
			// Get the actual value
			if(points.length > 1) {
				var point = points[points.length-1];
			} else {
				var point = points;
			}
			
			// convert time (string) to Unix time for plotting
			var date = Date.parse(point.Timestamp);
            console.log(point)
            
			// Append a data point to the chart's series if the timestamp is new
			L = chart.series[0].xData.length;
			if(L == 0) {
				chart.series[0].addPoint([date, point.Value]);
                console.log("Initial fill: " + point.Timestamp)                
			} else if(L > 0 && date != chart.series[0].xData[L-1]){
				chart.series[0].addPoint([date, point.Value]);
				console.log("Updated point at: " + point.Timestamp)
			}

			// Check for new data after @timeUpdate milliseconds
			setTimeout(requestData, IndexInfo.timeUpdate); 
        },
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
				load: requestData,
			}
		},
		credits: {
			text: 'DIAdata',
			href: 'https://diadata.org'
		},
        title: {
            text: IndexInfo.name,
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
				name: IndexInfo.symbol,
				data: []
			},
		]
	});	
});

