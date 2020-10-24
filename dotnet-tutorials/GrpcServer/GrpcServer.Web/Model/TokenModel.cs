using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace GrpcServer.Web.Model
{
    public class TokenModel
    {
        public string Token { get; set; }
        public DateTime? Expiration { get; set; }
        public bool isSuccess { get; set; }
    }
}
