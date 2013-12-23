# generate dummy images

simple command line utility that generates a set of white jpeg images.

## usage

			generate-dummy-images -count=10 -minWidth=140 -maxWidth=400 -minHeight=30 -maxHeight=500 -baseName="dummy_"

will generate a count of 10 white jpegs with random width between minWidth and maxWidth, random height between minHeight and maxHeight, and baseName of "dummy_". A number between 0 and count - 1 is appended to the baseName
