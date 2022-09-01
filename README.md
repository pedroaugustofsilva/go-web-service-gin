# Go Web Service with Gin

## Listar todos os álbums

```sh
curl http://localhost:3000/albums
```

## Adicionar um novo item

```sh
curl http://localhost:3000/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

## Buscar um item por id

```sh
curl http://localhost:3000/albums/2
```

## Rodar o container do Postgres

```sh
docker run --name postgres -e POSTGRES_PASSWORD=postgres -d postgres
```

```sh
docker cp ./albums.sql postgres:/
```

```sh
docker exec postgres psql -h localhost -p 5432 -U postgres -f /albums.sql
```
