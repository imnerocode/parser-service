package parser

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	// Import generated proto package
	proto "github.com/imnerocode/parser-service/proto/generated"
)

// ParseOBJ parses an OBJ file and returns a Model structure
func ParseOBJ(filePath string) (*proto.Model, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	model := &proto.Model{
		Vertices: []*proto.Vertex{},
		Faces:    []*proto.Face{},
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue // Ignore comments and empty lines
		}

		parts := strings.Fields(line)
		switch parts[0] {
		case "v": // Vertex definition
			x, _ := strconv.ParseFloat(parts[1], 64)
			y, _ := strconv.ParseFloat(parts[2], 64)
			z, _ := strconv.ParseFloat(parts[3], 64)
			model.Vertices = append(model.Vertices, &proto.Vertex{X: float32(x), Y: float32(y), Z: float32(z)})

		case "f": // Face definition
			var indices []int32
			for _, part := range parts[1:] {
				vertexIndex, _ := strconv.Atoi(strings.Split(part, "/")[0])
				indices = append(indices, int32(vertexIndex-1)) // OBJ is 1-based
			}
			model.Faces = append(model.Faces, &proto.Face{VertexIndices: indices})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return model, nil
}
