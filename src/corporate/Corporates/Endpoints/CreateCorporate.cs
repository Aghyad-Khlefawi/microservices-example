using CorporateManagement.API.Corporates.Entities;
using CorporateManagement.API.Infrastructure;
using CorporateManagement.API.Infrastructure.Data;
using CorporateManagement.API.Infrastructure.Endpoints;
using FluentValidation;
using MongoDB.Driver;

namespace CorporateManagement.API.Corporates.Endpoints;

public class CreateCorporate:IEndpoint
{
  public static void Map(IEndpointRouteBuilder app)
    => app.MapPost("/", async (CreateCorporateRequest request, IMongoClient mongoClient) =>
    {
      var corp = Corporate.Create(request.Name);
      await mongoClient.GetCorporatesCollection().InsertOneAsync(corp);
      return TypedResults.Ok(new CreateCorporateResponse(corp.Id));
    }).WithRequestValidation<Validator>();

  public record CreateCorporateResponse(string Id);
  public record CreateCorporateRequest(string Name);


  public class Validator : AbstractValidator<CreateCorporateRequest>
  {
    public Validator()
    {
      RuleFor(e => e.Name).NotEmpty().MaximumLength(100);
    }
  }
}
