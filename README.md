# go-expert-lab-o11y-otel
Pós Go Expert Lab Otel

### Buildar a imagem docker e inicar os serviços
```bash
    make docker-compose-all-up
```

### Parar os serviços
```bash
    make docker-compose-all-down
```

### Acessar o Zipkin:
```bash
    http://localhost:9411
```

### Testar

## 1 - Navegue até a pasta api no diretório zipcodeservice:
```bash
    cd ./internal/services/zipcodeservice/api/
```

## 2 - Execute o arquivo .http (VSCode: REST Client Plugin):
```bash
    zipcodeservice.http
```

### Exemplo

## Requisição com CEP válido

```bash
    curl -X POST "http://localhost:8081/api/v1/zipcode/validate" \
    -H 'Content-Type: application/json' \
    -d '{ "cep": "37275000"}'    
```    

## Resposta
```bash
    {"city":"Cristais","temp_C":19.9,"temp_F":67.82,"temp_K":293.05}
```    


## <a name="license"></a> License

Copyright (c) 2024 [Hugo Castro Costa]

[Hugo Castro Costa]: https://github.com/hgtpcastro