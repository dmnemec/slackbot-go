package elasticsearch

import "time"

type RestoreJSON struct {
	Indices             string        `json:"indices"`
	IndexSettings       IndexSettings `json:"index_settings"`
	IgnoreIndexSettings []string      `json:"ignore_index_settings"`
	IgnoreUnavailable   bool          `json:"ignore_unavailable"`
	RenamePattern       string        `json:"rename_pattern"`
	RenameReplacement   string        `json:"rename_replacement"`
}

type IndexSettings struct {
	IndexNumberOfReplicas int `json:"index.number_of_replicas"`
}

type SnapshotsResponseJSON struct {
	Snapshots []struct {
		Snapshot           string        `json:"snapshot"`
		UUID               string        `json:"uuid"`
		VersionID          int           `json:"version_id"`
		Version            string        `json:"version"`
		Indices            []string      `json:"indices"`
		IncludeGlobalState bool          `json:"include_global_state"`
		State              string        `json:"state"`
		StartTime          time.Time     `json:"start_time"`
		StartTimeInMillis  int64         `json:"start_time_in_millis"`
		EndTime            time.Time     `json:"end_time"`
		EndTimeInMillis    int64         `json:"end_time_in_millis"`
		DurationInMillis   int           `json:"duration_in_millis"`
		Failures           []interface{} `json:"failures"`
		Shards             struct {
			Total      int `json:"total"`
			Failed     int `json:"failed"`
			Successful int `json:"successful"`
		} `json:"shards"`
	} `json:"snapshots"`
}
