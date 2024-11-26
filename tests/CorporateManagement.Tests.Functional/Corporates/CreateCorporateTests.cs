using System.Net.Http.Json;
using CorporateManagement.API.Corporates.Endpoints;
using Newtonsoft.Json;

namespace CorporateManagement.Tests.Functional.Corporates;

public class CreateCorporateTests(ApiFactory apiFactory) : IClassFixture<ApiFactory>, IAsyncLifetime
{
    private readonly HttpClient _client = apiFactory.CreateClient();

    [Fact]
    public async Task CreateCorporate_ShouldReturn200WithCreateCorporateResponse_WhenEntryIsValid()
    {
        // Arrange
        var request = new CreateCorporate.CreateCorporateRequest("Test");
        
        // Act
        var response = await _client.PostAsJsonAsync("/corporate",request);

        // Assert
        response.Should().BeSuccessful();
        
        string responseContent = await response.Content.ReadAsStringAsync();
        var responseObject = JsonConvert.DeserializeObject<CreateCorporate.CreateCorporateResponse>(responseContent);
        responseObject.Should().NotBeNull();
        responseObject!.Id.Should().NotBeNullOrEmpty();
    }
    [Fact]
    public async Task CreateCorporate_ShouldReturn400_WhenCorporateNameIsNotProvided()
    {
        // Arrange
        var request = new CreateCorporate.CreateCorporateRequest("");
        
        // Act
        var response = await _client.PostAsJsonAsync("/corporate",request);

        // Assert
        response.Should().HaveClientError();
    }
    public Task InitializeAsync() => Task.CompletedTask;

    public async Task DisposeAsync()
    {
        await apiFactory.ResetDatabase();
    }
}