package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ImageTag represents the ImageTag schema from the OpenAPI specification
type ImageTag struct {
	Tagid string `json:"tagId,omitempty"`
	Tagname string `json:"tagName,omitempty"`
	Created string `json:"created,omitempty"`
}

// StoredImagePrediction represents the StoredImagePrediction schema from the OpenAPI specification
type StoredImagePrediction struct {
	Iteration string `json:"iteration,omitempty"`
	Predictions []Prediction `json:"predictions,omitempty"`
	Project string `json:"project,omitempty"`
	Thumbnailuri string `json:"thumbnailUri,omitempty"`
	Created string `json:"created,omitempty"`
	Domain string `json:"domain,omitempty"`
	Id string `json:"id,omitempty"`
	Imageuri string `json:"imageUri,omitempty"`
}

// ImageFileCreateEntry represents the ImageFileCreateEntry schema from the OpenAPI specification
type ImageFileCreateEntry struct {
	Contents string `json:"contents,omitempty"`
	Name string `json:"name,omitempty"`
	Regions []Region `json:"regions,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
}

// ImageFileCreateBatch represents the ImageFileCreateBatch schema from the OpenAPI specification
type ImageFileCreateBatch struct {
	Images []ImageFileCreateEntry `json:"images,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
}

// ProjectSettings represents the ProjectSettings schema from the OpenAPI specification
type ProjectSettings struct {
	Domainid string `json:"domainId,omitempty"` // Gets or sets the id of the Domain to use with this project
}

// BoundingBox represents the BoundingBox schema from the OpenAPI specification
type BoundingBox struct {
	Height float32 `json:"height,omitempty"`
	Left float32 `json:"left,omitempty"`
	Top float32 `json:"top,omitempty"`
	Width float32 `json:"width,omitempty"`
}

// PredictionQueryResult represents the PredictionQueryResult schema from the OpenAPI specification
type PredictionQueryResult struct {
	Results []StoredImagePrediction `json:"results,omitempty"`
	Token PredictionQueryToken `json:"token,omitempty"`
}

// ImageCreateSummary represents the ImageCreateSummary schema from the OpenAPI specification
type ImageCreateSummary struct {
	Images []ImageCreateResult `json:"images,omitempty"`
	Isbatchsuccessful bool `json:"isBatchSuccessful,omitempty"`
}

// ImageRegionCreateEntry represents the ImageRegionCreateEntry schema from the OpenAPI specification
type ImageRegionCreateEntry struct {
	Left float32 `json:"left,omitempty"`
	Tagid string `json:"tagId,omitempty"`
	Top float32 `json:"top,omitempty"`
	Width float32 `json:"width,omitempty"`
	Height float32 `json:"height,omitempty"`
	Imageid string `json:"imageId,omitempty"`
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Thumbnailuri string `json:"thumbnailUri,omitempty"` // Gets the thumbnail url representing the project
	Created string `json:"created,omitempty"` // Gets the date this project was created
	Description string `json:"description,omitempty"` // Gets or sets the description of the project
	Id string `json:"id,omitempty"` // Gets the project id
	Lastmodified string `json:"lastModified,omitempty"` // Gets the date this project was last modified
	Name string `json:"name,omitempty"` // Gets or sets the name of the project
	Settings ProjectSettings `json:"settings,omitempty"` // Represents settings associated with a project
}

// Domain represents the Domain schema from the OpenAPI specification
type Domain struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	TypeField string `json:"type,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Exportable bool `json:"exportable,omitempty"`
}

// ImageUrlCreateBatch represents the ImageUrlCreateBatch schema from the OpenAPI specification
type ImageUrlCreateBatch struct {
	Tagids []string `json:"tagIds,omitempty"`
	Images []ImageUrlCreateEntry `json:"images,omitempty"`
}

// Iteration represents the Iteration schema from the OpenAPI specification
type Iteration struct {
	Id string `json:"id,omitempty"` // Gets the id of the iteration
	Trainedat string `json:"trainedAt,omitempty"` // Gets the time this iteration was last modified
	Exportable bool `json:"exportable,omitempty"` // Whether the iteration can be exported to another format for download
	Name string `json:"name,omitempty"` // Gets or sets the name of the iteration
	Projectid string `json:"projectId,omitempty"` // Gets the project id of the iteration
	Status string `json:"status,omitempty"` // Gets the current iteration status
	Created string `json:"created,omitempty"` // Gets the time this iteration was completed
	Isdefault bool `json:"isDefault,omitempty"` // Gets or sets a value indicating whether the iteration is the default iteration for the project
	Lastmodified string `json:"lastModified,omitempty"` // Gets the time this iteration was last modified
	Domainid string `json:"domainId,omitempty"` // Get or sets a guid of the domain the iteration has been trained on
}

// ImageRegionCreateSummary represents the ImageRegionCreateSummary schema from the OpenAPI specification
type ImageRegionCreateSummary struct {
	Created []ImageRegionCreateResult `json:"created,omitempty"`
	Duplicated []ImageRegionCreateEntry `json:"duplicated,omitempty"`
	Exceeded []ImageRegionCreateEntry `json:"exceeded,omitempty"`
}

// ImageRegionProposal represents the ImageRegionProposal schema from the OpenAPI specification
type ImageRegionProposal struct {
	Imageid string `json:"imageId,omitempty"`
	Projectid string `json:"projectId,omitempty"`
	Proposals []RegionProposal `json:"proposals,omitempty"`
}

// IterationPerformance represents the IterationPerformance schema from the OpenAPI specification
type IterationPerformance struct {
	Pertagperformance []TagPerformance `json:"perTagPerformance,omitempty"` // Gets the per-tag performance details for this iteration
	Precision float32 `json:"precision,omitempty"` // Gets the precision
	Precisionstddeviation float32 `json:"precisionStdDeviation,omitempty"` // Gets the standard deviation for the precision
	Recall float32 `json:"recall,omitempty"` // Gets the recall
	Recallstddeviation float32 `json:"recallStdDeviation,omitempty"` // Gets the standard deviation for the recall
	Averageprecision float32 `json:"averagePrecision,omitempty"` // Gets the average precision when applicable
}

// ImageIdCreateBatch represents the ImageIdCreateBatch schema from the OpenAPI specification
type ImageIdCreateBatch struct {
	Images []ImageIdCreateEntry `json:"images,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
}

// ImageIdCreateEntry represents the ImageIdCreateEntry schema from the OpenAPI specification
type ImageIdCreateEntry struct {
	Id string `json:"id,omitempty"`
	Regions []Region `json:"regions,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
}

// PredictionQueryToken represents the PredictionQueryToken schema from the OpenAPI specification
type PredictionQueryToken struct {
	Starttime string `json:"startTime,omitempty"`
	Orderby string `json:"orderBy,omitempty"`
	Session string `json:"session,omitempty"`
	Continuation string `json:"continuation,omitempty"`
	Endtime string `json:"endTime,omitempty"`
	Application string `json:"application,omitempty"`
	Maxcount int `json:"maxCount,omitempty"`
	Tags []PredictionQueryTag `json:"tags,omitempty"`
	Iterationid string `json:"iterationId,omitempty"`
}

// ImageRegion represents the ImageRegion schema from the OpenAPI specification
type ImageRegion struct {
	Left float32 `json:"left,omitempty"`
	Regionid string `json:"regionId,omitempty"`
	Tagid string `json:"tagId,omitempty"`
	Tagname string `json:"tagName,omitempty"`
	Top float32 `json:"top,omitempty"`
	Width float32 `json:"width,omitempty"`
	Created string `json:"created,omitempty"`
	Height float32 `json:"height,omitempty"`
}

// ImageRegionCreateBatch represents the ImageRegionCreateBatch schema from the OpenAPI specification
type ImageRegionCreateBatch struct {
	Regions []ImageRegionCreateEntry `json:"regions,omitempty"`
}

// RegionProposal represents the RegionProposal schema from the OpenAPI specification
type RegionProposal struct {
	Boundingbox BoundingBox `json:"boundingBox,omitempty"`
	Confidence float32 `json:"confidence,omitempty"`
}

// ImageRegionCreateResult represents the ImageRegionCreateResult schema from the OpenAPI specification
type ImageRegionCreateResult struct {
	Created string `json:"created,omitempty"`
	Imageid string `json:"imageId,omitempty"`
	Tagname string `json:"tagName,omitempty"`
	Top float32 `json:"top,omitempty"`
	Height float32 `json:"height,omitempty"`
	Regionid string `json:"regionId,omitempty"`
	Tagid string `json:"tagId,omitempty"`
	Left float32 `json:"left,omitempty"`
	Width float32 `json:"width,omitempty"`
}

// ImageUrl represents the ImageUrl schema from the OpenAPI specification
type ImageUrl struct {
	Url string `json:"url,omitempty"`
}

// TagPerformance represents the TagPerformance schema from the OpenAPI specification
type TagPerformance struct {
	Precision float32 `json:"precision,omitempty"` // Gets the precision
	Precisionstddeviation float32 `json:"precisionStdDeviation,omitempty"` // Gets the standard deviation for the precision
	Recall float32 `json:"recall,omitempty"` // Gets the recall
	Recallstddeviation float32 `json:"recallStdDeviation,omitempty"` // Gets the standard deviation for the recall
	Averageprecision float32 `json:"averagePrecision,omitempty"` // Gets the average precision when applicable
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ImageTagCreateSummary represents the ImageTagCreateSummary schema from the OpenAPI specification
type ImageTagCreateSummary struct {
	Created []ImageTagCreateEntry `json:"created,omitempty"`
	Duplicated []ImageTagCreateEntry `json:"duplicated,omitempty"`
	Exceeded []ImageTagCreateEntry `json:"exceeded,omitempty"`
}

// Region represents the Region schema from the OpenAPI specification
type Region struct {
	Height float32 `json:"height,omitempty"`
	Left float32 `json:"left,omitempty"`
	Tagid string `json:"tagId,omitempty"`
	Top float32 `json:"top,omitempty"`
	Width float32 `json:"width,omitempty"`
}

// ImagePerformance represents the ImagePerformance schema from the OpenAPI specification
type ImagePerformance struct {
	Tags []ImageTag `json:"tags,omitempty"`
	Created string `json:"created,omitempty"`
	Height int `json:"height,omitempty"`
	Id string `json:"id,omitempty"`
	Imageuri string `json:"imageUri,omitempty"`
	Regions []ImageRegion `json:"regions,omitempty"`
	Width int `json:"width,omitempty"`
	Thumbnailuri string `json:"thumbnailUri,omitempty"`
	Predictions []Prediction `json:"predictions,omitempty"`
}

// ImageCreateResult represents the ImageCreateResult schema from the OpenAPI specification
type ImageCreateResult struct {
	Sourceurl string `json:"sourceUrl,omitempty"`
	Status string `json:"status,omitempty"`
	Image Image `json:"image,omitempty"` // Image model to be sent as JSON
}

// Prediction represents the Prediction schema from the OpenAPI specification
type Prediction struct {
	Tagid string `json:"tagId,omitempty"`
	Tagname string `json:"tagName,omitempty"`
	Boundingbox BoundingBox `json:"boundingBox,omitempty"`
	Probability float32 `json:"probability,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Description string `json:"description,omitempty"` // Gets or sets the description of the tag
	Id string `json:"id,omitempty"` // Gets the Tag ID
	Imagecount int `json:"imageCount,omitempty"` // Gets the number of images with this tag
	Name string `json:"name,omitempty"` // Gets or sets the name of the tag
}

// Export represents the Export schema from the OpenAPI specification
type Export struct {
	Status string `json:"status,omitempty"`
	Downloaduri string `json:"downloadUri,omitempty"`
	Flavor string `json:"flavor,omitempty"`
	Platform string `json:"platform,omitempty"`
}

// ImageTagCreateBatch represents the ImageTagCreateBatch schema from the OpenAPI specification
type ImageTagCreateBatch struct {
	Tags []ImageTagCreateEntry `json:"tags,omitempty"`
}

// ImageUrlCreateEntry represents the ImageUrlCreateEntry schema from the OpenAPI specification
type ImageUrlCreateEntry struct {
	Tagids []string `json:"tagIds,omitempty"`
	Url string `json:"url,omitempty"`
	Regions []Region `json:"regions,omitempty"`
}

// PredictionQueryTag represents the PredictionQueryTag schema from the OpenAPI specification
type PredictionQueryTag struct {
	Id string `json:"id,omitempty"`
	Maxthreshold float32 `json:"maxThreshold,omitempty"`
	Minthreshold float32 `json:"minThreshold,omitempty"`
}

// Image represents the Image schema from the OpenAPI specification
type Image struct {
	Height int `json:"height,omitempty"`
	Id string `json:"id,omitempty"`
	Imageuri string `json:"imageUri,omitempty"`
	Regions []ImageRegion `json:"regions,omitempty"`
	Tags []ImageTag `json:"tags,omitempty"`
	Thumbnailuri string `json:"thumbnailUri,omitempty"`
	Width int `json:"width,omitempty"`
	Created string `json:"created,omitempty"`
}

// ImagePrediction represents the ImagePrediction schema from the OpenAPI specification
type ImagePrediction struct {
	Predictions []Prediction `json:"predictions,omitempty"`
	Project string `json:"project,omitempty"`
	Created string `json:"created,omitempty"`
	Id string `json:"id,omitempty"`
	Iteration string `json:"iteration,omitempty"`
}

// ImageTagCreateEntry represents the ImageTagCreateEntry schema from the OpenAPI specification
type ImageTagCreateEntry struct {
	Tagid string `json:"tagId,omitempty"`
	Imageid string `json:"imageId,omitempty"`
}
