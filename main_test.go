package main

import "testing"

func TestExactExcludedFiles(t *testing.T) {

	filepaths := []string{"/tmp/File1", "/tmp/file2"}
	excluded_filenames := []string{"file1", "file2"}

	for _, fp := range filepaths {
		if !IsExcluded(fp, excluded_filenames) {
			t.Errorf("%s should be excluded from %v\n", fp, excluded_filenames)
		}
	}
}

func TestWildcardExcludedFiles(t *testing.T) {

	filepaths := []string{"/tmp/File1", "/tmp/file2"}
	excluded_filenames := []string{"file*"}

	for _, fp := range filepaths {
		if !IsExcluded(fp, excluded_filenames) {
			t.Errorf("%s should be excluded from %v\n", fp, excluded_filenames)
		}
	}
}
