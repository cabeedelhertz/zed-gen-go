package zedgengo_test

import (
	"fmt"
	"testing"
	"zedgen/generated/zschema"
)

func Test(t *testing.T) {
	relationships := zschema.NewRelBuilder().Create().
		TeamRelationship(
			zschema.TeamObj("123"),
			zschema.TeamRelMember,
			zschema.UserObj("456"),
		).
		FolderRelationship(
			zschema.FolderObj("789"),
			zschema.FolderRelReader,
			zschema.TeamObj("456"),
		).
		DocumentRelationship(
			zschema.DocumentObj("789"),
			zschema.DocumentRelReader,
			zschema.TeamObj("123"),
		).
		DocumentRelationship(
			zschema.DocumentObj("789"),
			zschema.DocumentRelParentFolder,
			zschema.FolderObj("789"),
		).Build()
	fmt.Printf("relationships: %+v\n", relationships)
}
