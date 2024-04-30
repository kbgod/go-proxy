## Quick start

1. Create config.json file
2. Put your config
```json
{
  "host": ":7777",
  "proxy": [
    {
      "path": "/api",
      "target": "http://localhost:8080"
    },
    {
      "path": "/",
      "target": "http://localhost:5173"
    }
  ]
}
```

3. Run
```bash
go run main.go
```

Or run with specific config file
```bash
go run main.go aboba.json
```