const fs = require("fs");
const grpc = require("grpc");
const path = require("path");

const messages = require("./GrpcDemoService_pb");
const services = require("./GrpcDemoService_grpc_pb");


let credentials = grpc.credentials.createSsl(
	fs.readFileSync(path.join(__dirname, "./ca.crt")), 
	fs.readFileSync(path.join(__dirname, "./client.key")), 
	fs.readFileSync(path.join(__dirname, "./client.crt"))
);

var client = new services.GrpcDemoServiceClient("localhost:7777", credentials);

var req = new messages.HelloRequest();
req.setName("jaska");

client.hello(req, (err, res) => {
	if (err) console.log("error happended ", err);
	else {
		console.log("returnString", res.getReturnstring());
	}
});