package whereami

import "testing"

func TestShouldOutputCorrectData(t *testing.T) {
	out := WhereAmIData(WhereAmIDataInput{
		DepthList:          []int{1},
		FullFilePath:       false,
		FullFunctionOutput: false,
	})

	const expectedLineNumber = 6
	const expectedFileName = "whereami_test.go"
	const expectedFunction = "whereami.TestShouldOutputCorrectData"

	if out.File != expectedFileName {
		t.Fatalf("File is expected to be <%s>, but is <%s>!", expectedFileName, out.File)
	}

	if out.Function != expectedFunction {
		t.Fatalf("Function is expected to be <%s>, but is <%s>!", expectedFunction, out.Function)
	}

	if out.LineNumber != expectedLineNumber {
		t.Fatalf("LineNumber is expected to be <%d>, but is <%d>!", expectedLineNumber, out.LineNumber)
	}
}

func TestShouldChopPath(t *testing.T) {
	const testPath = "github.com/testreopsitory/golang"
	const expectedChoppedPath = "golang"

	choppedPath := chopPath(testPath)

	if choppedPath != expectedChoppedPath {
		t.Fatalf("choppedPath is expected to be <%s>, but is <%s>!", expectedChoppedPath, choppedPath)
	}
}
