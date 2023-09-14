package io.github.alice52.configuration

import io.github.alice52.interceptor.LogGrpcInterceptor
import net.devh.boot.grpc.server.interceptor.GrpcGlobalServerInterceptor
import org.springframework.context.annotation.Configuration

@Configuration(proxyBeanMethods = false)
class ServerConfiguration {

    @GrpcGlobalServerInterceptor
    fun logServerInterceptor(): LogGrpcInterceptor {
        return LogGrpcInterceptor()
    }
}
