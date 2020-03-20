
function getApi() {
	return 'https://api.diadata.org/v1/interestrate/SOFR/2020-03-14'
}


function getData(url, callback) {

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

getData(getApi(), function(obj) {
	console.log(obj)
});



document.addEventListener('DOMContentLoaded', () => {
		Highcharts.chart('container', {
			chart: {
				type: 'area'
			},
			credits: {
				text: 'DIAdata',
				href: 'https://diadata.org'
			},
			title: {
				text: 'My first chart'
			},
			xAxis: {
				title: {
					text: 'time'
				},
			},
			yAxis: {
				title: {
					text: 'value'
				},

			},
			
			series: [
				{
				name: 'SOFR',
				data: [[1, 12], [2,10], [3, 20], [4, 40], [5, 9]]
			},
			{
				data: [[1, 1], [2,10], [3, 10], [4, 20], [5, 1]],
				name: 'ESTER'
			}
			]
		});
});
