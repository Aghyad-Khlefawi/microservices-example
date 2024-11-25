using CorporateManagement.API.Corporates.Entities;
using MongoDB.Bson;
using MongoDB.Driver;

namespace CorporateManagement.API;

public class DbSeeder
{
  public static void SeedDb(IMongoClient client)
  {
    var collection = client.GetDatabase(DatabaseName.Corporate).GetCollection<Corporate>(Collections.Coporates);
    var count = collection.CountDocuments(new BsonDocument());
    if (count > 0)
      return;

    collection.InsertOne(Corporate.Create("Demo corp."));
  }
}
