package main

import "testing"

func TestIsInteresting(t *testing.T) {
	var context = runconf{
		FilterOut: []string{"Windows"},
		LookFor:   []string{".jpg", ".png", ".jpeg"},
		rootPath:  "/Volumes/Acer/demo/Desktop",
	}

	got := isInteresting(".pdf", context)
	if got {
		t.Errorf("Extension .pdf is not in %s but was recognised wrong", context.LookFor)
	}

	got = isInteresting(".JPEG", context)
	if !got {
		t.Errorf("Extension .JPEG is in %s but was recognised wrong", context.LookFor)
	}

	got = isInteresting(".png", context)
	if !got {
		t.Errorf("Extension .png is in %s but was recognised wrong", context.LookFor)
	}
}

func TestFileIsToBeFilteredOut(t *testing.T) {
	var context = runconf{
		FilterOut: []string{"Windows", "Program Files"},
		LookFor:   []string{".jpg", ".png", ".jpeg"},
		rootPath:  "/Volumes/Acer/demo/Desktop",
	}

	got := fileIsToBeFilteredOut("/Acer/Windows/System32/a123/a.jpg", context)
	if !got {
		t.Errorf("File in Windows/ should have been filtered out, %s", context.FilterOut)
	}

	got = fileIsToBeFilteredOut("C://Program Files/Google Chrome/assets/icons/tulip.png", context)
	if !got {
		t.Errorf("File in Program Files/ should have been filtered out, %s", context.FilterOut)
	}

	got = fileIsToBeFilteredOut("C://Users/jack/Desktop/mom.JPEG", context)
	if got {
		t.Errorf("File should have been copied but was filtered out, %s", context.FilterOut)
	}

	got = fileIsToBeFilteredOut("mom.JPEG", context)
	if got {
		t.Errorf("File should have been copied but was filtered out, %s", context.FilterOut)
	}

}

func TestStringContainsAnyOfSubstrings(t *testing.T) {
	got := stringContainsAnyOfSubstrings("C://Users/jack/Desktop/mom.JPEG", []string{"Desktop"})
	if !got {
		t.Error("Failed to spot Desktop in C://Users/jack/Desktop/mom.JPEG")
	}

	got = stringContainsAnyOfSubstrings("/Acer/Windows/System32/a123/a.jpg", []string{"DS_Store"})
	if got {
		t.Error("Incorrectly spotted DS_Store in /Acer/Windows/System32/a123/a.jpg")
	}

	got = stringContainsAnyOfSubstrings("/Google Chrome/assets/icons/tulip.png", []string{"Google Chrome"})
	if !got {
		t.Error("Failed to spot Google Chrome in /Google Chrome/assets/icons/tulip.png")
	}
}

func TestStrIsAnyOf(t *testing.T) {
	got := strIsAnyOf("cat", []string{"dog", "cat", "parrot"})
	if !got {
		t.Error("Failed to spot a cat in [dog, cat, parrot]")
	}

	got = strIsAnyOf("animal", []string{"dog", "cat", "parrot"})
	if got {
		t.Error("Got a little bit too clever here")
	}

	got = strIsAnyOf("Google Chrome", []string{"Google Chrome", "Firefox"})
	if !got {
		t.Error("Failed to spot Google Chrome in [Google Chrome, Firefox]")
	}
}
