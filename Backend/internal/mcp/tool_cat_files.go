package mcp

import (
	"os"
	"path/filepath"

	"github.com/ItsMe-RiiK/LocalWebxMCP/Backend/internal/config"
)

func schemaCatFile() map[string]interface{} {
	return map[string]interface{}{
		"name":        "cat_file",
		"description": "Reads and outputs the entire text content of a local file (like 'cat').",
		"inputSchema": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"filename": map[string]interface{}{"type": "string", "description": "Full filename (e.g., data.txt)"},
			},
			"required": []string{"filename"},
		},
	}
}

func executeCatFile(req *Request, resp *Response) {
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
		textContext = textContext[:50000] + "\n...(truncated because too large)"
	}
	resp.Result = map[string]interface{}{
		"content": []map[string]interface{}{{"type": "text", "text": textContext}},
	}
}
