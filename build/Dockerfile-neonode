FROM microsoft/dotnet:sdk AS build-env
WORKDIR /app

RUN git clone https://github.com/neo-project/neo-cli.git

WORKDIR /app/neo-cli/neo-cli

RUN dotnet restore && dotnet publish -c Release -o out

# Build runtime image
FROM microsoft/dotnet:aspnetcore-runtime
WORKDIR /app

RUN apt-get update && apt-get install wget libleveldb-dev sqlite3 libsqlite3-dev libunwind8-dev -y

COPY --from=build-env /app/neo-cli/neo-cli/out .
ENTRYPOINT ["dotnet", "neo-cli.dll", "/rpc"]
