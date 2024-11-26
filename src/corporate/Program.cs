using CorporateManagement.API;
using CorporateManagement.API.Corporates.Endpoints;
using CorporateManagement.API.Infrastructure.Data;
using CorporateManagement.API.Infrastructure.Endpoints;
using FluentValidation;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddDb(builder.Configuration);

builder.Services.AddValidatorsFromAssembly(typeof(AssemblyMarker).Assembly);

var app = builder.Build();


app.MapGroup("/corporate")
   .MapEndpoint<CreateCorporate>()
   .MapEndpoint<GetCorporates>()
  
   .MapGroup("/{corporateId}")
   .MapEndpoint<CreateEmployee>();


app.SeedDb();
app.Run();


