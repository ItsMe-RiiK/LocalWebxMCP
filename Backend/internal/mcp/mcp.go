package mcp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ItsMe-RiiK/LocalWebxMCP/internal/config"
)

type Request struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  struct {
		Name      string                 `json:"name"`
		Arguments map[string]interface{} `json:"arguments"`
	} `json:"params"`
}

type Response struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Handler acts as the main entry point for the AI client
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request from AI: Method=%s | URL=%s\n", r.Method, r.URL.String())

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "MCP Server Valid", "ready": true}`))
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Use POST method", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("-> AI requesting instruction: %s\n", req.Method)
	resp := Response{JSONRPC: "2.0", ID: req.ID}

	switch req.Method {
	case "initialize":
		resp.Result = map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities":    map[string]interface{}{"tools": map[string]interface{}{}},
			"serverInfo":      map[string]interface{}{"name": "local-go-server", "version": "1.0.0"},
		}
	case "notifications/initialized":
		w.WriteHeader(http.StatusOK)
		return
	case "ping":
		resp.Result = map[string]interface{}{}
	case "tools/list":
		resp.Result = getToolsList()
	case "tools/call":
		handleToolCall(&req, &resp)
	default:
		resp.Error = map[string]interface{}{"code": -32601, "message": "MCP Method Not Supported"}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// getToolsList returns the available MCP tools
func getToolsList() map[string]interface{} {
	return map[string]interface{}{
		"tools": []map[string]interface{}{
			{
				"name":        "list_files",
				"description": "Lists all files in the local storage.",
				"inputSchema": map[string]interface{}{"type": "object", "properties": map[string]interface{}{}},
			},
			{
				"name":        "read_file",
				"description": "Reads the text content of a local file.",
				"inputSchema": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"filename": map[string]interface{}{
							"type":        "string",
							"description": "Full filename (e.g., data.txt)",
						},
					},
					"required": []string{"filename"},
				},
			},
		},
	}
}

// handleToolCall executes the requested AI tool
func handleToolCall(req *Request, resp *Response) {
	switch req.Params.Name {
	case "list_files":
		files, _ := os.ReadDir(config.StorageDir)
		var fileNames []string
		for _, f := range files {
			fileNames = append(fileNames, f.Name())
		}
		resp.Result = map[string]interface{}{
			"content": []map[string]interface{}{{"type": "text", "text": fmt.Sprintf("%v", fileNames)}},
		}

	case "read_file":
		filename, ok := req.Params.Arguments["filename"].(string)
		if !ok {
			resp.Error = map[string]interface{}{"code": -32602, "message": "Invalid filename parameter"}
			return
		}

		content, err := os.ReadFile(filepath.Join(config.StorageDir, filename))
		if err != nil {
			resp.Result = map[string]interface{}{
				"content": []map[string]interface{}{{"type": "text", "text": "Error: File not found"}},
			}
			return
		}

		textContext := string(content)
		if len(textContext) > 50000 {
			textContext = textContext[:50000] + "...(truncated because too large)"
		}
		resp.Result = map[string]interface{}{
			"content": []map[string]interface{}{{"type": "text", "text": textContext}},
		}

	default:
		resp.Error = map[string]interface{}{"code": -32601, "message": "Tool Not Found"}
	}
}
