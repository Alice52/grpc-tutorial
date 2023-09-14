package io.github.alice52.service

import io.github.alice52.grpc.HelloReply
import io.github.alice52.grpc.HelloRequest
import io.github.alice52.grpc.SimpleGrpc
import io.grpc.stub.StreamObserver
import net.devh.boot.grpc.server.service.GrpcService

@GrpcService
class GrpcServerService : SimpleGrpc.SimpleImplBase() {

    override fun sayHello(req: HelloRequest, responseObserver: StreamObserver<HelloReply?>) {
        val reply = HelloReply.newBuilder().setMessage("Hello ==> " + req.name).build()
        responseObserver.onNext(reply)
        responseObserver.onCompleted()
    }
}
