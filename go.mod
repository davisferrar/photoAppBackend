module github.com/chhoengichen/distributed-system-mission1

go 1.21.5

replace github.com/chhoengichen/distributed-system-mission1/handler => ./handler

replace github.com/chhoengichen/distributed-system-mission1/server => ./server

replace github.com/chhoengichen/distributed-system-mission1/route => ./route

require (
	github.com/gofiber/fiber/v2 v2.51.0
	github.com/joho/godotenv v1.5.1
)

require (
	github.com/andybalholm/brotli v1.0.6 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)
