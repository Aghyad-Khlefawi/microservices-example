using CorporateManagement.API.Corporates.Entities;
using CorporateManagement.API.Infrastructure;
using CorporateManagement.API.Infrastructure.Data;
using CorporateManagement.API.Infrastructure.Endpoints;
using FluentValidation;
using MongoDB.Driver;

namespace CorporateManagement.API.Corporates.Endpoints;

public class CreateEmployee : IEndpoint
{
  public static void Map(IEndpointRouteBuilder app)
    => app.MapPost("/employee", async (CreateEmployeeRequest request, string corporateId, IMongoClient mongoClient) =>
    {
      var employee = Employee.Create(request.FirstName, request.LastName, request.Email);

      var collection = mongoClient.GetCorporatesCollection();
      if (collection.CountDocuments(e => e.Id == corporateId) < 1)
        throw new ValidationException("Invalid corporate id");

      await collection.UpdateOneAsync(e => e.Id == corporateId, Builders<Corporate>.Update.Push(e => e.Employees, employee));
      return TypedResults.Ok(new CreateEmployeeResponse(employee.Id));
    }).WithRequestValidation<CreateEmployeeRequest>();

  public record CreateEmployeeResponse(string Id);
  public record CreateEmployeeRequest(string FirstName, string LastName, string Email, string CorporateId);


  public class Validator : AbstractValidator<CreateEmployeeRequest>
  {
    public Validator(IMongoClient mongoClient)
    {
      RuleFor(e => e.FirstName).NotEmpty().MaximumLength(100);
      RuleFor(e => e.LastName).NotEmpty().MaximumLength(100);
      RuleFor(e => e.Email)
        .EmailAddress()
        .MustAsync(async (request, email, ct) =>
          !(await mongoClient.GetCorporatesCollection()
                             .Aggregate()
                             .Match(Builders<Corporate>.Filter.ElemMatch(e => e.Employees, e => e.Email == email))
                             .AnyAsync(ct))
          ).WithMessage("Another employee with the same email address already exists");
    }
  }
}
