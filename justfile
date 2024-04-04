openslo_spec_dir := "pkg/openslo"
openslo_schema_dir := "openslo/schemas"

gen-v1-spec:
	mkdir -p {{ openslo_spec_dir }}/v1
	quicktype --quiet --telemetry disable --field-tags yaml,json --omit-empty -l go -s schema -o {{ openslo_spec_dir }}/v1/openslo_v1.go --package openslo_v1 \
	    {{ openslo_schema_dir }}/v1/alertcondition.schema.json \
	    {{ openslo_schema_dir }}/v1/alertnotificationtarget.schema.json \
	    {{ openslo_schema_dir }}/v1/alertpolicy.schema.json \
	    {{ openslo_schema_dir }}/v1/budgetadjustment.schema.json \
	    {{ openslo_schema_dir }}/v1/datasource.schema.json \
	    {{ openslo_schema_dir }}/v1/service.schema.json \
	    {{ openslo_schema_dir }}/v1/sli.schema.json \
	    {{ openslo_schema_dir }}/v1/slo.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/alertcondition-spec.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/alertnotificationtarget-spec.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/alertpolicy-spec.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/budgetadjustment-spec.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/datasource-spec.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/description.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/duration-shorthand.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/general.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/metadata.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/metricsource.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/name.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/service-spec.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/sli-spec.schema.json \
	    -S {{ openslo_schema_dir }}/v1/parts/slo-spec.schema.json
