package tagginges

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic"
	"github.com/trackit/jsonlog"

	"github.com/trackit/trackit/tagging/utils"
)

type domain struct {
	ID     string              `json:"domainId"`
	Region string              `json:"region"`
	Tags   []utils.TagDocument `json:"tags"`
}

type source struct {
	Domain domain `json:"domain"`
}

const sourceIndexName = "es-reports"
const urlFormat = "TODO"

// Process generates tagging reports from ES reports
func Process(ctx context.Context, account int, awsAccount string, resourceTypeString string) ([]utils.TaggingReportDocument, error) {
	logger := jsonlog.LoggerFromContextOrDefault(ctx)
	logger.Info(fmt.Sprintf("Processing %s reports.", resourceTypeString), nil)

	hits, err := fetchReports(ctx, account)
	if err != nil {
		return nil, err
	}

	var documents []utils.TaggingReportDocument
	for _, hit := range hits {
		document, success := processHit(ctx, hit, awsAccount, resourceTypeString)
		if success {
			documents = append(documents, document)
		}
	}

	logger.Info(fmt.Sprintf("%d %s reports processed.", len(documents), resourceTypeString), nil)
	return documents, nil
}

// processHit converts an elasticSearch hit into a TaggingReportDocument
// Second argument is true if operation is a success
func processHit(ctx context.Context, hit *elastic.SearchHit, awsAccount string, resourceTypeString string) (utils.TaggingReportDocument, bool) {
	logger := jsonlog.LoggerFromContextOrDefault(ctx)
	var source source
	err := json.Unmarshal(*hit.Source, &source)
	if err != nil {
		logger.Error(fmt.Sprintf("Could not process a %s report.", resourceTypeString), nil)
		return utils.TaggingReportDocument{}, false
	}

	// regionForURL := utils.GetRegionForURL(source.Snapshot.Region)

	document := utils.TaggingReportDocument{
		Account:      awsAccount,
		ResourceID:   source.Domain.ID,
		ResourceType: resourceTypeString,
		Region:       source.Domain.Region,
		URL:          fmt.Sprintf(urlFormat),
		Tags:         source.Domain.Tags,
	}
	return document, true
}
