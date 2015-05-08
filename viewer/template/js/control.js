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
				"province":$('#province').val()				
			},
			async: false ,
			success : function(result) {
				if (result == null){
					return ;
				}
				var data = [];
				for (i=0;i<result.length;i++){
					var item = [result[i].City, result[i].Total];
					data.push(item);
				}
				var options = {
						region : $('#province').val(),
						resolution:'provinces',
						displayMode:'markers'

					};	
				
				charts($('#province').val(), data, options);

			}
		});
		
		
	})
		

	



	

	

})