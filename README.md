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
