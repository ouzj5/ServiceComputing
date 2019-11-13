$(document).ready(function() {
    $.ajax({
        url: "/getHeader"
    }).then(function(data) {
    	alert("simple js acess: " + data);
       $('#head').html(data);
    });
});

