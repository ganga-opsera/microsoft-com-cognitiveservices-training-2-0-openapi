package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/trainingapi/mcp-server/config"
	"github.com/trainingapi/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UpdateprojectHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		projectIdVal, ok := args["projectId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: projectId"), nil
		}
		projectId, ok := projectIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: projectId"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Project
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/projects/%s", cfg.BaseURL, projectId)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.APIKey != "" {
			req.Header.Set("Training-Key", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["Training-Key"]; ok {
			req.Header.Set("Training-Key", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Project
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateUpdateprojectTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_projects_projectId",
		mcp.WithDescription("Update a specific project"),
		mcp.WithString("projectId", mcp.Required(), mcp.Description("The id of the project to update")),
		mcp.WithString("Training-Key", mcp.Required(), mcp.Description("")),
		mcp.WithString("name", mcp.Description("Input parameter: Gets or sets the name of the project")),
		mcp.WithObject("settings", mcp.Description("Input parameter: Represents settings associated with a project")),
		mcp.WithString("thumbnailUri", mcp.Description("Input parameter: Gets the thumbnail url representing the project")),
		mcp.WithString("created", mcp.Description("Input parameter: Gets the date this project was created")),
		mcp.WithString("description", mcp.Description("Input parameter: Gets or sets the description of the project")),
		mcp.WithString("id", mcp.Description("Input parameter: Gets the project id")),
		mcp.WithString("lastModified", mcp.Description("Input parameter: Gets the date this project was last modified")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdateprojectHandler(cfg),
	}
}
