package io.github.alice52.configuration

import io.github.alice52.interceptor.LogGrpcInterceptor
import net.devh.boot.grpc.client.interceptor.GrpcGlobalClientInterceptor
import org.springframework.context.annotation.Configuration
import org.springframework.core.Ordered
import org.springframework.core.annotation.Order

@Order(Ordered.LOWEST_PRECEDENCE)
@Configuration(proxyBeanMethods = false)
class ClientConfiguration {

    @GrpcGlobalClientInterceptor
    fun logClientInterceptor(): LogGrpcInterceptor {

        return LogGrpcInterceptor()
    }
}
