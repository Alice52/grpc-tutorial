package io.github.alice52.interceptor

import io.github.alice52.logger
import io.grpc.Metadata
import io.grpc.ServerCall
import io.grpc.ServerCallHandler
import io.grpc.ServerInterceptor

class LogGrpcInterceptor : ServerInterceptor {
    override fun <ReqT, RespT> interceptCall(
        serverCall: ServerCall<ReqT, RespT>,
        metadata: Metadata,
        serverCallHandler: ServerCallHandler<ReqT, RespT>
    ): ServerCall.Listener<ReqT> {
        logger().info(serverCall.methodDescriptor.fullMethodName)
        return serverCallHandler.startCall(serverCall, metadata)
    }
}
