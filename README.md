# e-commerce

Web app for buying goods

## Architecture

```console
george@hotz:e-commerce$ tree -L 1 --dirsfirst
├── client
|   # all client code
├── configs
|   # main configs (for backend)
├── constants
|   # constants (for backend)
├── graph
|   # graphql API
├── handlers
|   # http handlers
├── middleware
├── models
|   # mongodb models
├── services
|   # mongodb services
├── tests
├── utils
|   # helpers
├── Dockerfile
├── go.mod
├── go.sum
├── gqlgen.yml
|   # yml file to generate gql schema
├── README.md
├── server.go
|   # server init
└── tasks.md
    # dev notes/tasks
```
