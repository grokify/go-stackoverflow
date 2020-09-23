package elasticsearch

import v5 "github.com/grokify/elastirad-go/models/v5"

func CreateUsersRequestBody(shardCount, replicaCount uint32, objectName string) v5.CreateIndexBody {
	if shardCount == 0 {
		shardCount = 1
	}
	if replicaCount == 0 {
		replicaCount = 1
	}
	body := v5.CreateIndexBody{
		Index: &v5.Index{
			NumberOfShards:   shardCount,
			NumberOfReplicas: replicaCount,
		},
		Mappings: map[string]v5.Mapping{
			objectName: {
				All: v5.All{Enabled: true},
				Properties: map[string]v5.Property{
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
