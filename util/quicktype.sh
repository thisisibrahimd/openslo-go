#!/usr/bin/env fish

argparse h/help d/schema-dir o/output-dir -- $argv
or return

# If -h or --help is given, we print a little help text and return
if set -ql _flag_help
    echo "quicktype.sh [-h|--help] [-d|--schema-dir] [-o|--output-dir] [ARGUMENT ...]"
    return 0
end

quicktype --quiet --telemetry disable --just-types -l go -s schema -o openslo_v1.go --package openslo_v1 \
    schemas/v1/alertcondition.schema.json \
    schemas/v1/alertnotificationtarget.schema.json \
    schemas/v1/alertpolicy.schema.json \
    schemas/v1/budgetadjustment.schema.json \
    schemas/v1/datasource.schema.json \
    schemas/v1/service.schema.json \
    schemas/v1/sli.schema.json \
    schemas/v1/slo.schema.json \
    -S schemas/v1/parts/alertcondition-spec.schema.json \
    -S schemas/v1/parts/alertnotificationtarget-spec.schema.json \
    -S schemas/v1/parts/alertpolicy-spec.schema.json \
    -S schemas/v1/parts/budgetadjustment-spec.schema.json \
    -S schemas/v1/parts/datasource-spec.schema.json \
    -S schemas/v1/parts/description.schema.json \
    -S schemas/v1/parts/duration-shorthand.schema.json \
    -S schemas/v1/parts/general.schema.json \
    -S schemas/v1/parts/metadata.schema.json \
    -S schemas/v1/parts/metricsource.schema.json \
    -S schemas/v1/parts/name.schema.json \
    -S schemas/v1/parts/service-spec.schema.json \
    -S schemas/v1/parts/sli-spec.schema.json \
    -S schemas/v1/parts/slo-spec.schema.json
