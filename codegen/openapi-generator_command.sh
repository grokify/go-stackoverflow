java -jar openapi-generator-cli.jar generate -c openapi-generator_config.json -i openapi_spec.yaml -g go -o client
echo "\n\nfunc (apiClient *APIClient) HTTPClient() *http.Client { return apiClient.cfg.HTTPClient }" >> client/client.go
gofmt -s -w client/*.go
