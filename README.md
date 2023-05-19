# curso-go-extensive-desafio-multithreading

Executando o desafio:
```
go run src/main.go
```

Simulando timeout da chamada para a api viaCep:
```
SET_VIACEP_TIMEOUT=true go run src/main.go
```

Simulando timeout da chamada para a api cdnApi:
```
SET_CDNAPI_TIMEOUT=true go run src/main.go
```

Simulando timeout das duas chamadas:
```
SET_CDNAPI_TIMEOUT=true SET_VIACEP_TIMEOUT=true go run src/main.go
```