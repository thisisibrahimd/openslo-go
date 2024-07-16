openslo_spec_dir := "pkg/openslo"
openapi_schema_url := "https://raw.githubusercontent.com/thisisibrahimd/OpenSLO/feature/add-budget-adjustment-spec/schemas/v1/schema.yaml"

gen: gen-v1-spec

gen-v1-spec:
	GO_POST_PROCESS_FILE="/usr/local/bin/gofmt -w" openapi-generator \
		generate \
		-i {{ openapi_schema_url }} \
		-g go \
		-o {{ openslo_spec_dir }}/v1 \
		--inline-schema-name-mappings AlertPolicyCondition_inner=AlertPolicyCondition,AlertPolicyNotificationTarget_inner=AlertPolicyNotificationTarget \
		--global-property models,modelDocs=false \
		--package-name openslo_v1

