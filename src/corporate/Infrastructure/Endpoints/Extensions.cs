namespace CorporateManagement.API.Infrastructure.Endpoints;


public static class Extensions
{
  public static IEndpointRouteBuilder MapEndpoint<T>(this IEndpointRouteBuilder app) where T : IEndpoint
  {
    T.Map(app);
    return app;
  }
}
