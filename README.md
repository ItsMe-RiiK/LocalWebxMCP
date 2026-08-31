# 🚀 LocalWeb x MCP

![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-323330?style=for-the-badge&logo=javascript&logoColor=F7DF1E)

**LocalWeb x MCP** is a lightweight, high-performance local web server built with Go (Golang). It features a sleek,
modern UI for managing local files and a built-in **Model Context Protocol (MCP)** endpoint that allows AI
assistants (like Gemini Spark) to securely interact with your local environment.

---

## ✨ Key Features

*   **📂 Modern File Management:** A responsive, dark-mode Web UI with drag-and-drop support for easy local file uploads.
*   **🤖 Native MCP Endpoint:** A dedicated `/mcp` route that handshakes with AI models, granting them specific tools to
    interact with your data.
*   **🛠️ AI Tools Included:** AI can dynamically execute `list_files` and `read_file`
    to analyze your local documents in real-time.
*   **🚀 Zero Dependencies:** The Go backend uses standard libraries (`net/http`, `os`, `encoding/json`) for maximum
    performance and minimal footprint.
*   **🔒 Secure by Default:** Built-in CORS middleware to safely handle requests from external AI web clients.

---

## 🏗️ Project Architecture

This repository follows standard Go project layouts (Monorepo setup):

```text
LocalWebxMCP/
├── Backend/
│   ├── cmd/
│   │   └── main.go                 # Main entry point for the server
│   └── internal/                   # Modular internal packages
│       ├── api/                    # Web UI API handlers (upload, list)
│       ├── config/                 # Global configuration constants
│       ├── mcp/                    # MCP handshake and AI tool logic
│       ├── middleware/             # CORS and security middlewares
│       └── storage/                # Local disk management
├── Frontend/
│   ├── index.html                  # Main UI structure
│   ├── css/style.css               # Modern dark-mode styling
│   └── js/
│       ├── api.js                  # Frontend-to-Backend fetch logic
│       └── app.js                  # Drag-and-drop & DOM manipulation
├── Storage/                        # Auto-generated directory for uploaded files
├── go.mod                          # Go module definition
└── README.md                       # This documentation
```

---

## 🚀 Getting Started

### 1. Prerequisites
*   [Go](https://go.dev/dl/) (version 1.21 or higher) installed on your machine.
*   Git installed.

### 2. Installation
Clone this repository to your local machine:
```bash
git clone https://github.com/ItsMe-RiiK/LocalWebxMCP.git
cd LocalWebxMCP
```

### 3. Run the Server
Execute the application directly from the root directory:
```bash
go run Backend/cmd/main.go
```
*The server will start instantly. It will automatically create a `Storage/` directory if it doesn't exist.*

### 4. Access the UI
Open your favorite web browser and navigate to:
👉 **`http://localhost:8080`**

---

## 🤖 Connecting to an AI Client (MCP Setup)

To allow a cloud-based AI (like Gemini Spark) to communicate with your local MCP endpoint, you need to expose your local port `8080` to the internet using a secure SSH tunnel.

### Option A: Using localhost.run (Recommended)
Generate an SSH key (if you haven't already), and run:
```bash
ssh -R 80:127.0.0.1:8080 localhost.run
```
*(This gives you a persistent public URL if you register your SSH key).*

### Option B: Using Serveo
```bash
ssh -R your-custom-name:80:127.0.0.1:8080 serveo.net
```

### Adding to the AI Client
Once you have your public URL (e.g., `[https://your-tunnel-url.lhr.life](https://your-tunnel-url.lhr.life)`), append `/mcp` to it. Provide this full URL to your AI client's configuration:
👉 **`[https://your-tunnel-url.lhr.life/mcp](https://your-tunnel-url.lhr.life/mcp)`**

---

## 🔌 API & Endpoints Summary

### Web UI Endpoints
*   `GET /` - Serves the frontend UI.
*   `GET /api/files` - Returns a JSON array of uploaded files.
*   `POST /api/upload` - Accepts `multipart/form-data` to save files to the `Storage/` directory.
*   `GET /api/files/{filename}` - Direct access to download/view an uploaded file.

### AI Endpoint
*   `POST /mcp` - The core Model Context Protocol interface. It handles `initialize`, `ping`, `tools/list`, and `tools/call`.

---

## 📝 License
This project is licensed under the [Apache 2.0 License](license).
