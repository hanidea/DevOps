const http = require('http');
const hostname = '0.0.0.0';
const port = 3000;
var express = require('express');
var fs=require("fs");
var app = express();
app.get("/main.html",function(request,response){ 
    fs.readFile("./"+request.path.substr(1), function(err, data) {
            if (err) {
                console.log(err);
            }
            else{
		response.writeHead(200,{"Content-Type":"text/html"});
	        response.write(data.toString());
	    }
       	    response.end();
        });
}); 

app.listen(port, hostname, () => { 
    console.log(`Server running at http://${hostname}:${port}/`);
});
