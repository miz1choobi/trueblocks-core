//go:build integration
// +build integration

package index

// func Test_EstablishValidFilename(t *testing.T) {
// 	// Such a range should not be listed in manifest, so the function will fail and we can investigate
// 	r := cache.FileRange{First: 0, Last: 1}
// 	_, err := EstablishIndexChunk("mainnet", r)
// 	if err == nil {
// 		t.Fatal("error expected")
// 	}
// }

// func IndexPathFromRange(fileRange cache.FileRange) (string, error) {
// 	tmp := cache.NewCachePath("mainnet", cache.Index_Final)
// 	path := tmp.GetFullPath(cache.FilenameFromRange(fileRange, ""))
// 	if file.FileExists(path) {
// 		return path, nil
// 	}
// 	return path, errors.New("file not found for " + path)
// }

// func Test_EstablishChunk(t *testing.T) {

// 	// TODO: This kills real data. Do we really want to do that? Remove this test.
// 	fileRange := cache.FileRange{First: 1, Last: 590501}
// 	path, err := IndexPathFromRange(fileRange)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if !file.Remove(path) {
// 		t.Error("could not remove file " + path)
// 	}

// 	if file.FileExists(path) {
// 		t.Error("remove file did not work for " + path)
// 	}

// 	_, err = EstablishIndexChunk("mainnet", fileRange)
// 	if err != nil {
// 		t.Error("EstablishIndexChunk returns error: " + err.Error())
// 	}

// 	if !file.FileExists(path) {
// 		t.Error("EstablishIndexChunk did not work for " + path)
// 	}
// }
