$(document).ready(function() {
	
	function charts(rows){
		google.load('visualization', '1.0', {'packages':['corechart'],callback:function(){
			
			// Create the data table.
	        var data = new google.visualization.DataTable();
	        data.addColumn('string', $('#continent').val());
	        data.addColumn('number', 'Total');
	        data.addRows(rows);
	        
	        var options = {
					 region: $('#continent').val()
					
			};

	        // Instantiate and draw our chart, passing in some options.
	        var chart1 = new google.visualization.GeoChart(document.getElementById('chart_geo'));
	        chart1.draw(data, options);
	        var chart2 = new google.visualization.PieChart(document.getElementById('chart_pie'));
	        chart2.draw(data);
			
			
			
			
		}});
		
	}
	
	function GetDatabyContinent(data){
		var continent = $('#continent').val()
		var countrylist;
		if (continent == '002'){
			countrylist = "DZ,EG,EH,LY,MA,SD,TN,BF,BJ,CI,CV,GH,GM,GN,GW,LR,ML,MR,NE,NG,SH,SL,SN,TG,AO,CD,ZR,CF,CG,CM,GA,GQ,ST,TD,BI,DJ,ER,ET,KE,KM,MG,MU,MW,MZ,RE,RW,SC,SO,TZ,UG,YT,ZM,ZW,BW,LS,NA,SZ,ZA";			
		}else if(continent == '150'){
			countrylist = "GG,JE,AX,DK,EE,FI,FO,GB,IE,IM,IS,LT,LV,NO,SE,SJ,AT,BE,CH,DE,DD,FR,FX,LI,LU,MC,NL,BG,BY,CZ,HU,MD,PL,RO,RU,SU,SK,UA,AD,AL,BA,ES,GI,GR,HR,IT,ME,MK,MT,CS,RS,PT,SI,SM,VA,YU";
		}else if(continent == '019'){
			countrylist = "BM,CA,GL,PM,US,AG,AI,AN,AW,BB,BL,BS,CU,DM,DO,GD,GP,HT,JM,KN,KY,LC,MF,MQ,MS,PR,TC,TT,VC,VG,VI,BZ,CR,GT,HN,MX,NI,PA,SV,AR,BO,BR,CL,CO,EC,FK,GF,GY,PE,PY,SR,UY,VE";
		}else if(continent == '142'){
			countrylist = "TM,TJ,KG,KZ,UZ,CN,HK,JP,KP,KR,MN,MO,TW,AF,BD,BT,IN,IR,LK,MV,NP,PK,BN,ID,KH,LA,MM,BU,MY,PH,SG,TH,TL,TP,VN,AE,AM,AZ,BH,CY,GE,IL,IQ,JO,KW,LB,OM,PS,QA,SA,NT,SY,TR,YE,YD";
		}else if(continent == '009'){
			countrylist = "AU,NF,NZ,FJ,NC,PG,SB,VU,FM,GU,KI,MH,MP,NR,PW,AS,CK,NU,PF,PN,TK,TO,TV,WF,WS";
		}else{
			return data;			
		}
		
		var i = 0;
		while (data[i]) {
			if (countrylist.search(data[i][0]) == -1){
				data.splice(i,1);
			}else{
				i++;
				}
		
				
		}
		return data;	
		
	}

	
	$('#btnworld').click(function() {
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
//		charts(data);
		
		d = GetDatabyContinent(data);
		charts(d);
		
	})
			    

		

	



	

	

})