package dir

type diffState struct {
	dir1Files []*FileInfo
	dir2Files []*FileInfo
	patches   []*Patch
}

// Diff compares two versions of a directory over time. The first directory
// is the original or older version. The second directory is the new version.
func Diff(originalVersion *Directory, newerVersion *Directory) []*Patch {
	dir1Files := originalVersion.Flatten()
	dir2Files := newerVersion.Flatten()
	state := &diffState{
		dir1Files: dir1Files,
		dir2Files: dir2Files,
		patches:   []*Patch{},
	}

	state.computePatches()
	return state.patches
}

func (s *diffState) computePatches() {
	dir1Len := len(s.dir1Files)
	dir2Len := len(s.dir2Files)
	dir1Index, dir2Index := 0, 0

DIR_COMPARE_LOOP:
	for {
		switch {
		case dir1Index >= dir1Len && dir2Index >= dir2Len:
			break DIR_COMPARE_LOOP

		case dir1Index >= dir1Len:
			// We are at the end of the list for dir1Files any files in dir2Files are
			// not in dir1Files. We treat as a creation, and add the entries from dir2Files
			// to dir1Files
			patch := &Patch{
				File:    s.dir2Files[dir2Index],
				Type:    PatchCreate,
				ApplyTo: OriginalDirectory,
			}
			s.patches = append(s.patches, patch)

		case dir2Index >= dir2Len:
			// We are at the end of the list for dir2Files. Any files in dir1Files were added
			// and are not in dir2Files. We treat this as an add.

		case s.dir1Files[dir1Index].Path > s.dir2Files[dir2Index].Path:
			// There is a file in dir1Files that is not in dir2Files - add it

			// Decrement dir1Index because we are going to increment at the bottom
			// of the loop. Thus this decrement means we will be comparing this same dir1Files
			// entry again against another entry in dir2Files. We will keep doing this until
			// dir2Files catches up or we run out of dir2Files entries.
			dir1Index--

		case s.dir1Files[dir1Index].Path < s.dir2Files[dir2Index].Path:
			// There is a file in dir2Files that is not in dir1Files - delete it
			// from dir2Files.

			// Decrement dir2Index because we are going to increment at the bottom
			// of the loop. Thus this decrement means we will be comparing this same dir2Files
			// entry again against another entry in dir1Files. We will keep doing this until
			// dir1Files catches up or we run out of dir1Files entries.
			dir2Index--

		default:
			// The names would match - check stats to determine if there is a change.
			// if original is newer than newer then we need to update the new file. Else if
			// original is older than newer then we need to update older. Else we don't
			// need to do anything.
			dir1MTime := s.dir1Files[dir1Index].MTime
			dir2MTime := s.dir2Files[dir2Index].MTime
			switch {
			case dir1MTime.After(dir2MTime):
				// Update new file
			case dir1MTime.Before(dir2MTime):
				// Update original file
			default:
				// MTimes are the same, nothing to do.
			}
		}
		dir1Index++
		dir2Index++
	}
}
