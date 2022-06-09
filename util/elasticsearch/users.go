package elasticsearch

import "github.com/grokify/elastirad-go/models/es5"

func CreateUsersRequestBody(shardCount, replicaCount uint32, objectName string) es5.CreateIndexBody {
	if shardCount == 0 {
		shardCount = 1
	}
	if replicaCount == 0 {
		replicaCount = 1
	}
	body := es5.CreateIndexBody{
		Index: &es5.Index{
			NumberOfShards:   shardCount,
			NumberOfReplicas: replicaCount,
		},
		Mappings: map[string]es5.Mapping{
			objectName: {
				All: es5.All{Enabled: true},
				Properties: map[string]es5.Property{
					"about_me":      {Type: "text"},
					"display_name":  {Type: "text"},
					"link":          {Type: "text"},
					"location":      {Type: "text"},
					"profile_image": {Type: "text"},
					"reputation":    {Type: "long"},
					"website_url":   {Type: "text"},
				},
			},
		},
	}
	return body
}
