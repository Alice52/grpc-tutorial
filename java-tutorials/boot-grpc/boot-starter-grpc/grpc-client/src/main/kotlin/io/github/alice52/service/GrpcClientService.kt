package io.github.alice52.service

import io.github.alice52.grpc.HelloReply
import io.github.alice52.grpc.HelloRequest
import io.github.alice52.grpc.SimpleGrpc
import io.grpc.StatusRuntimeException
import net.devh.boot.grpc.client.inject.GrpcClient
import org.springframework.stereotype.Service

@Service
class GrpcClientService {

    @GrpcClient("local-grpc-server")
    private lateinit var simpleStub: SimpleGrpc.SimpleBlockingStub

    fun sendMessage(name: String?): String? {
        return try {
            val response: HelloReply = simpleStub.sayHello(HelloRequest.newBuilder().setName(name).build())
            response.getMessage()
        } catch (e: StatusRuntimeException) {
            "FAILED with " + e.status.code.name
        }
    }
}
