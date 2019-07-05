//   Copyright 2018 MSolution.IO
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package ebs

import (
	"context"
	"time"

	"github.com/trackit/jsonlog"

	"github.com/trackit/trackit-server/es"
)

const TypeEBSReport = "ebs-report"
const IndexPrefixEBSReport = "ebs-reports"
const TemplateNameEBSReport = "ebs-reports"

// put the ElasticSearch index for *-ebs-reports indices at startup.
func init() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := es.Client.IndexPutTemplate(TemplateNameEC2Report).BodyString(TemplateEc2Report).Do(ctx)
	if err != nil {
		jsonlog.DefaultLogger.Error("Failed to put ES index EC2Report.", err)
	} else {
		jsonlog.DefaultLogger.Info("Put ES index EC2Report.", res)
		ctxCancel()
	}
}

const TemplateEc2Report = `
{
	"template": "*-ebs-reports",
	"version": 7,
	"mappings": {
		"ebs-report": {
			"properties": {
				"account": {
					"type": "keyword"
				},
				"reportDate": {
					"type": "date"
				},
				"reportType": {
					"type": "keyword"
				},
				"instance": {
					"properties": {
						"id": {
							"type": "keyword"
						},
						"region": {
							"type": "keyword"
						},
						"state": {
							"type": "keyword"
						},
						"purchasing": {
							"type": "keyword"
						},
						"keyPair": {
							"type": "keyword"
						},
						"type": {
							"type": "keyword"
						},
						"platform": {
							"type": "keyword"
						},
						"tags": {
							"type": "nested",
							"properties": {
								"key": {
									"type": "keyword"
								},
								"value": {
									"type": "keyword"
								}
							}
						},
						"costs": {
							"type": "object"
						},
						"stats": {
							"type": "object",
							"properties": {
								"cpu": {
									"type": "object",
									"properties": {
											"average": {
												"type": "double"
											},
											"peak": {
												"type": "double"
											}
									}
								},
								"network": {
									"type": "object",
									"properties": {
											"in": {
												"type": "double"
											},
											"out": {
												"type": "double"
											}
									}
								},
								"volumes": {
									"type": "nested",
									"properties": {
										"id": {
											"type": "keyword"
										},
										"read": {
											"type": "double"
										},
										"write": {
											"type": "double"
										}
									}
								}
							}
						}
					}
				}
			},
			"_all": {
				"enabled": false
			},
			"numeric_detection": false,
			"date_detection": false
		}
	}
}
`
