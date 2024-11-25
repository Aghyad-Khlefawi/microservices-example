using CorporateManagement.API.Infrastructure.Filters;

namespace CorporateManagement.API.Infrastructure;

public static class Extensions
{

  public static RouteHandlerBuilder WithRequestValidation<TRequest>(this RouteHandlerBuilder builder)
    => builder.AddEndpointFilter<RequestValidationFilter<TRequest>>()
      .ProducesValidationProblem();

}
