### A Demo for Go RESTful API

Demostrate how to build api in a Go way that deliver Scalability and Extensibility.

##### Caveat: 

Not very polished in style though, you might delete whenever you like. 
       
Nevertheless, Logic is straightforward.

This project is capable of:
    
    1.  Simple deployment
        Deployed as Binary file. 
        Drag and drop, no dependencies relied on Production environment.
        Run like a bash cli.
        
    2.  Greate performance
        SQL queries are logged in console for detailed view of cost.
    
    3.  Greate modular code structure
        Normally, go project are not well-structured. This project triage files into right place.
    
    4.  DB first
        - Performance-care: 
          DBA can work on more features on db side. Backend can generate `class` to take advantage of these.
        Without writing explicite line to manage db specific details.
        - Less ORM, less M2M to 1 to many bindings. Class based ORM make things complicated and un-necessary.
    
    5.  Echo is built on top of Gin (i think L:) 
        You can plug and play Gin libraries to replace Echo. Or maybe your own.
    
    6.  Query builder is open and easy. Not like sqlx or other raw methods: Scanner(&q)
        - Generate SQL in function like GetPost() either eagerly or lazily.
    
    7.  Clean DTO for serializer or deserializer.
        - handle day to day life
    
    8.  Store has data layer CRUD api. Independent of upper Web layer.
    
    9.  Web layer api are in handler. Bit like Controller / Service folder.

## Main Folder structure:

models - generated go structs, interface to operate db

store - data layer, protocal

handler - web layer, HTTP controller, gRPC as well

router - middleware

url.go - as high level API blueprint: [location](https://github.com/ilovejs/Go-Echo-Boiler/blob/feature%2Fupdate_model_use_boiler/handler/url.go)

### Install
- Windows:
1. run `SETX GO111MODULES on`
2. open new console && cd to project
3. Pick package 
```
go get -u github.com/sirupsen/logrus
go get -u github.com/labstack/echo/...
go get -u github.com/labstack/echo
go get -u -v github.com/kat-co/vala
```

Code hot reload [air](https://github.com/cosmtrek/air)

4. now check go.mod values 
5. run magic command to download source and create vendor:
```
go mod download
go mod vendor
```

### If not using vendor mode Go module in different folder

copy onsite folder to GOPATH

`go get -v`

### Library patch in GoPATH

[1]
sqlboilder / queries / query_builders.go
LINE 238 and 220:
fmt.Fprintf(buf, " OFFSET %d ROWS", q.offset)

### Deployment

[Docker and Host bi-direction](https://docs.docker.com/docker-for-mac/networking/#use-cases-and-workarounds)
  
  #### Steps:
    
    * folder
    
        cd C:\Users\mzhuang\go\src\onsite

    * file name is 'app'

        gox -osarch="linux/amd64" --output="build/app"

    ```{bash}
    docker build -t durgaprasad-budhwani/azurego .
    docker image ls
    docker run -p 8585:8585 --env PORT=8585 durgaprasad-budhwani/azurego
    
    docker run --publish 8585:8585 --net="host" --env PORT=8585 durgaprasad-budhwani/azurego
    docker run --publish 85:8585 --net="host" durgaprasad-budhwani/azurego
    ```

### Azure snippets

#### Create web app
    az webapp create --resource-group AzureGoRG --plan AzureGoSP --name AzureGoApp --deployment-container-image-name durgaprasadbudhwani/azurego:latest

#### Set env for docker image
    az webapp config appsettings set --resource-group AzureGoRG --name AzureGoApp --settings PORT=80 connection_string="connection_string"

#### Registry login server
    index.docker.io
    ilovejs/idk:latest

#### Example in 'azure-go-labs'
    - https://github.com/Azure-Samples/azure-go-labs/blob/master/3-web-app-aci-aks/Dockerfile
    - https://github.com/Azure-Samples/azure-go-labs/blob/master/1-app-hello-echo/README.md
#### Issues:
    1. Linux worker can't be created under windows resource group.
        Web app service use linux Docker container can not be used under old resource group incl vnet, azure sql server in subnet.
        - windows container is preview and on expensive teir: premium
        Q: how to change RG to mix up
### Dev logs

Project Serial No | Name should be unique

- Project Update

    u creator_id and manager_id MUST exists OR FK issue from db
        sol: detect that in store level
    u need to have permission system
        not every user can modify creator,manager etc.
    + show latest updates so as to notify user.

- Trade
    
    skip field 'temp'

- Double check Echo version before compile in production.    
