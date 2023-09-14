package io.github.alice52

import common.swagger.annotation.EnableSwagger
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@EnableSwagger
@SpringBootApplication
class GrpcClientApplication

fun main(args: Array<String>) {
    runApplication<GrpcClientApplication>(*args)
}