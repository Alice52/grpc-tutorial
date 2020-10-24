using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using GrpcServer.Web.Auth;
using GrpcServer.Web.Data;
using GrpcServer.Web.Model;
using GrpcServer.Web.Protos;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Authorization;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace GrpcServer.Web.Service
{
    [Authorize(AuthenticationSchemes = JwtBearerDefaults.AuthenticationScheme)]
    public class MyEmployeeService : EmployeeService.EmployeeServiceBase
    {
        private readonly ILogger<MyEmployeeService> _logger;
        private readonly JwtTokenValidationService _tokenValidationService;

        public MyEmployeeService(ILogger<MyEmployeeService> logger, JwtTokenValidationService tokenValidationService)
        {
            _logger = logger;
            _tokenValidationService = tokenValidationService;
        }

        [AllowAnonymous]
        public override Task<TokenResponse> GenerateToken(TokenRequest request, ServerCallContext context)
        {
            TokenModel tokenModel = _tokenValidationService.GenerateToken(new UserModel
            {
                Username = request.Username,
                Password = request.Password
            });

            return Task.FromResult(new TokenResponse
            {
                Token = tokenModel.Token,
                Expiration = Timestamp.FromDateTime(tokenModel.Expiration ?? DateTime.MinValue),
                IsSuccess = tokenModel.isSuccess
            });
        }

        public override Task<EmployeeResponse> GetByNo(GetByNoRequest request, ServerCallContext context)
        {
            var md = context.RequestHeaders;

            foreach (var pair in md)
            {
                _logger.LogInformation($"{pair.Key}: {pair.Value}");
            }

            var employee = InMemoryData.Employees.SingleOrDefault(x => x.No == request.No);

            if (employee != null)
            {
                var response = new EmployeeResponse
                {
                    Employee = employee
                };

                return Task.FromResult(response);
            }

            throw new RpcException(new Status(StatusCode.DataLoss, "data loss"), $"Employee not found with no : {request.No}");
        }

        public override async Task GetByAll(GetByAllRequest request, IServerStreamWriter<EmployeeResponse> responseStream, ServerCallContext context)
        {
            foreach (var employee in InMemoryData.Employees)
            {
                await responseStream.WriteAsync(new EmployeeResponse
                {
                    Employee = employee
                });
            }
        }

        public override async Task<AddPhoneResponse> AddPhoto(IAsyncStreamReader<AddPhoneRequest> requestStream, ServerCallContext context)
        {
            var data = new List<byte>();
            while (await requestStream.MoveNext())
            {
                data.AddRange(requestStream.Current.Data);
            }

            Console.WriteLine($"Receive file with {data.Count} bytes");


            return new AddPhoneResponse
            {
                IsOk = true
            };
        }

        public override async Task SaveAll(IAsyncStreamReader<EmployeeRequest> requestStream, IServerStreamWriter<EmployeeResponse> responseStream, ServerCallContext context)
        {
            while (await requestStream.MoveNext())
            {
                InMemoryData.Employees.Add(requestStream.Current.Employee);
                await responseStream.WriteAsync(new EmployeeResponse
                {
                    Employee = requestStream.Current.Employee
                });
            }

            foreach (var e in InMemoryData.Employees)
            {
                Console.WriteLine(e);
            }
        }
    }
}
