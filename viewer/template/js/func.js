function charts(column,rows, options) {
	
	google.load('visualization', '1.0', {
		'packages' : [ 'corechart' ],
		callback : function() {

			// Create the data table.
			var data = new google.visualization.DataTable();
			data.addColumn('string', column);
			data.addColumn('number', 'Total');
			data.addRows(rows);
			
			$('#chart_geo').remove();
			$("#geo").append("<div id='chart_geo' style='width: 512px; height: 384px;'></div>")

			// Instantiate and draw our chart, passing in some options.
			var chart1 = new google.visualization.GeoChart(document
					.getElementById('chart_geo'));
			chart1.clearChart();
			chart1.draw(data, options);
			var chart2 = new google.visualization.PieChart(document
					.getElementById('chart_pie'));
		//	chart2.clearChart();
			chart2.draw(data);

		}
	});
}

function GetDatabyContinent(data) {
	var continent = $('#continent').val()
	var countrylist;
	if (continent == '002') {
		countrylist = "DZ,EG,EH,LY,MA,SD,TN,BF,BJ,CI,CV,GH,GM,GN,GW,LR,ML,MR,NE,NG,SH,SL,SN,TG,AO,CD,ZR,CF,CG,CM,GA,GQ,ST,TD,BI,DJ,ER,ET,KE,KM,MG,MU,MW,MZ,RE,RW,SC,SO,TZ,UG,YT,ZM,ZW,BW,LS,NA,SZ,ZA";
	} else if (continent == '150') {
		countrylist = "GG,JE,AX,DK,EE,FI,FO,GB,IE,IM,IS,LT,LV,NO,SE,SJ,AT,BE,CH,DE,DD,FR,FX,LI,LU,MC,NL,BG,BY,CZ,HU,MD,PL,RO,RU,SU,SK,UA,AD,AL,BA,ES,GI,GR,HR,IT,ME,MK,MT,CS,RS,PT,SI,SM,VA,YU";
	} else if (continent == '019') {
		countrylist = "BM,CA,GL,PM,US,AG,AI,AN,AW,BB,BL,BS,CU,DM,DO,GD,GP,HT,JM,KN,KY,LC,MF,MQ,MS,PR,TC,TT,VC,VG,VI,BZ,CR,GT,HN,MX,NI,PA,SV,AR,BO,BR,CL,CO,EC,FK,GF,GY,PE,PY,SR,UY,VE";
	} else if (continent == '142') {
		countrylist = "TM,TJ,KG,KZ,UZ,CN,HK,JP,KP,KR,MN,MO,TW,AF,BD,BT,IN,IR,LK,MV,NP,PK,BN,ID,KH,LA,MM,BU,MY,PH,SG,TH,TL,TP,VN,AE,AM,AZ,BH,CY,GE,IL,IQ,JO,KW,LB,OM,PS,QA,SA,NT,SY,TR,YE,YD";
	} else if (continent == '009') {
		countrylist = "AU,NF,NZ,FJ,NC,PG,SB,VU,FM,GU,KI,MH,MP,NR,PW,AS,CK,NU,PF,PN,TK,TO,TV,WF,WS";
	} else {
		return data;
	}

	var i = 0;
	while (data[i]) {
		if (countrylist.search(data[i][0]) == -1) {
			data.splice(i, 1);
		} else {
			i++;
		}

	}
	return data;

}