// This Demo section deals with capturing the server time and creating a clientSide time. When we subtract the two we get the difference
// Need to be tested, but general idea works in jConsole...
var hmsServerTimer = '16:23:33'; // Server Timer
var clientTime = '16:45:00'; // Client Timer

// split both times at the colons
var a = hmsServerTimer.split(':'); 
var b = clientTime.split(':'); 

// minutes are worth 60 seconds. Hours are worth 60 minutes.
var timerOne = (+a[0]) * 60 * 60 + (+a[1]) * 60 + (+a[2]); 
var timerTwo = (+b[0]) * 60 * 60 + (+b[1]) * 60 + (+b[2]); 

console.log(timerOne);
console.log(timerTwo);

var diff = timerTwo - timerOne;

console.log("The difference is: " + diff);