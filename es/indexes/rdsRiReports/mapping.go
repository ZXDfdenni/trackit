//   Copyright 2020 MSolution.IO
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

package rdsRiReports

const Template = `
{
	"template": "*-rds-ri-reports",
	"version": 2,
	"mappings": {
		"rds-ri-report": {
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
					"type": "object",
					"properties": {
						"id": {
							"type": "keyword"
						},
						"offeringId": {
							"type": "keyword"
						},
						"availabilityZone": {
							"type": "keyword"
						},
						"instanceClass": {
							"type": "keyword"
						},
						"instanceCount": {
							"type": "integer"
						},
						"duration": {
							"type": "integer"
						},
						"multiAz": {
							"type": "boolean"
						},
						"productDescription": {
							"type": "keyword"
						},
						"offeringType": {
							"type": "keyword"
						},
						"state": {
							"type": "keyword"
						},
						"startTime": {
							"type": "date"
						},
						"recurringCharges": {
							"type": "nested",
							"properties": {
								"amount": {
									"type": "double"
								},
								"frequency": {
									"type": "keyword"
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