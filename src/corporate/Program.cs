using CorporateManagement.API;
using CorporateManagement.API.Corporates.Endpoints;
using CorporateManagement.API.Infrastructure.Data;
using CorporateManagement.API.Infrastructure.Endpoints;
using FluentValidation;
using Identity.Grpc;

var builder = WebApplication.CreateBuilder(args);
builder.Configuration.AddEnvironmentVariables();

builder.Services.AddDb(builder.Configuration);

builder.Services.AddValidatorsFromAssembly(typeof(AssemblyMarker).Assembly);

builder.Services.AddGrpcClient<IdentityService.IdentityServiceClient>((sp, config) =>
{
  var identityUrl = sp.GetRequiredService<IConfiguration>().GetValue<string>("IdentityServiceGrpcUrl");
  if (string.IsNullOrEmpty(identityUrl))
    throw new Exception("Identity service url was not configured");
  config.Address = new Uri(identityUrl);
});

if (builder.Configuration.GetValue<bool>("EnableHealthChecks"))
  builder.Services.AddHealthChecks();

var app = builder.Build();


app.MapGroup("/corporate")
   .MapEndpoint<CreateCorporate>()
   .MapEndpoint<GetCorporates>()
   .MapGroup("/{corporateId}")
   .MapEndpoint<CreateEmployee>();

if (app.Configuration.GetValue<bool>("EnableHealthChecks"))
  app.MapHealthChecks("/hc");

app.SeedDb();
app.Run();


