using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;
 
namespace CorporateManagement.API.Corporates.Entities;

public class Employee
{
    private Employee()
    {
        
    }
    
    [BsonRepresentation(BsonType.ObjectId)]
    public required string Id { get; init; }
    public required string FirstName { get; init; }
    public required string LastName { get; init; }
    public required string Email { get; set; }

    public static Employee Create(string firstName, string lastName,string email) =>
        new()
        {
            Id = ObjectId.GenerateNewId().ToString(),
            FirstName = firstName,
            LastName = lastName,
            Email=email
        };
}
