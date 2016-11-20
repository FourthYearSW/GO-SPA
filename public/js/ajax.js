// To be documented...
$(document).ready(function(){
	$("button").click(function(){
		$.ajax({
			url: "/test", 
			method: "GET",
			success: function(result){
				$("#responseHead").html(result);
			}
		});
	});
});