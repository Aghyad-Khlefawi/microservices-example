namespace CorporateManagement.API.Infrastructure.Endpoints;

public interface IEndpoint{
  public abstract static void Map(IEndpointRouteBuilder app);
}
