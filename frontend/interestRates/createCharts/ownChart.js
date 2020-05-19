// ------------------------------------------------------------------------------------------------
// General variables and functions
// ------------------------------------------------------------------------------------------------

// // Load meta information on rates
var dateUrl = 'http://localhost:8081/v1/interestrates';
$.holdReady(true);
var firstPublications = null;
$.getJSON(dateUrl, function(data) {
    firstPublications = data;
    $.holdReady(false);
});

let today = new Date().toISOString().slice(0, 10);
var yourOwnChart;
var firstPublications = {};

function addDays(date, days) {
    var result = new Date(date).getTime();
    result += days * 864e5
    result = new Date(result);
    return result.toISOString()
}

// getHistoric fetches historic data from our API with address @url
function getData(url, callback) {
    
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

function makechart(rate) {
	yourOwnChart = Highcharts.stockChart(rate.container, {
		rangeSelector: {
			buttonTheme: {
				   width: 20,
			},
			inputBoxWidth: 75,
	   	}, 
		chart: {
            type: 'spline',
        
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
				name: rate.name,
				data: []
            },
            // {
            //     name: 'ESTER',
			// 	data: [[1, 1], [2,10], [3, 10], [4, 20], [5, 1]]
			// }
		]
	});	
}

// ------------------------------------------------------------------------------------------------
// First fill of chart when loading the page
// ------------------------------------------------------------------------------------------------

// Rate info for the first fill
var RateInfo = {
	name: 'SOFR30',
	container: 'yourOwnContainer',
    firstPublication: "2018-04-03",
    url: 'http://localhost:8081/v1/compoundedAvg/SOFR/30/360?dateInit=2018-05-14&dateFinal=' + today,    
};

// Initial fill
getData(RateInfo.url, function(obj) {
    prefillArray = []
    for(i = 0; i < obj.length; i++) {  
        prefillArray.push([Date.parse(obj[i].EffectiveDate), obj[i].Value]);
    }
    prefillArray.sort()
    yourOwnChart.series[0].setData(prefillArray)
    // yourOwnChartSOFR.redraw();               
});
makechart(RateInfo);

// ------------------------------------------------------------------------------------------------
// Update upon clicking button
// ------------------------------------------------------------------------------------------------
function updateChart() {

    // Retrieve user data --------------------------------------------------------------------
    var lenPeriod = document.getElementById("lenPeriod").value;
    var dpy = document.getElementById("dpy").value;
    var symbol = document.getElementById("symbol").value;
    

    // update rate information ---------------------------------------------------------------
    RateInfo.name = symbol + lenPeriod;
    // retrieve first publication date
    const found = Object.values(firstPublications).find(element => element.Symbol == symbol);
    RateInfo.firstPublication = found.FirstDate.slice(0,10);
    // Increase initial date according to observation period
    dateInit = addDays(RateInfo.firstPublication, lenPeriod).slice(0,10);            
    RateInfo.url = 'http://localhost:8081/v1/compoundedAvg/' + symbol + '/' + lenPeriod + '/' + dpy + '?dateInit=' + dateInit + '&dateFinal=' + today;
    
    // Fill the chart with updated information
    getData(RateInfo.url, function(obj) {
        var prefillArray = [];
        for(i = 0; i < obj.length; i++) {  
            prefillArray.push([Date.parse(obj[i].EffectiveDate), obj[i].Value]);
        }
        prefillArray.sort(function(a,b) { return a - b; });
        yourOwnChart.series[0].setData(prefillArray);               
    });

    //  // Fill the chart with reference rate
    //  getData('http://localhost:8081/v1/interestrate/' + symbol + '/' + '?dateInit=' + dateInit + '&dateFinal=' + today, function(obj) {
    //     var prefillArray = [];
    //     for(i = 0; i < obj.length; i++) {  
    //         prefillArray.push([Date.parse(obj[i].EffectiveDate), obj[i].Value]);
    //     }
    //     prefillArray.sort();
    //     yourOwnChart.series[1].setData(prefillArray);
                   
    // });

    makechart(RateInfo);

};
