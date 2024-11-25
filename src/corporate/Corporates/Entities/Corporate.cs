using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace CorporateManagement.API.Corporates.Entities;

public class Corporate
{
    private Corporate()
    {
    }
    [BsonId]
    [BsonRepresentation(BsonType.ObjectId)]
    public required string Id { get; init; }
    public required string Name { get; init; }
    public List<Employee> Employees { get; init; } = [];

    public static Corporate Create(string name) =>
        new()
        {
            Id = ObjectId.GenerateNewId().ToString(),
            Name = name
        };
}
