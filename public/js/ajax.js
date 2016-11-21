
/* Make an ajax call to get the timer value at a point 
   in time from the server and set the value to an html element */
$(document).ready(function(){
	$("button").click(function(){
		$.ajax({
			url: "/test", 
			method: "GET",
			success: function(result){
				$("#responseHead").html(result);
			}
		});
		var getTimerVal = $("#responseHead").html();
	});
});

// console.log(getTimerVal);

// Capture what's been sent by the server and store in a javascript variable
function getTimervalue() {
	var getTimerVal = $("#responseHead").html();
}