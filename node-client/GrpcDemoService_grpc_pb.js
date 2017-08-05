// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var GrpcDemoService_pb = require('./GrpcDemoService_pb.js');

function serialize_GrpcDemoService_HelloRequest(arg) {
  if (!(arg instanceof GrpcDemoService_pb.HelloRequest)) {
    throw new Error('Expected argument of type GrpcDemoService.HelloRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_GrpcDemoService_HelloRequest(buffer_arg) {
  return GrpcDemoService_pb.HelloRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_GrpcDemoService_HelloResponse(arg) {
  if (!(arg instanceof GrpcDemoService_pb.HelloResponse)) {
    throw new Error('Expected argument of type GrpcDemoService.HelloResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_GrpcDemoService_HelloResponse(buffer_arg) {
  return GrpcDemoService_pb.HelloResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var GrpcDemoServiceService = exports.GrpcDemoServiceService = {
  hello: {
    path: '/GrpcDemoService.GrpcDemoService/Hello',
    requestStream: false,
    responseStream: false,
    requestType: GrpcDemoService_pb.HelloRequest,
    responseType: GrpcDemoService_pb.HelloResponse,
    requestSerialize: serialize_GrpcDemoService_HelloRequest,
    requestDeserialize: deserialize_GrpcDemoService_HelloRequest,
    responseSerialize: serialize_GrpcDemoService_HelloResponse,
    responseDeserialize: deserialize_GrpcDemoService_HelloResponse,
  },
};

exports.GrpcDemoServiceClient = grpc.makeGenericClientConstructor(GrpcDemoServiceService);
