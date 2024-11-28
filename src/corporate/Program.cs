using CorporateManagement.API;
using CorporateManagement.API.Corporates.Endpoints;
using CorporateManagement.API.Infrastructure.Data;
using CorporateManagement.API.Infrastructure.Endpoints;
using FluentValidation;
using Identity.Grpc;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddDb(builder.Configuration);

builder.Services.AddValidatorsFromAssembly(typeof(AssemblyMarker).Assembly);

builder.Services.AddGrpcClient<IdentityService.IdentityServiceClient>((config) =>
{
    config.Address = new Uri("http://localhost:5001");
});

var app = builder.Build();


app.MapGroup("/corporate")
   .MapEndpoint<CreateCorporate>()
   .MapEndpoint<GetCorporates>()
  
   .MapGroup("/{corporateId}")
   .MapEndpoint<CreateEmployee>();


app.SeedDb();
app.Run();


