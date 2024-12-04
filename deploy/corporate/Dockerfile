FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
WORKDIR /app
EXPOSE 8080
EXPOSE 8081
RUN apt-get update && apt-get install -y curl
USER app

FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
ARG BUILD_CONFIGURATION=Release
WORKDIR /src/corporate
COPY ["src/corporate/CorporateManagement.API.csproj", "."]
RUN dotnet restore "CorporateManagement.API.csproj"

COPY ["src/corporate", "."]

COPY ["src/identity/protos/identity.proto", "/src/identity/protos/identity.proto"]

RUN dotnet build "./CorporateManagement.API.csproj" -c $BUILD_CONFIGURATION -o /app/build

FROM build AS publish
ARG BUILD_CONFIGURATION=Release
RUN dotnet publish "./CorporateManagement.API.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD curl -f http://localhost:8080/hc || exit 1 

ENTRYPOINT ["dotnet", "CorporateManagement.API.dll"]  

