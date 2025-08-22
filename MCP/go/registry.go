package main

import (
	"github.com/trainingapi/mcp-server/config"
	"github.com/trainingapi/mcp-server/models"
	tools_imageapi "github.com/trainingapi/mcp-server/tools/imageapi"
	tools_tagsapi "github.com/trainingapi/mcp-server/tools/tagsapi"
	tools_projectapi "github.com/trainingapi/mcp-server/tools/projectapi"
	tools_predictionsapi "github.com/trainingapi/mcp-server/tools/predictionsapi"
	tools_imageregionproposalapi "github.com/trainingapi/mcp-server/tools/imageregionproposalapi"
	tools_domainsapi "github.com/trainingapi/mcp-server/tools/domainsapi"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_imageapi.CreateCreateimagesfrompredictionsTool(cfg),
		tools_tagsapi.CreateDeletetagTool(cfg),
		tools_tagsapi.CreateGettagTool(cfg),
		tools_tagsapi.CreateUpdatetagTool(cfg),
		tools_projectapi.CreateGetiterationsTool(cfg),
		tools_tagsapi.CreateGettagsTool(cfg),
		tools_tagsapi.CreateCreatetagTool(cfg),
		tools_imageapi.CreateDeleteimagetagsTool(cfg),
		tools_imageapi.CreateCreateimagetagsTool(cfg),
		tools_projectapi.CreateGetimageperformancesTool(cfg),
		tools_imageapi.CreateGetimagesbyidsTool(cfg),
		tools_imageapi.CreateGetuntaggedimagesTool(cfg),
		tools_imageapi.CreateCreateimagesfromurlsTool(cfg),
		tools_imageapi.CreateDeleteimagesTool(cfg),
		tools_predictionsapi.CreateQuerypredictionsTool(cfg),
		tools_predictionsapi.CreateQuicktestimageurlTool(cfg),
		tools_imageregionproposalapi.CreateGetimageregionproposalsTool(cfg),
		tools_imageapi.CreateGettaggedimagecountTool(cfg),
		tools_projectapi.CreateTrainprojectTool(cfg),
		tools_projectapi.CreateDeleteprojectTool(cfg),
		tools_projectapi.CreateGetprojectTool(cfg),
		tools_projectapi.CreateUpdateprojectTool(cfg),
		tools_projectapi.CreateGetiterationperformanceTool(cfg),
		tools_projectapi.CreateGetimageperformancecountTool(cfg),
		tools_projectapi.CreateDeleteiterationTool(cfg),
		tools_projectapi.CreateGetiterationTool(cfg),
		tools_projectapi.CreateUpdateiterationTool(cfg),
		tools_projectapi.CreateGetexportsTool(cfg),
		tools_projectapi.CreateExportiterationTool(cfg),
		tools_imageapi.CreateDeleteimageregionsTool(cfg),
		tools_imageapi.CreateCreateimageregionsTool(cfg),
		tools_projectapi.CreateGetprojectsTool(cfg),
		tools_projectapi.CreateCreateprojectTool(cfg),
		tools_imageapi.CreateGettaggedimagesTool(cfg),
		tools_predictionsapi.CreateDeletepredictionTool(cfg),
		tools_imageapi.CreateGetuntaggedimagecountTool(cfg),
		tools_domainsapi.CreateGetdomainsTool(cfg),
		tools_domainsapi.CreateGetdomainTool(cfg),
		tools_imageapi.CreateCreateimagesfromfilesTool(cfg),
	}
}
