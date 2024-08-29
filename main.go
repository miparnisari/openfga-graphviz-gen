package main

import (
	"flag"
	"io"
	"log"
	"os"

	graph "github.com/openfga/language/pkg/go/graph"
	language "github.com/openfga/language/pkg/go/transformer"
)

func main() {
	modelPathFlag := flag.String("model-path", "", "the file path for the OpenFGA model (in DSL format)")
	outputPathFlag := flag.String("output-path", "", "the file path for the output graph (default to stdout)")

	flag.Parse()

	bytes, err := os.ReadFile(*modelPathFlag)
	if err != nil {
		log.Fatalf("failed to read model file: %v", err)
	}

	model := language.MustTransformDSLToProto(string(bytes))
	graph, err := graph.NewAuthorizationModelGraph(model)

	result := graph.GetDOT()

	var writer io.Writer
	if *outputPathFlag != "" && *outputPathFlag != "-" {
		writer, _ = os.Create(*outputPathFlag)
	} else {
		writer = os.Stdout
	}

	_, err = writer.Write([]byte(result))
	if err != nil {
		log.Fatalf("failed to render graph: %v", err)
	}
}
