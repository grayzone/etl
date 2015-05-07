function get_gchart_data(data) {
    var g_data = new google.visualization.DataTable();
    g_data.addColumn('string', 'Country');
    g_data.addColumn('number', 'Popularity');
    g_data.addRows(data);
    return g_data;
  }	


// Callback that creates and populates a data table, 
// instantiates the pie chart, passes in the data and
// draws it.
function drawChart() {

// Create the data table.
var data = new google.visualization.DataTable();
data.addColumn('string', 'Topping');
data.addColumn('number', 'Slices');
data.addRows([
  ['Mushrooms', 3],
  ['Onions', 1],
  ['Olives', 1],
  ['Zucchini', 1],
  ['Pepperoni', 2]
]);

// Set chart options
var options = {'title':'How Much Pizza I Ate Last Night',
               'width':400,
               'height':300};

// Instantiate and draw our chart, passing in some options.
var chart = new google.visualization.PieChart(document.getElementById('chart_div'));
chart.draw(data, options);
}



function drawRegionsMap() {
	
	var data =  [['AE',214],
	             ['AG',2],
	             ['AL',8],
	             ['AN',4],
	             ['AR',5304],
	             ['AT',107],
	             ['AU',44204],
	             ['AW',6],
	             ['BA',8],
	             ['BD',46],
	             ['BE',367491],
	             ['BG',38],
	             ['BH',11],
	             ['BM',154],
	             ['BN',3],
	             ['BO',160],
	             ['BR',52474],
	             ['BS',227],
	             ['BZ',1],
	             ['CA',862],
	             ['CH',3556],
	             ['CL',6392],
	             ['CN',7514],
	             ['CO',7543],
	             ['CR',890],
	             ['CS',3],
	             ['CY',4],
	             ['CZ',41],
	             ['DE',63608],
	             ['DK',13183],
	             ['DM',7],
	             ['DO',32],
	             ['DZ',20],
	             ['EC',604],
	             ['EE',8],
	             ['EG',85],
	             ['ES',4001],
	             ['FI',22702],
	             ['FR',31986],
	             ['GB',2530],
	             ['GR',524],
	             ['GT',361],
	             ['GU',59],
	             ['HK',10762],
	             ['HN',32],
	             ['HR',30],
	             ['HT',3],
	             ['HU',220],
	             ['ID',59],
	             ['IE',29701],
	             ['IL',8903],
	             ['IN',41278],
	             ['IR',118],
	             ['IS',19],
	             ['IT',4651],
	             ['JM',288],
	             ['JO',49],
	             ['JP',253295],
	             ['KH',206],
	             ['KP',442],
	             ['KR',111132],
	             ['KW',261],
	             ['KY',2],
	             ['KZ',4],
	             ['LC',1],
	             ['LI',4],
	             ['LK',93],
	             ['LT',8],
	             ['LU',2],
	             ['LV',6],
	             ['MA',8],
	             ['MP',2],
	             ['MT',10],
	             ['MX',10003],
	             ['MY',553],
	             ['NC',31],
	             ['NG',5],
	             ['NI',164],
	             ['NL',7847],
	             ['NO',213],
	             ['NP',25],
	             ['NZ',482],
	             ['OM',13],
	             ['PA',23789],
	             ['PE',2064],
	             ['PF',147],
	             ['PH',1049],
	             ['PK',126],
	             ['PL',922],
	             ['PR',16365],
	             ['PT',621],
	             ['PY',145],
	             ['QA',18],
	             ['RO',14],
	             ['RU',417],
	             ['SA',1787],
	             ['SE',749],
	             ['SG',167757],
	             ['SI',188],
	             ['SK',1],
	             ['SO',2],
	             ['SV',100],
	             ['TH',3353],
	             ['TN',4],
	             ['TR',880],
	             ['TT',27],
	             ['TW',4154],
	             ['UA',76],
	             ['US',3747865],
	             ['UY',13138],
	             ['VE',221],
	             ['VG',1],
	             ['VI',9],
	             ['VN',23],
	             ['ZA',3881]];

		
		var g_data = get_gchart_data(data);
     
		var options = {
				 region: '019'
				
		};

     var chart = new google.visualization.GeoChart(document.getElementById('chart_div'));

     chart.draw(g_data, options);

	$.ajax({
		type : "POST",
		url : "/world",
		dataType:"json",
		async: false ,
		success : function(result) {
			

	        
			
			d = eval(result)
	//		alert(d)
		}
	});
	      }