Structure

go-todo/
│
├── cmd/
│   └── server/
│       └── main.go         # entrypoint
│
├── internal/
│   ├── todo/               # todo module
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── model.go
│   ├── config/             # config loading
│   └── database/           # db connection
│
├── pkg/
│   └── response/           # reusable response helpers
│
├── go.mod
└── go.sum