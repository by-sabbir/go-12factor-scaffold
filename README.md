# Go Rest micro-service

---
Using 12Factor Methodology

## Instructions and Usage

```bash
go build -o app main.go
./app
```

output:

```bash
Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  migrate     Migrate Database
  server      runs the server

Flags:
      --config string   config file (default is ./config.yaml)
  -h, --help            help for 12Factor-App
  -t, --toggle          Help message for toggle
```

### [The Twelve Factors](https://12factor.net/)

---

I. **Codebase**
&nbsp;&nbsp;&nbsp;&nbsp;Deployment is maintained from branches ie. `master -> production`, `dev -> test`, `staging -> uat`.

II. **Dependencies**
&nbsp;&nbsp;&nbsp;&nbsp;In our case, `go.mod` files will seperate all the dependencies specifically for your project.

III. **Config**
&nbsp;&nbsp;&nbsp;&nbsp;In this project we can pass configurations as parameter, `go run main.go server --config <yaml path>`.

IV. **Backing services**
&nbsp;&nbsp;&nbsp;&nbsp;We used postgres and rabbitmq as the backing services and they can be attached from the config as we need.

V. Build, release, run
Strictly separate build and run stages
VI. Processes
Execute the app as one or more stateless processes
VII. Port binding
Export services via port binding
VIII. Concurrency
Scale out via the process model
IX. Disposability
Maximize robustness with fast startup and graceful shutdown
X. Dev/prod parity
Keep development, staging, and production as similar as possible
XI. Logs
Treat logs as event streams
XII. Admin processes
Run admin/management tasks as one-off processes