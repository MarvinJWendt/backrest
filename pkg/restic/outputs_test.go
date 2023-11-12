package restic

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestReadBackupProgressEntries(t *testing.T) {
	t.Parallel()
	testInput := `{"message_type":"status","percent_done":0,"total_files":1,"total_bytes":15}
	{"message_type":"summary","files_new":0,"files_changed":0,"files_unmodified":166,"dirs_new":0,"dirs_changed":0,"dirs_unmodified":128,"data_blobs":0,"tree_blobs":0,"data_added":0,"total_files_processed":166,"total_bytes_processed":16754463,"total_duration":0.235433378,"snapshot_id":"bca1043e"}`

	b := bytes.NewBuffer([]byte(testInput))

	summary, err := readBackupProgressEntries(&exec.Cmd{}, b, func(event *BackupProgressEntry) {
		t.Logf("event: %v", event)
	})
	if err != nil {
		t.Fatalf("failed to read backup events: %v", err)
	}
	if summary == nil {
		t.Fatalf("wanted summary, got: nil")
	}
	if summary.TotalFilesProcessed != 166 {
		t.Errorf("wanted 166 files processed, got: %d", summary.TotalFilesProcessed)
	}
}


func TestReadLs(t *testing.T) {
	testInput := `{"time":"2023-11-10T19:14:17.053824063-08:00","tree":"3e2918b261948e69602ee9504b8f475bcc7cdc4dcec0b3f34ecdb014287d07b2","paths":["/resticui"],"hostname":"pop-os","username":"dontpanic","uid":1000,"gid":1000,"id":"db155169d788e6e432e320aedbdff5a54cc439653093bb56944a67682528aa52","short_id":"db155169","struct_type":"snapshot"}
	{"name":".git","type":"dir","path":"/.git","uid":1000,"gid":1000,"mode":2147484157,"mtime":"2023-11-10T18:32:38.156599473-08:00","atime":"2023-11-10T18:32:38.156599473-08:00","ctime":"2023-11-10T18:32:38.156599473-08:00","struct_type":"node"}
	{"name":".gitignore","type":"file","path":"/.gitignore","uid":1000,"gid":1000,"size":22,"mode":436,"mtime":"2023-11-10T00:41:26.611346634-08:00","atime":"2023-11-10T00:41:26.611346634-08:00","ctime":"2023-11-10T00:41:26.611346634-08:00","struct_type":"node"}
	{"name":"README.md","type":"file","path":"/README.md","uid":1000,"gid":1000,"size":762,"mode":436,"mtime":"2023-11-10T00:59:06.842538768-08:00","atime":"2023-11-10T00:59:06.842538768-08:00","ctime":"2023-11-10T00:59:06.842538768-08:00","struct_type":"node"}`

	b := bytes.NewBuffer([]byte(testInput))

	snapshot, entries, err := readLs(b)
	if err != nil {
		t.Fatalf("failed to read ls output: %v", err)
	}
	if snapshot == nil {
		t.Fatalf("wanted snapshot, got: nil")
	}
	if len(entries) != 3 {
		t.Errorf("wanted 3 entries, got: %d", len(entries))
	}
}

func TestSnapshotToProto(t *testing.T) {
	snapshot := &Snapshot{
		Id: "db155169d788e6e432e320aedbdff5a54cc439653093bb56944a67682528aa52",
		Time: "2023-11-10T19:14:17.053824063-08:00",
		Tree: "3e2918b261948e69602ee9504b8f475bcc7cdc4dcec0b3f34ecdb014287d07b2",
		Paths: []string{"/resticui"},
		Hostname: "pop-os",
		Username: "dontpanic",
		Tags: []string{},
		Parent: "",
	}

	proto := snapshot.ToProto()

	if proto.Id != snapshot.Id {
		t.Errorf("wanted id %q, got: %q", snapshot.Id, proto.Id)
	}
	if proto.Tree != snapshot.Tree {
		t.Errorf("wanted tree %q, got: %q", snapshot.Tree, proto.Tree)
	}
	if proto.Hostname != snapshot.Hostname {
		t.Errorf("wanted hostname %q, got: %q", snapshot.Hostname, proto.Hostname)
	}
	if proto.Username != snapshot.Username {
		t.Errorf("wanted username %q, got: %q", snapshot.Username, proto.Username)
	}
	if len(proto.Tags) != len(snapshot.Tags) {
		t.Errorf("wanted %d tags, got: %d", len(snapshot.Tags), len(proto.Tags))
	}
	if proto.Parent != snapshot.Parent {
		t.Errorf("wanted parent %q, got: %q", snapshot.Parent, proto.Parent)
	}
	if proto.UnixTimeMs != 1699672457053 {
		t.Errorf("wanted unix time %d, got: %d", 1699672457053, proto.UnixTimeMs)
	}
}