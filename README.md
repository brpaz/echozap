
# echozap

> Middleware for Golang [Echo](https://echo.labstack.com/) framework that provides integration with Uber´s [Zap](https://github.com/uber-go/zap)  logging library for logging HTTP requests.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](LICENSE)
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg?style=for-the-badge)](http://commitizen.github.io/cz-cli/)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg?style=for-the-badge)](https://github.com/semantic-release/semantic-release?style=for-the-badge)

[![Actions Status](https://github.com/brpaz/echozap/workflows/CI/badge.svg?style=for-the-badge)](https://github.com/brpaz/echozap/actions)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/99c5875d156440c0b861dad80c76c01f)](https://www.codacy.com/manual/brpaz/echozap?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=brpaz/echozap&amp;utm_campaign=Badge_Grade)
[![Codacy Badge](https://api.codacy.com/project/badge/Coverage/99c5875d156440c0b861dad80c76c01f)](https://www.codacy.com/manual/brpaz/echozap?utm_source=github.com&utm_medium=referral&utm_content=brpaz/echozap&utm_campaign=Badge_Coverage)

## Pre-requisites

*  Go with Go modules enabled.
*  [Echo v4](https://echo.labstack.com/)
*  [Zap](https://github.com/uber-go/zap)

## Usage

```go
package main

import (
	"net/http"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	e := echo.New()

	zapLogger, _ := zap.NewProduction()

	e.Use(echozap.ZapLogger(zapLogger))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

## Logged details

The following information is logged:

*  Status Code
*  Time
*  Uri
*  Method
*  Hostname
*  Remote IP Address

## Todo

*  Add more customization options.

## 🤝 Contributing

Contributions, issues and feature requests are welcome!

## Show your support

If this project have been useful for you, I would be grateful to have your support.

Give a ⭐️ to the project, or just:

<a href="https://www.buymeacoffee.com/Z1Bu6asGV" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

## Author

👤 **Bruno Paz**

*  Website: [https://github.com/brpaz](https://github.com/brpaz)
*  Github: [@brpaz](https://github.com/brpaz)

## 📝 License

Copyright © 2019 [Bruno Paz](https://github.com/brpaz).

This project is [MIT](LICENSE) licensed.
