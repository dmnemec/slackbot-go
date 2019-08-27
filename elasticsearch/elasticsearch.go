package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	port = "9200"
)

// RestoreIndices sends a request to restore elasticsearch snapshots to an ES cluster
func RestoreIndices(esURL, env, fbVersion string, days ...string) (err error) {
	//Validate days
	//TODO verify dates to conform to YYYY.MM.DD format

	//Find snapshot with all days
	resp, err := http.Get(esURL + ":" + port + "/_snapshot/" + env + "/_all")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(bb))
	//TODO check resp.StatusCode

	//Get snapshot name (currator)
	snapshots := SnapshotsResponseJSON{}
	if err := json.Unmarshal(bb, &snapshots); err != nil {
		return err
	}

	var use bool
	var shot string
Snapshots:
	for _, v := range snapshots.Snapshots {
		use = true
	Days:
		for _, d := range days {
			for _, i := range v.Indices {

				use = false
				fmt.Printf("d: %s, i: %s, snapshot: %s\n", d, i, v.Snapshot)
				if strings.HasSuffix(i, d) {
					use = true
					continue Days
				}
			}
			if !use {
				continue Snapshots
			}
		}
		if use {
			shot = v.Snapshot
			break Snapshots
		}
	}
	//Build json with days
	for i, v := range days {
		days[i] = "filebeat-" + fbVersion + "-" + v
	}
	data, err := json.Marshal(RestoreJSON{Indices: strings.Join(days, ","),
		IndexSettings:       IndexSettings{IndexNumberOfReplicas: 0},
		IgnoreIndexSettings: []string{"index.refresh_interval"},
		IgnoreUnavailable:   true,
		RenamePattern:       "(.+)",
		RenameReplacement:   "restored-$1"})
	if err != nil {
		return err
	}
	//POST url with with json
	fmt.Println("snapshot = " + shot)
	fmt.Printf("%s\n", data)
	fmt.Println("restore request url = " + esURL + ":" + port + "/_snapshot/" + env + "/" + shot + "/_restore?wait_for_completion=false")
	resp2, err := http.Post(esURL+":"+port+"/_snapshot/"+env+"/"+shot+"/_restore?wait_for_completion=false", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp2.Body.Close()
	bb, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(bb))
	fmt.Println(resp2.StatusCode)
	for i, v := range days {
		days[i] = "restored-" + v
	}
	fmt.Println("recovery progress url: " + esURL + ":" + port + "/_cat/recovery/" + strings.Join(days, ",") + "?human")
	return err
}
