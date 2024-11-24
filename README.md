# Systatus 🔎



[![GitHub License](https://img.shields.io/github/license/gdegiorgio/systatus?style=for-the-badge&color=blue&link=https%3A%2F%2Fgithub.com%gdegiorgio%systatus%2Fblob%2Fmain%2FLICENSE)](https://github.com/YourUsername/GopherMetrics/blob/main/LICENSE) 
![Go Reference](https://img.shields.io/badge/reference-grey?style=for-the-badge&logo=go&link=https%3A%2F%2Fgithub.com%gdegiorgio%FsystatUs) 

> **Note:** Systatus is currently in beta, and we are actively working on expanding functionality. Be sure to update frequently to get the latest features and improvements.

<img src="./resources/assets/systatus.png" align="right">

### What is Systatus?

Systatus is a lightweight Go library inspired by [Spring Boot's Actuator](https://docs.spring.io/spring-boot/reference/actuator/index.html), designed to add system observability and monitoring endpoints to any Go application. It allows you to expose essential system information and application health metrics through HTTP routes, enabling quick insights and diagnostics.
With just two lines of code, Systatus can provide your application with predefined routes to monitor metrics like CPU, memory, and disk usage, as well as application uptime and a simple health check.



- [Quick start](#installation)
- [Systatys Options](#systatus-options)
- [Available Handlers](#available-handlers)
  - [`/health`](#health)
  - [`/env`](#env)
  - [`/cpu`](#cpu)
  - [`/disk`](#disk)
  - [`/mem`](#mem)
  - [`/uptime`](#uptime)


### Quickstart

Install Systatus in your go project:

```bash
go get github.com/gdegiorgio/systatus
```

Then import and simply use the `Enable` method

```golang
package main

import (
	"fmt"
	"net/http"
	"github.com/gdegiorgio/systatus"
)

func main() {
	opts := SystatusOptions{ Prefix : "/dev"}
	systatus.Enable(opts)
	http.ListenAndServe(":8080", nil)
}
```

### Systatus Options

| _**Option**_         | **_Type_**                   | **_Default_** | **_Description_**                                                                                                  |
|----------------------|------------------------------|---------------|--------------------------------------------------------------------------------------------------------------------|
| `Prefix`             | `string`                     | `""`          | Specifies a URL prefix for all systatus endpoints. For example, setting `Prefix: "/dev"` results in `/dev/health`. |
| `ExposeEnv`          | `boolean`                    | `false`       | Enables the `/env` endpoint, exposing environment variables. Use cautiously as this may reveal sensitive data.     |
| `PrettyLogger`       | `boolean`                    | `false`       | 	If set to `true`, replaces JSON-formatted logging with human-readable plain text logs.                            |
| `HealthHandlerOpts`  | [`/health`](#health) options | `nil`         | Configuration for the `/health` handler                                                                            |
| `CPUHandlerOpts`     | [`/cpu`](#cpu) options       | `nil`         | Configuration for the `/cpu` handler                                                                               |
| `EnvHandlerOpts`     | [`/env`](#env) options       | `nil`         | Configuration for the `/env` handler                                                                               |
| `DiskHandlerOpts`    | [`/disk`](#disk) options     | `nil`         | Configuration for the `/disk` handler                                                                              |
| `MemHandlerOpts`     | [`/mem`](#mem) options       | `nil`         | Configuration for the `/mem` handler                                                                               |
| 	`UptimeHandlerOpts` | [`/uptime`](#uptime) options | `nil`         | UConfiguration for the `/uptime` handler                                                                           |

#### Notes:
- **Options with nil defaults**: These options are optional and customizable through their respective configurations. Reference specific handler documentation (linked in the table) for details.
- **Logging Behavior**: Setting PrettyLogger to true is recommended for debugging purposes but might increase log size due to less structured output.
- **Prefix Best Practices**: Use Prefix to isolate endpoints in shared environments, e.g., /api/systatus.



## Available Handlers

### `/health`

This endpoint is used to report system health status.


#### Options



| _**Option**_  | **_Type_**                                       | **_Default_** | **_Description_**                                                                                                                                               |
|---------------|--------------------------------------------------|---------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `Healtcheck`  | `func(w http.ResponseWriter, r *http.Request)`   | `nil`         | Allows users to define a custom health check function. This function overrides the default /health handler logic and must write the HTTP response directly.     |
| `Middlewares` | `[]func(next http.HandlerFunc) http.HandlerFunc` | `[]`          | A list of middleware functions applied to the request pipeline. Each middleware wraps the handler, enabling custom preprocessing or postprocessing of requests. | 

#### HealthResponse


| _**Option**_ | **_Type_**                                                                                   | **_Description_**     |
|--------------|----------------------------------------------------------------------------------------------|-----------------------|
| `status`     | `string`                                                                                     | Current System status |


```json
{
  "status" : "HEALTHY"
}
```

### `/env`

This endpoint is used to expose system environment variables.
> **Note:** This endpoint will not be exposed by default unless you set `ExposeEnv` to `true` in [`SystatusOptions`](#systatus-options).


#### Options


| _**Option**_    | **_Type_**                                       | **_Default_** | **_Description_**                                                                                                                                               |
|-----------------|--------------------------------------------------|---------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `SensitiveKeys` | `[]string`                                       | `[]`          | A list of key names whose values should be masked in logs or responses to prevent exposure of sensitive data.                                                   |
| `Middlewares`   | `[]func(next http.HandlerFunc) http.HandlerFunc` | `[]`          | A list of middleware functions applied to the request pipeline. Each middleware wraps the handler, enabling custom preprocessing or postprocessing of requests. | 

####  EnvResponse


| _**Field**_ | **_Type_**          | **_Description_**                              |
|-------------|---------------------|------------------------------------------------|
| `env`       | `map[string]string` | A map representing system environment variable |

```json
{
  "env" : {
    "GO_VERSION": "1.23.0",
    "APIKEY" : "************************"
  }
}

```
### `/cpu`

This endpoint is not available yet and will soon be implemented.

### `/mem`

This endpoint is used to retrieve runtime memory usage statistics

#### Options


| _**Option**_    | **_Type_**                                       | **_Default_** | **_Description_**                                                                                                                                               |
|-----------------|--------------------------------------------------|---------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `SensitiveKeys` | `[]string`                                       | `[]`          | A list of key names whose values should be masked in logs or responses to prevent exposure of sensitive data.                                                   |
| `Middlewares`   | `[]func(next http.HandlerFunc) http.HandlerFunc` | `[]`          | A list of middleware functions applied to the request pipeline. Each middleware wraps the handler, enabling custom preprocessing or postprocessing of requests. | 

####  MemResponse


| _**Field**_  | **_Type_** | **_Description_**                                                                                    |
|--------------|------------|------------------------------------------------------------------------------------------------------|
| `TotalAlloc` | `uint64`   | The cumulative bytes allocated and freed by the application.                                         |
| `Alloc`      | `uint64`   | The current number of bytes allocated in heap objects.                                               |
| `Sys`        | `uint64`   | The total bytes of memory obtained from the operating system (includes heap, stack, and other uses). |

```json
{
  "total_alloc": 1048576,
  "alloc": 524288,
  "sys": 2097152
}

```

### `/uptime`




#### Options


| _**Option**_    | **_Type_**                                       | **_Default_** | **_Description_**                                                                                                                                               |
|-----------------|--------------------------------------------------|---------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `Middlewares`   | `[]func(next http.HandlerFunc) http.HandlerFunc` | `[]`          | A list of middleware functions applied to the request pipeline. Each middleware wraps the handler, enabling custom preprocessing or postprocessing of requests. | 



#### UptimeResponse

Systime	string	The current system time in a human-readable format (e.g., ISO 8601).
Uptime	string	The server's uptime since it was started, formatted as hh:mm:ss.

| _**Field**_ | **_Type_** | **_Description_**                                                                                    |
|-------------|------------|------------------------------------------------------------------------------------------------------|
| `Systime`   | `string`   | The current system time in a human-readable format (e.g., ISO 8601)                                        |
| `Uptime`    | `string`   | The server's uptime since it was started, formatted as hh:mm:ss.                                             |

```json
{
  "systime": "2024-11-24T12:34:56Z",
  "uptime": "01:23:45"
}
```
