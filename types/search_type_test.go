package types

import (
	commontypes "github.com/a-castellano/music-manager-common-types/types"
	"testing"
)

func TestSelectRecordType(t *testing.T) {
	if SelectRecordType("Nonsense") != commontypes.Other {
		t.Errorf("'Nonsense' string should be type 'Other'.")
	}
	if SelectRecordType("Full-length") != commontypes.FullLength {
		t.Errorf("'Full-length' string should be type 'FullLength'.")
	}
	if SelectRecordType("Demo") != commontypes.Demo {
		t.Errorf("'Demo' string should be type 'Demo'.")
	}
	if SelectRecordType("EP") != commontypes.EP {
		t.Errorf("'EP' string should be type 'EP'.")
	}
	if SelectRecordType("Compilation") != commontypes.Compilation {
		t.Errorf("'Compilation' string should be type 'Compilation'.")
	}
	if SelectRecordType("Single") != commontypes.Single {
		t.Errorf("'Single' string should be type 'Single'.")
	}
	if SelectRecordType("Live album") != commontypes.Live {
		t.Errorf("'Live album' string should be type 'Live'.")
	}
	if SelectRecordType("Boxed set") != commontypes.BoxedSet {
		t.Errorf("'Boxed Set' string should be type 'BoxedSet'.")
	}
	if SelectRecordType("Video") != commontypes.Video {
		t.Errorf("'Video' string should be type 'Video'.")
	}
	if SelectRecordType("Split") != commontypes.Split {
		t.Errorf("'Split' string should be type 'Split'.")
	}
	if SelectRecordType("Other") != commontypes.Other {
		t.Errorf("'Other' string should be type 'Other'.")
	}
}
