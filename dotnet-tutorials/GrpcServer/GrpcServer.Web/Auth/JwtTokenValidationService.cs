using Google.Protobuf.WellKnownTypes;
using GrpcServer.Web.Model;
using Microsoft.IdentityModel.Tokens;
using System;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.Linq;
using System.Security.Claims;
using System.Text;
using System.Threading.Tasks;

namespace GrpcServer.Web.Auth
{
    public class JwtTokenValidationService
    {
        public TokenModel GenerateToken(UserModel user)
        {
            JwtSecurityToken token = null;
            if (user.Username == "admin" && user.Password == "abc123_")
            {
                var claims = new[]
                {
                    new Claim(JwtRegisteredClaimNames.Sub, "email@163.com"),
                    new Claim(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString()),
                    new Claim(JwtRegisteredClaimNames.UniqueName, "admin")
                };

                var key = new SymmetricSecurityKey(Encoding.UTF8.GetBytes("a64f60a2-62d7-11ea-974c-00163e0a1128"));
                var credentials = new SigningCredentials(key, SecurityAlgorithms.HmacSha256);
                token = new JwtSecurityToken(
                    issuer: "localhost",
                    audience: "localhost",
                    claims: claims,
                    expires: DateTime.UtcNow.AddMinutes(10),
                    signingCredentials: credentials);
            }

            return new TokenModel
            {
                Token = new JwtSecurityTokenHandler().WriteToken(token),
                Expiration = token?.ValidTo,
                isSuccess = token == null ? false : true
            };
        }

    }
}
