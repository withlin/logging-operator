package plugins

// ElasticsearchOutput CRD name
const ElasticsearchOutput = "elasticsearch"

// ElasticsearchDefaultValues for Elasticsearch output plugin
var ElasticsearchDefaultValues = map[string]string{
	"bufferPath":      "/buffers/elasticsearch",
	"logLevel":        "info",
	"logstashFormat":  "true",
	"logstashPrefix":  "logstash",
	"chunkLimit":      "2M",
	"queueLimit":      "8",
	"timekey":         "1h",
	"timekey_wait":    "10m",
	"timekey_use_utc": "true",
}

// ElasticsearchTemplate for Elasticsearch output plugin
const ElasticsearchTemplate = `
<match {{ .pattern }}.**>
  @type elasticsearch
  @log_level {{ .logLevel }}
  include_tag_key true
  type_name fluentd
  host {{ .host }}
  port {{ .port }}
  scheme  {{ .schema }}
  logstash_format {{ .logstashFormat }}
  logstash_prefix {{ .logstashPrefix }}
  reconnect_on_error true
  <buffer tag, time>
    @type file
    path {{ .bufferPath }}
    timekey {{ .timekey }}
    timekey_wait {{ .timekey_wait }}
    timekey_use_utc {{ .timekey_use_utc }}
    flush_mode interval
    retry_type exponential_backoff
    flush_thread_count 2
    flush_interval 5s
    retry_forever
    retry_max_interval 30
    chunk_limit_size {{ .chunkLimit }}
    queue_limit_length {{ .queueLimit }}
    overflow_action block
  </buffer>
</match>`
