package middleware

var whiteListUrl = map[string]bool{ //api url白名单
	"/swagger/index.html":                      true,
	"/swagger/swagger-ui.css":                  true,
	"/swagger/swagger-ui-bundle.js":            true,
	"/swagger/swagger-ui-standalone-preset.js": true,
	"/swagger/favicon-32x32.png":               true,
	"/swagger/favicon-16x16.png":               true,
	"/swagger/doc.json":                        true,
	"/cloud/pprof/":                            true,
	"/cloud/api/user/login":                    true,
	"/cloud/api/device/heartbeat":              true,
}
