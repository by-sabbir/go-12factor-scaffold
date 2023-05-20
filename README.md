# Go Rest micro-service

---
Using 12Factor Methodology

## Instructions and Usage

|   Action  |  Command      |
|-----------|---------------|
|**Build**  |  `make build` |
|**Migrate**| `make migrate`|
|**Deploy** | `make run`    |


###Help Text

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

V. **Build, release, run**
Strictly separate build and run stages, means we have to maintain a CICD to seperate build, test, and deploy environments. For this example, I have seperated the stages with [Makefile](./Makefile)

VI. **Processes**
This means we have to maintain our backend as a stateless independent process so that we can scale better. We are managing the application state in `rabbitmq` and it's fully independent.

VII. **Port binding**
Export services via port binding, in our case `9091`

VIII. **Concurrency**
Concurrency is built into go.

IX. **Disposability**
I have used `SIGTERM` and `SIGKILL` for graceful shutdown. So that, the app waits for a specific amount of time to process existing concurrent request before it forcefully shuts down.

X. **Dev/prod parity**
Again this depends on how we setup/plan our CICD. For our application we have `config.yaml` and `dev.yaml` for simulating this case.
Also for gin specific application, we can sperate debug/production with a environment variable - `GIN_MODE=release`

XI. **Logs**
I have used `logrus` for structure logging.

XII. **Admin processes**
[`cobra`](https://github.com/spf13/cobra) is used to create admin processes. Following admin processes are available in our app:

```bash
  migrate     Migrate Database with golang-migrate
    up        Apply Migration
    down      Undo Migration (not implemented)
  server      runs the server
```
