$(document).ready(function() {


	
	$('#btnworld').click(function() {
		
		$.ajax({
			type : "POST",
			url : "/world",
			dataType:"json",
			async: false ,
			success : function(result) {
				var data = [];
				for (i=0;i<result.length;i++){
					var item = [result[i].Country, result[i].Total];
					data.push(item);
				}
							
				d = GetDatabyContinent(data);
				

				var options = {
					region : $('#continent').val()

				};
				
				charts($('#continent').val(), d, options);

			}
		});
		
		
	})
			    
	$('#btncountry').click(function() {
		
		$.ajax({
			type : "POST",
			url : "/us",
			dataType:"json",
			async: false ,
			data:{
				"devicetype":$('#devicetype').val()				
			},
			success : function(result) {
				var data = [];
				for (i=0;i<result.length;i++){
					var item = [result[i].Province, result[i].Total];
					data.push(item);
				}
				var options = {
						region : $('#country').val(),
						resolution:'provinces'

					};	
				
				charts($('#country').val(), data, options);

			}
		});
		
		
	})
	
		$('#btnprovince').click(function() {
	
		$.ajax({
			type : "POST",
			url : "/province",
			dataType:"json",
			data:{
				"province":$('#province').val()	,
				"devicetype":$('#devicetype').val()
			},
			async: false ,
			success : function(result) {
				if (result == null){
					alert("no results found");
					return ;
				}
				var data = [];
				for (i=0;i<result.length;i++){
					var coordinate = result[i].Coordinate;
					var len = coordinate.length
					if(len > 0){
						var index = coordinate.indexOf(",");
						var lat = parseFloat(coordinate.substring(1,index));
						var lng = parseFloat(coordinate.substring(index+1,len-1));
						
						var item = [lat,lng,result[i].City, result[i].Total];
						data.push(item);						
					}
					
				}
				var options = {
						region : $('#province').val(),
						resolution:'provinces',
						displayMode:'markers'

					};	
				if ($('#province').val() == "US-ALL"){
					options["region"] = "US";
				}

				charts_coordinate($('#province').val(), data, options);

			}
		});
		
		
	})
		

	



	

	

})