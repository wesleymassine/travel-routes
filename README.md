# Rota de Viagem #

Rota de viagem é uma aplicação orgulhosamente escrita em Golang. Que tem por objetivo encontrar a melhor rota entre dois pontos com
o melhor custo benefício.

## Como executar a aplicação
Para usar a API é necessário ter o [Docker](https://docs.docker.com/get-docker/) instalado.

Após o git clone/instalação, basta usar o comando abaixo na raiz do projeto o [seguinte comando](https://docs.docker.com/compose/reference/up/)

```bash
 docker build -t go/travel .
```
 - Interface Rest: Execute o comando abaixo na pasta raiz do projeto que informa o nome do arquivo input-routes.csv por arqumento:
 ```bash
 sudo docker run -p 5000:5000 -e file=input-routes.csv -t go/travel &
```
Desta forma, a API estará disponível para ser usada na porta [5000](http://localhost:5000)

 - Intergace Cli: Abra outro terminal e na pasta raiz do projeto execute o comando abaixo:
```bash
 ./cli/cli input-routes.csv
```

## Outras alternativas para executar o aplicativo caso não queira utilizar o docker:
- Interface Rest:
Utilize o executável rest na pasta raiz:
 ```bash 
./travel-routes input-routes.csv
```
- Interface Cli:
Utilize o executável cli na pasta raiz: 
```bash 
./cli/cli input-routes.csv
```
## Versões
- Assegure-se de ter Go 1.15+ instalado:

```bash
go version  # go version go1.15.4 linux/amd64
```
## Package
O gorilla/mux pacote de roteamento HTTP em Go (github.com/gorilla/mux) v1.8.0
O urfave/cli para aplicativos de linha de comando em Go	github.com/urfave/cli v1.22.5

## Api Interface rest
* 1. Fligth creation routes
    * URL: http://localhost:5000/flights
    * METHOD: POST 
    * REPONSE: CREATED (201)
    * PAYLOAD: (JSON)

    ```Input: Criação de rota de voos DE-PARA e valor
    {
        "from" : GRU",
        "to"   : "CDG",
        "price" : 20
    }
    ```
    ```Output:
     {
        "from" : "GRU",
        "to"   : "CDG",
        "price" : 20
    }
    ```

* 2.  Fligth consult routes
    * URL: http://localhost:5000/flights
    * METHOD: GET 
    * REPONSE: CREATED (200)
    * PAYLOAD: (JSON)

     ```Input: Informe dois códigos de aeroportos DE-PARA para consultar a melhor rota
    {
        "from" : "BRC",
        "to"   : "CDG"
    }
    ```

    ```Output: Interface API Rest melhor rota
    {
        "price": 40,
        "route": "GRU-BRC-SCL-ORL-CDG"
    }
    ```

## Cli Interface de console
     ```Input: Informe dois códigos de aeroportos DE-PARA para consultar a melhor rota:
        
            Please enter the route: GRU-CDG
     ```      

      ```Output: Interface cli melhor rota:
            
            Best route: "GRU-BRC-SCL-ORL-CDG" > 40
      ```

### Execução com o Postman
Você pode utilizar o link abaixo ou importar o arquivo travel-routes.postman_collection.json da pasta raiz para 
o postman.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/8062a57a73e39480e357)

## Decisões de arquitetura

A escolha da arquitetura MVC foi feita devido ao foco na regra de negócios através do modelo que representa os dados de aplicação e as regras de negócios que regem o acesso e a modificação dos dados. O modelo mantém o estado persistente do negócio e fornece ao controlador a capacidade de acessar as funcionalidades da aplicação encapsuladas pelo próprio modelo.

A aplicação não depende de nenhuma base externa para funcionar. E mantem sua persistência de dados em arquivo CSV.

 - Dijkstra algoritmo

Por se tratar de um problema clássico da ciência da computação, foi utilizado o algoritmo Dijkstra, pois é um
algoritmo bastante utilizado para solucionar problemas de melhor rota em um gráfico.

O algoritmo Dijkstra encontra o caminho mais curto entre os pontos de um gráfico.
Um gráfico é um mapa de pontos e um mapa para os pontos vizinhos no gráfico e o custo para alcançá-los.

Referências: 
* https://pkg.go.dev/github.com/albertorestifo/dijkstra
* https://github.com/skalski/dijkstra-algorithm-golang/blob/master/main.go

## Estrutura de pastas e arquivos
    .
    ├── travel-routes 
        ├── cli                     # Cli Interface de console 
            ├── app                 # Pasta que armazena o arquivo app.go com as funções main de cli
        ├── cli                     # Arquivo main para executar a aplicativos de linha de comando
    ├── file                        # Pasta que armazena o arquivo de persistência CSV
    ├── src                         # Rest aplicação ficheiros do código-fonte
        ├── controllers             # Camada de controller da aplicação  
        ├── exceptions              # Exceptions trata de erros e excessẽos da aplicação
        ├── models                  # Camada do modelo de negócios
        ├── repositories            # Camada de persistência de dados
        ├── responses               # Router manages HTTP requests
        ├── router                  # Gerenciamento de rotas e requisições HTTP request/response
        ├── tests                   # Arquivos de testes da aplicação     
        ├── utils                   # Utils funções genericas auxiliáres
    ├── travel-routes               # Arquivo binário de execução da interface Rest
    ├── docker-compose.yaml
    ├── Dockfile
    ├── go.mod
    ├── go.sum
    ├── main
    └── readme.md

## Testing
Para os testes, foi utilizado somente a biblioteca nativa do Golang.

    Os testes cobrem as principais regras de negócios relacionados a criação rotas de voos, consulta de melhor rota e suas validações.

Para executar os testes utilize os comando abaixo em seu terminal:

    * Na pasta raiz execute: go test -v -coverpkg=./... -coverprofile=coverage.txt ./src/tests/

    * Para visualizar o resultado do coverage no terminal: go tool cover --func=coverage.txt

    * Para visualizar o resultado do coverage no browser: go tool cover --html=coverage.txt

## Licença
[MIT](https://choosealicense.com/licenses/mit/)

## Referências ##
* https://golang.org/

## Challenge Description ##

Um turista deseja viajar pelo mundo pagando o menor preço possível independentemente do número de conexões necessárias.
Vamos construir um programa que facilite ao nosso turista, escolher a melhor rota para sua viagem.

Para isso precisamos inserir as rotas através de um arquivo de entrada.

## Input Example ##
```csv
GRU,BRC,10
BRC,SCL,5
GRU,CDG,75
GRU,SCL,20
GRU,ORL,56
ORL,CDG,5
SCL,ORL,20
```

## Explicando ## 
Caso desejemos viajar de **GRU** para **CDG** existem as seguintes rotas:

1. GRU - BRC - SCL - ORL - CDG ao custo de **$40**
2. GRU - ORL - CGD ao custo de **$64**
3. GRU - CDG ao custo de **$75**
4. GRU - SCL - ORL - CDG ao custo de **$45**

O melhor preço é da rota **1** logo, o output da consulta deve ser **GRU - BRC - SCL - ORL - CDG**.

### Execução do programa ###
A inicializacao do teste se dará por linha de comando onde o primeiro argumento é o arquivo com a lista de rotas inicial.

```shell
$ mysolution input-routes.csv
```

Duas interfaces de consulta devem ser implementadas:
- Interface de console deverá receber um input com a rota no formato "DE-PARA" e imprimir a melhor rota e seu respectivo valor.
  Exemplo:
  
```shell
please enter the route: GRU-CGD
best route: GRU - BRC - SCL - ORL - CDG > $40
please enter the route: BRC-CDG
best route: BRC - ORL > $30
```

- Interface Rest
    A interface Rest deverá suportar:
    - Registro de novas rotas. Essas novas rotas devem ser persistidas no arquivo csv utilizado como input(input-routes.csv),
    - Consulta de melhor rota entre dois pontos.

Também será necessária a implementação de 2 endpoints Rest, um para registro de rotas e outro para consula de melhor rota.

## Recomendações ##
Para uma melhor fluides da nossa conversa, atente-se aos seguintes pontos:

* Envie apenas o código fonte,
* Estruture sua aplicação seguindo as boas práticas de desenvolvimento,
* Evite o uso de frameworks ou bibliotecas externas à linguagem. Utilize apenas o que for necessário para a exposição do serviço,
* Implemente testes unitários seguindo as boas praticas de mercado,
* Documentação
  Em um arquivo Texto ou Markdown descreva:
  * Como executar a aplicação,
  * Estrutura dos arquivos/pacotes,
  * Explique as decisões de design adotadas para a solução,
  * Descreva sua APÌ Rest de forma simplificada.