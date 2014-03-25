package dir

// Diff compares two versions of a directory over time. The first directory
// is the original or older version. The second directory is the new version.
func Diff(originalVersion *Directory, newerVersion *Directory) []*Patch {
	originalFiles := originalVersion.Flatten()
	newFiles := newerVersion.Flatten()
	originalFilesLen, newFilesLen := len(originalFiles), len(newFiles)
	originalFilesIndex, newFilesIndex := 0, 0
	patches := []*Patch{}

	// The checks need to become more complex. In order to determine if we should
	// add or delete a file, we need to look at a couple of things:
	// 1. Who has the latest set of updates? We assume here that newer wins.
	// 2. What are the events if we are monitoring events. Because the events can
	//    tell us what really happened.
	// 3.

DIR_COMPARE_LOOP:
	for {
		switch {
		case originalFilesIndex >= originalFilesLen && newFilesIndex >= newFilesLen:
			break DIR_COMPARE_LOOP

		case originalFilesIndex >= originalFilesLen:
			// We are at the end of the list for originalFiles any files in newFiles are
			// not in originalFiles. We treat as a creation, and add the entries from newFiles
			// to originalFiles
			patch := &Patch{
				File:    newFiles[newFilesIndex],
				Type:    PatchCreate,
				ApplyTo: OriginalDirectory,
			}
			patches = append(patches, patch)

		case newFilesIndex >= newFilesLen:
			// We are at the end of the list for newFiles. Any files in originalFiles were added
			// and are not in newFiles. We treat this as an add.

		case originalFiles[originalFilesIndex].Path > newFiles[newFilesIndex].Path:
			// There is a file in originalFiles that is not in newFiles - add it

			// Decrement originalFilesIndex because we are going to increment at the bottom
			// of the loop. Thus this decrement means we will be comparing this same originalFile
			// entry again against another entry in newFiles. We will keep doing this until
			// newFiles catches up or we run out of newFiles entries.
			originalFilesIndex--

		case originalFiles[originalFilesIndex].Path < newFiles[newFilesIndex].Path:
			// There is a file in newFiles that is not in originalFiles - delete it
			// from newFiles.

			// Decrement newFilesIndex because we are going to increment at the bottom
			// of the loop. Thus this decrement means we will be comparing this same newFile
			// entry again against another entry in originalFiles. We will keep doing this until
			// originalFiles catches up or we run out of originalFiles entries.
			newFilesIndex--

		default:
			// The names would match - check stats to determine if there is a change.
			// if original is newer than newer then we need to update the new file. Else if
			// original is older than newer then we need to update older. Else we don't
			// need to do anything.
			originalMTime := originalFiles[originalFilesIndex].MTime
			newMTime := newFiles[newFilesIndex].MTime
			switch {
			case originalMTime.After(newMTime):
				// Update new file
			case originalMTime.Before(newMTime):
				// Update original file
			default:
				// MTimes are the same, nothing to do.
			}
		}
		originalFilesIndex++
		newFilesIndex++
	}
	return nil
}
