package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	// Import generated proto package
	proto "github.com/imnerocode/parser-service/proto/generated"
	"github.com/oakmound/ofbx"
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

// ParseFBX parses an FBX file and returns a Model structure
func ParseFBX(filePath string) (*proto.Model, error) {
	// Open the FBX file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open FBX file: %w", err)
	}
	defer file.Close()

	// Parse the FBX file
	scene, err := ofbx.Load(file)
	if err != nil {
		return nil, fmt.Errorf("failed to parse FBX file: %w", err)
	}

	// Initialize the Model structure
	model := &proto.Model{
		Vertices: []*proto.Vertex{},
		Faces:    []*proto.Face{},
	}

	// Iterate over the meshes in the FBX scene
	for _, mesh := range scene.Meshes {
		// Extract vertices
		vertices := mesh.Geometry.Vertices
		for _, vertex := range vertices {
			model.Vertices = append(model.Vertices, &proto.Vertex{
				X: float32(vertex[0]),
				Y: float32(vertex[1]),
				Z: float32(vertex[2]),
			})
		}

		// Extract faces
		for _, face := range mesh.Geometry.Faces {
			faceIndices := make([]int32, len(face))
			for i, idx := range face {
				faceIndices[i] = int32(idx)
			}
			model.Faces = append(model.Faces, &proto.Face{
				VertexIndices: faceIndices,
			})
		}
	}

	return model, nil
}
