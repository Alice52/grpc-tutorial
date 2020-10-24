using Google.Protobuf;
using Grpc.Core;
using Grpc.Net.Client;
using GrpcServer.Web.Protos;
using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using static GrpcServer.Web.Protos.EmployeeService;

namespace GrpcClient
{
    class Program
    {
        private static DateTime _expiration = DateTime.MinValue;
        private static string _token;

        private static bool NeedGetToken() => string.IsNullOrEmpty(_token) || _expiration > DateTime.UtcNow;

        static async Task Main(string[] args)
        {
            var userInfo = new Metadata
            {
                {"username", "admin" },
                {"password", "abc123_" }
            };

            using var channel = GrpcChannel.ForAddress("https://localhost:5001");
            var client = new EmployeeServiceClient(channel);
            var option = int.Parse(args[0]);

            if (!NeedGetToken() || GetToken(client, userInfo))
            {

                var headers = new Metadata
                {
                    {"Authorization", $"Bearer {_token}" }
                };

                switch (option)
                {
                    case 1:
                        await GetByNoAsync(client, headers);
                        break;

                    case 2:
                        await GetAllAsync(client, headers);
                        break;

                    case 3:
                        await AddPhotoAsync(client, headers);
                        break;
                    case 4:
                        await SaveAllAsync(client, headers);
                        break;

                    default:
                        break;
                }

            }


            Console.WriteLine("Press any key to exit.");
            Console.ReadKey();
        }

        private static bool GetToken(EmployeeServiceClient client, Metadata entries)
        {
            TokenResponse tokenResponse = client.GenerateToken(new TokenRequest
            {
                Username = entries.GetValue("username"),
                Password = entries.GetValue("password")
            });

            _token = tokenResponse.Token;
            _expiration = tokenResponse.Expiration.ToDateTime();

            return tokenResponse.IsSuccess;
        }


        public static async Task<Employee> GetByNoAsync(EmployeeServiceClient client, Metadata md)
        {
            var response = await client.GetByNoAsync(new GetByNoRequest
            {
                No = 1994
            }, md);

            return response.Employee;
        }

        public static async Task<IList<Employee>> GetAllAsync(EmployeeServiceClient client, Metadata md)
        {
            using var call = client.GetByAll(new GetByAllRequest(), md);

            var responseStream = call.ResponseStream;
            List<Employee> response = new List<Employee>();

            while (await responseStream.MoveNext())
            {
                var employee = responseStream.Current.Employee;
                Console.WriteLine(employee);
                response.Add(employee);
            }

            return response;
        }

        public static async Task AddPhotoAsync(EmployeeServiceClient client, Metadata md)
        {
            FileStream fs = File.OpenRead("14894-8.jpg");

            using var call = client.AddPhoto(md);
            var stream = call.RequestStream;

            while (true)
            {
                byte[] buffer = new byte[1024];
                int numRead = await fs.ReadAsync(buffer, 0, buffer.Length);
                if (numRead == 0)
                {
                    break;
                }

                if (numRead < buffer.Length)
                {
                    Array.Resize(ref buffer, numRead);
                }

                await stream.WriteAsync(new AddPhoneRequest()
                {
                    Data = ByteString.CopyFrom(buffer)
                });

            }

            await stream.CompleteAsync();

            var response = await call.ResponseAsync;
            Console.WriteLine(response.IsOk);
        }

        public static async Task<IList<Employee>> SaveAllAsync(EmployeeServiceClient client, Metadata md)
        {
            List<Employee> employees = new List<Employee>
            {
                new Employee
                {
                    Id = 12,
                    No = 1994,
                    FirstName = "zack",
                    LastName = "zhang",
                    //Salary = 100,
                },
                new Employee
                {
                    Id = 22,
                    No = 1995,
                    FirstName = "tim",
                    LastName = "tang",
                    //Salary = 500,
                }
            };

            var call = client.SaveAll(md);

            var responseStream = call.ResponseStream;
            var list = new List<Employee>();
            var responseTask = Task.Run(async () =>
            {
                while (await responseStream.MoveNext())
                {
                    Console.WriteLine(responseStream.Current.Employee);
                    list.Add(responseStream.Current.Employee);
                }
            });


            var requestStream = call.RequestStream;
            foreach (var employee in employees)
            {
                await requestStream.WriteAsync(new EmployeeRequest
                {
                    Employee = employee
                });
            }

            await requestStream.CompleteAsync();
            await responseTask;

            return list;
        }
    }
}
