$(document).ready(function() {
	
	
	$('#btnworld').click(function() {
		$.ajax({
			type : "POST",
			async : false,
			url : "/world",
			success : function(result) {
				alert(result)
			}
		});
	});
	

})