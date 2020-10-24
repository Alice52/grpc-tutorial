using Google.Protobuf.WellKnownTypes;
using GrpcServer.Web.Protos;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace GrpcServer.Web.Data
{
    public class InMemoryData
    {
        public static List<Employee> Employees = new List<Employee>
        {
            new Employee
            {
                Id = 1,
                No = 1994,
                FirstName = "zack",
                LastName = "zhang",
                //Salary = 100,
                MonthSalary = new MonthSalary
                {
                    Basic=15,
                    Bonus=20
                },
                Status =  EmployeeStatus.Normal,
                LastModified = Timestamp.FromDateTime(DateTime.UtcNow)
            },
            new Employee
            {
                Id = 2,
                No = 1995,
                FirstName = "tim",
                LastName = "tang",
                //Salary = 500,
                MonthSalary = new MonthSalary
                {
                    Basic=115,
                    Bonus=210
                },
                Status =  EmployeeStatus.Normal,
                LastModified = Timestamp.FromDateTime(DateTime.UtcNow)
            },
            new Employee
            {
                Id = 3,
                No = 1999,
                FirstName = "kayla",
                LastName = "qi",
                //Salary = 8100,
                         MonthSalary = new MonthSalary
                {
                    Basic=315,
                    Bonus=310
                },
                Status =  EmployeeStatus.Normal,
                LastModified = Timestamp.FromDateTime(DateTime.UtcNow)
            },
            new Employee
            {
                Id = 4,
                No = 1991,
                FirstName = "able",
                LastName = "yan",
                //Salary = 1050,        
                MonthSalary = new MonthSalary
                {
                    Basic=415,
                    Bonus=410
                },
                Status =  EmployeeStatus.Normal,
                LastModified = Timestamp.FromDateTime(DateTime.UtcNow)
            },
            new Employee
            {
                Id = 5,
                No = 1990,
                FirstName = "klein",
                LastName = "zhou",
                //Salary = 100,
                 MonthSalary = new MonthSalary
                {
                    Basic=515,
                    Bonus=510
                },
                Status =  EmployeeStatus.Normal,
                LastModified = Timestamp.FromDateTime(DateTime.UtcNow)
            }
        };
    }
}
