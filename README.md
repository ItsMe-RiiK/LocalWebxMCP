# Local Web With MCP

MyLocalWeb/
├── Backend/
│   ├── cmd/
│   │   └── main.go                 (Entry point to run the server)
│   ├── internal/
│   │   ├── api/
│   │   │   └── handlers.go         (Frontend API: upload, list files)
│   │   ├── config/
│   │   │   └── config.go           (Global variables like folder names)
│   │   ├── mcp/
│   │   │   └── mcp.go              (AI logic, tools, and handshake)
│   │   ├── middleware/
│   │   │   └── cors.go             (Security and CORS headers)
│   │   └── storage/
│   │       └── manager.go          (Disk operations)
│   └── go.mod
├── Frontend/                       (Your HTML/CSS files)
└── Storage/                        (Your uploaded files)
