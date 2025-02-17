package zedgengo_test

import (
	"fmt"
	"testing"
	"zedgen/generated/zschema"
)

func Test(t *testing.T) {
	relationships := zschema.NewRelBuilder().Create().
		TeamRelation(
			zschema.TeamObj("123"),
			zschema.TeamRelTypeMember,
			zschema.UserObj("456"),
		).
		FolderRelation(
			zschema.FolderObj("789"),
			zschema.FolderRelTypeReader,
			zschema.TeamObj("456"),
		).
		DocumentRelation(
			zschema.DocumentObj("789"),
			zschema.DocumentRelTypeReader,
			zschema.TeamObj("123"),
		).
		DocumentRelation(
			zschema.DocumentObj("789"),
			zschema.DocumentRelTypeParentFolder,
			zschema.FolderObj("789"),
		).Build()
	fmt.Printf("relationships: %+v\n", relationships)
}
