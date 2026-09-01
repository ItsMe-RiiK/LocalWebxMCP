package mcp

// getToolsList assembles the schemas from all individual tool files
func getToolsList() map[string]interface{} {
	return map[string]interface{}{
		"tools": []map[string]interface{}{
			schemaListFiles(), // ls -la
			schemaCatFile(),   // cat
			schemaFindFiles(), // find
			schemaGrepFile(),  // grep
			schemaWriteFile(), // write/create
		},
	}
}

// handleToolCall routes the request to the specific tool execution function
func handleToolCall(req *Request, resp *Response) {
	switch req.Params.Name {
	case "list_files":
		executeListFiles(req, resp)
	case "cat_file":
		executeCatFile(req, resp)
	case "find_files":
		executeFindFiles(req, resp)
	case "grep_file":
		executeGrepFile(req, resp)
	case "write_file":
		executeWriteFile(req, resp)
	default:
		resp.Error = map[string]interface{}{"code": -32601, "message": "Tool Not Found"}
	}
}
