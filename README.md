# Systatus



[![GitHub License](https://img.shields.io/github/license/gdegiorgio/systatus?style=for-the-badge&color=blue&link=https%3A%2F%2Fgithub.com%gdegiorgio%systatus%2Fblob%2Fmain%2FLICENSE)](https://github.com/YourUsername/GopherMetrics/blob/main/LICENSE) 
![Go Reference](https://img.shields.io/badge/reference-grey?style=for-the-badge&logo=go&link=https%3A%2F%2Fgithub.com%gdegiorgio%FsystatUs) 
![pr's welcome](https://img.shields.io/badge/PR'S-WELCOME-green?style=for-the-badge) 
![GitHub Repo stars](https://img.shields.io/github/stars/gdegiorgio/systatus) 
[![GitHub followers](https://img.shields.io/github/followers/gdegiorgio?link=https%3A%2F%2Fgithub.com%2Fgdegiorgio)](https://github.com/gdegiorgio) 
![GitHub forks](https://img.shields.io/github/forks/gdegiorgio/systatus)

> **Note:** Systatus is currently in beta, and we are actively working on expanding functionality. Be sure to update frequently to get the latest features and improvements.

<img src="./resources/assets/systatus.png" align="right">

### What is Systatus?

Systatus is a lightweight Go library inspired by Spring Boot's Actuator, designed to add system observability and monitoring endpoints to any Go application. It allows you to expose essential system information and application health metrics through HTTP routes, enabling quick insights and diagnostics.

With just a of code, Systatus can provide your application with predefined routes to monitor metrics like CPU, memory, and disk usage, as well as application uptime and a simple health check.




### Available Routes

Systatus provides the following default routes, which are also documented in [`swagger.yml`](./swagger.yml):

- `/env`: JSON response with host environment variables
- `/health`: Simnple Health check route



### Installation

To install Systatus, use the following command:

```bash
go get github.com/gdegiorgio/systatus
```


###  Quick startup

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
