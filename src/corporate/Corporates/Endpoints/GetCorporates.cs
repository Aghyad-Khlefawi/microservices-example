
using CorporateManagement.API.Corporates.Entities;
using CorporateManagement.API.Infrastructure.Data;
using CorporateManagement.API.Infrastructure.Endpoints;
using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;
using MongoDB.Driver;

namespace CorporateManagement.API.Corporates.Endpoints;

public class GetCorporates: IEndpoint {
  public static void Map(IEndpointRouteBuilder app) => app.Map("/", async (IMongoClient mongoClient)=>
      TypedResults.Ok(await mongoClient.GetCorporatesCollectionAs<GetCorporatesResponse>().Find(new BsonDocument()).ToListAsync()) 
      ); 


  public record GetCorporatesResponse([property:BsonId] [property:BsonRepresentation(BsonType.ObjectId)] string Id, string Name, List<Employee> Employees);
}
