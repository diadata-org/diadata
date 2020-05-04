var chartCVI;

IndexInfo = {
	name: 'Crypto Volatility Index',
	symbol: 'CVI',
	timeInterval: 30000,
	timeUpdate: 10000
}

let now = Math.round(Date.now()/1000);
let beforeNow = (now - IndexInfo.timeInterval).toString();
now = now.toString()

getApi =  {
	prefill: 'https://api.diadata.org/v1/cviIndex?starttime=1&endtime=' + beforeNow,
	update: 'https://api.diadata.org/v1/cviIndex?starttime=' + beforeNow + '&endtime=' + now,
}

// ------------------------------------------------------------------------------------------------
// Prefill with data
// ------------------------------------------------------------------------------------------------

function getHistoric(url, callback) {
    var request = new XMLHttpRequest()
	request.open('GET', url, true)

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
	console.log("prefill");
    prefillArray = []
    for(i = 0; i < obj.length; i++) {  
        prefillArray.push([Date.parse(obj[i].Timestamp), obj[i].Value]);
	}
	chartCVI.series[0].setData(prefillArray);
});


// ------------------------------------------------------------------------------------------------
// As soon as chart is loaded update asynchonously
// ------------------------------------------------------------------------------------------------

function requestData() {
	let now = Math.round(Date.now()/1000);
	let beforeNow = (now - IndexInfo.timeInterval).toString();
    $.ajax({
		cache: true,
		url: 'https://api.diadata.org/v1/cviIndex?starttime=' + beforeNow + '&endtime=' + now,
		type: 'GET',

        success: function(points) {
			if(points.length > 1) {
				var point = points[points.length-1];
			} else {
				var point = points;
			}
			var date = Date.parse(point.Timestamp);
            console.log(point)
            
			L = chartCVI.series[0].xData.length;
			if(L == 0) {
				chartCVI.series[0].addPoint([date, point.Value]);
                console.log("Initial fill: " + point.Timestamp)                
			} else if(L > 0 && date != chartCVI.series[0].xData[L-1]){
				chartCVI.series[0].addPoint([date, point.Value]);
				console.log("Updated point at: " + point.Timestamp)
			}

			setTimeout(requestData, IndexInfo.timeUpdate); 
        },
	});
}

// ------------------------------------------------------------------------------------------------
// Draw chart
// ------------------------------------------------------------------------------------------------

document.addEventListener('DOMContentLoaded', function() {
    
    chartCVI = Highcharts.stockChart('container', {
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

