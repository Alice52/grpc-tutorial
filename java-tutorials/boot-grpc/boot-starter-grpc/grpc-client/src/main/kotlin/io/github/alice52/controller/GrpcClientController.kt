package io.github.alice52.controller

import io.github.alice52.service.GrpcClientService
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController
import javax.annotation.Resource

@RestController
@RequestMapping
class GrpcClientController {
    @Resource
    private lateinit var grpcClientService: GrpcClientService

    @GetMapping("/hello")
    fun printMessage(@RequestParam(defaultValue = "Michael") name: String?): String? {
        return grpcClientService.sendMessage(name)
    }
}
