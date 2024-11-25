using CorporateManagement.API.Corporates.Entities;
using MongoDB.Driver;

namespace CorporateManagement.API.Infrastructure.Data;

public static class Extensions
{
  public static IMongoCollection<T> GetCorporatesCollectionAs<T>(this IMongoClient mongoClient) =>
    mongoClient.GetDatabase(DatabaseName.Corporate).GetCollection<T>(Collections.Coporates);
  public static IMongoCollection<Corporate> GetCorporatesCollection(this IMongoClient mongoClient) =>
    mongoClient.GetCorporatesCollectionAs<Corporate>();


  public static void AddDb(this IServiceCollection services, IConfiguration configurations)
  {
    services.AddSingleton<IMongoClient>((sp) =>
    {
      var connection = configurations.GetConnectionString("CorporateDatabase");
      var client = new MongoClient(connection);
      return client;
    });

    DbSeeder.SeedDb(services.BuildServiceProvider().GetRequiredService<IMongoClient>());


  }
}
