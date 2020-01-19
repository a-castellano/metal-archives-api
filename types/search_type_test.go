package types

import (
	"testing"
)

func TestSelectRecordType(t *testing.T) {
	if SelectRecordType("Nonsense") != Other {
		t.Errorf("'Nonsense' string should be type 'Other'.")
	}
	if SelectRecordType("Full-length") != FullLength {
		t.Errorf("'Full-length' string should be type 'FullLength'.")
	}
	if SelectRecordType("Demo") != Demo {
		t.Errorf("'Demo' string should be type 'Demo'.")
	}
	if SelectRecordType("EP") != EP {
		t.Errorf("'EP' string should be type 'EP'.")
	}
	if SelectRecordType("Compilation") != Compilation {
		t.Errorf("'Compilation' string should be type 'Compilation'.")
	}
	if SelectRecordType("Single") != Single {
		t.Errorf("'Single' string should be type 'Single'.")
	}
	if SelectRecordType("Live album") != Live {
		t.Errorf("'Live album' string should be type 'Live'.")
	}
	if SelectRecordType("Boxed set") != BoxedSet {
		t.Errorf("'Boxed Set' string should be type 'BoxedSet'.")
	}
	if SelectRecordType("Video") != Video {
		t.Errorf("'Video' string should be type 'Video'.")
	}
	if SelectRecordType("Split") != Split {
		t.Errorf("'Split' string should be type 'Split'.")
	}
	if SelectRecordType("Other") != Other {
		t.Errorf("'Other' string should be type 'Other'.")
	}
}
