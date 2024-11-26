using CorporateManagement.API;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.AspNetCore.TestHost;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.DependencyInjection.Extensions;
using MongoDB.Driver;
using Testcontainers.MongoDb;

namespace CorporateManagement.Tests.Functional;

public class ApiFactory : WebApplicationFactory<AssemblyMarker>, IAsyncLifetime
{
    private readonly MongoDbContainer _dbContainer = new MongoDbBuilder().Build();

    protected override void ConfigureWebHost(IWebHostBuilder builder)
    {
        builder.ConfigureTestServices(services =>
        {
            services.RemoveAll(typeof(IMongoClient));
            services.AddSingleton<IMongoClient>((sp) =>
            {
                var client = new MongoClient(_dbContainer.GetConnectionString());
                return client;
            });
        });
    }

    public async Task ResetDatabase()
    {
        using var scope = Services.CreateScope();
        var mongoClient = scope.ServiceProvider.GetRequiredService<IMongoClient>();
        var db = mongoClient.GetDatabase(DatabaseName.Corporate);
        var collections = await (await db.ListCollectionNamesAsync()).ToListAsync();
        foreach (var collection in collections)
        {
            await db.DropCollectionAsync(collection);
        }
    }

    public async Task InitializeAsync()
    {
        await _dbContainer.StartAsync();
    }

    public new async Task DisposeAsync()
    {
        await _dbContainer.DisposeAsync();
    }
}