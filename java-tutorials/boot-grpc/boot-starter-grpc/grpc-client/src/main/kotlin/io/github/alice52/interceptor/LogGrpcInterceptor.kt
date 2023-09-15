package io.github.alice52.interceptor

import io.github.alice52.logger
import io.grpc.*

class LogGrpcInterceptor : ClientInterceptor {
    override fun <ReqT, RespT> interceptCall(
        method: MethodDescriptor<ReqT, RespT>, callOptions: CallOptions, next: Channel
    ): ClientCall<ReqT, RespT> {
        logger().info(method.fullMethodName)
        return next.newCall(method, callOptions)
    }
}
