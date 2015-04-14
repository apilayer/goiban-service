$(document).ready(function() {
	$("#start_validation").click(function() {
		var iban = $('#iban_input').val();
		if(iban.length < 1) {
			$('.form-group-iban-input').addClass("has-error");
			return;
		} else {
			$('.form-group-iban-input').removeClass("has-error").removeClass("has-success").removeClass("has-warning");
		}
		goiban.validate(iban, function(resultJSON) {
			if(resultJSON.valid) {

				$('#text_result_container').val("IBAN is valid.");
				if(resultJSON.bankData && resultJSON.bankData.bic) {
					$('#bic_result_container').val(resultJSON.bankData.bic);
				} else {
					$('#bic_result_container').val("Not available.");
				}
				$('.form-group-iban-input').addClass("has-success");
			} else {
				$('#text_result_container').val("IBAN is not valid!");
				$('.form-group-iban-input').addClass("has-error");
			}
			$('#result_container').val(JSON.stringify(resultJSON, null, " "));
		});

	});

	var ctx = document.getElementById('chart').getContext('2d');
	var chart24h,
	 knownLabels;

	function updateChart() {
		$.get('http://openiban.com/metrics').then(function (data) {
			var metrics = goiban.getMetrics24h(data);
			if(!chart24h || knownLabels.length != metrics.labels.length) {
				knownLabels = metrics.labels;
				chart24h = new Chart(ctx).Bar(metrics);
				return;
			}

			_.each(metrics.datasets[0].data, function (value, index) {
				chart24h.datasets[0].bars[index].value = value;
			});

			//chart24h.dataset[0] = metrics.dataset[0];
			chart24h.update();
		});
	}
	
	updateChart();
	setInterval(updateChart, 5000);

});

function getCount(x, key) {
	return {country: key, count: x.Count};
}

function withinLast24Hours(x) {
	var time = x.Interval;

	return moment(time).isAfter(moment().subtract('1', 'day'));
}

var goiban = {
	/*
	The MIT License (MIT)

	Copyright (c) 2014 Chris Grieger

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in
	all copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
	THE SOFTWARE.
	*/
	validate: function(iban, callback) {

		$.ajax({
			data: {"validateBankCode":true, "getBIC": true},
			url: '/validate/' + iban,
			success: function(data) {
				callback(data);
			},
			error: function(xhr) {
				callback("Empty request.");
			}});
	},

	getMetrics24h: function(data) {
		var chartData = {};

    data = _.chain(data)
      .filter(withinLast24Hours)
      .pluck('Counters')
      .reduce(function (acc, x) {
        var t = _.map(x, getCount);

        _.each(t, function (t) {
          acc[t.country] = (acc[t.country] || 0) + t.count;
        });

        return acc;
      }, {})
      .value();

    chartData.labels = _.keys(data);
    chartData.datasets = [{
      label: 'Sum',
			fillColor: "rgba(151,187,205,0.5)",
      strokeColor: "rgba(151,187,205,0.8)",
      highlightFill: "rgba(151,187,205,0.75)",
      highlightStroke: "rgba(151,187,205,1)",
      data: _.values(data)
    }];

    return chartData;
	}
};
