# Asset Manager

Single Source of truth for all assets

Resources needed

To start all required resources 

```
cd localenv
./start.sh

```

-  Cache Database (Redis)
-  Data Source (Mongo Db)

Fetch Asset from external sources (Exchanges): to run service which fetch assets from external exchanges and save it to Data Source

```
 go run assetfetch.go
 
```

Feed Asset to cache which  can be used by the Scrapers to get the asset by name and symbol

```
go run assetfetch.go -feedRedis=true

```

