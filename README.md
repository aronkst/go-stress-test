# Go StressTest

Este é um CLI em Go para realizar testes de carga em um serviço web.

## Pré-requisitos

Certifique-se de ter o Docker instalado no seu sistema. Você pode baixar e instalar o Docker a partir do [site oficial do Docker](https://www.docker.com/).

## Uso com Docker

1. Clone o repositório:

```bash
git clone https://github.com/aronkst/go-stress-test.git
```

2. Navegue até o diretório do projeto:

```bash
cd go-stress-test
```

3. Construa a imagem Docker:

```bash
docker build -t go-stress-test .
```

4. Execute o CLI usando Docker, com os parâmetros desejados:

```bash
docker run --rm go-stress-test --url <URL_DO_SERVIÇO> --requests <NÚMERO_DE_REQUESTS> --concurrency <CHAMADAS_SIMULTÂNEAS>
```

Substitua `<URL_DO_SERVIÇO>`, `<NÚMERO_DE_REQUESTS>` e `<CHAMADAS_SIMULTÂNEAS>` pelos valores desejados.

## Estrutura do Projeto

- `cmd/stresstest/main.go`: Contém a lógica principal do CLI.
- `internal/stresstest/usecase/stresstest_usecase.go`: Implementação do caso de uso para realizar testes de carga.

## Observation

- O retorno do relatório está em inglês.
