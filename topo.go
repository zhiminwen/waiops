package waiops

import (
	"fmt"
	"slices"

	"github.com/paulmach/orb/geojson"
	"github.com/zhiminwen/quote"
)

type Vertex struct {
	Operation   string          `json:"_operation,omitempty"`
	Name        string          `json:"name,omitempty"`
	UniqueId    string          `json:"uniqueId,omitempty"`
	EntityTypes []string        `json:"entityTypes,omitempty"`
	GeoLocation geojson.Feature `json:"geolocation,omitempty"` //Point, Polygon and Linestring.

	MatchTokens []string    `json:"matchTokens,omitempty"`
	Tags        []string    `json:"tags,omitempty"`
	References  []Reference `json:"_reference,omitempty"`

	Provider string `json:"_provider,omitempty"`

	Properties map[string]any `json:"properties,omitempty"` //extra properties
}
type Reference struct {
	FromUniqueId string `json:"_fromUniqueId,omitempty"` //vert doesn't have "from" exists yet
	ToUniqueId   string `json:"_toUniqueId,omitempty"`   //vert doesn't have to exists yet
	EdgeType     string `json:"_edgeType,omitempty"`
}
type VertexOpts func(*Vertex)

func WithName(name string) VertexOpts {
	return func(v *Vertex) {
		v.Name = name
	}
}

func WithUniqueId(uniqueId string) VertexOpts {
	return func(v *Vertex) {
		v.UniqueId = uniqueId
	}
}

func WithEntityTypes(entityTypes []string) VertexOpts {
	return func(v *Vertex) {
		v.EntityTypes = entityTypes
	}
}

func WithGeoLocation(geoLocation geojson.Feature) VertexOpts {
	return func(v *Vertex) {
		v.GeoLocation = geoLocation
	}
}

func WithMatchTokens(matchTokens []string) VertexOpts {
	return func(v *Vertex) {
		v.MatchTokens = matchTokens
	}
}

func WithTags(tags []string) VertexOpts {
	return func(v *Vertex) {
		v.Tags = tags
	}
}

func WithToReferences(toUniqueId, edgeType string) VertexOpts {
	return func(v *Vertex) {
		if !validateEdgeType(edgeType) {
			panic(fmt.Sprintf("invalid edge type: %s", edgeType))
		}
		v.References = append(v.References, Reference{ToUniqueId: toUniqueId, EdgeType: edgeType}) //v.References{
	}
}

func WithFromReferences(fromUniqueId, edgeType string) VertexOpts {
	return func(v *Vertex) {
		if !validateEdgeType(edgeType) {
			panic(fmt.Sprintf("invalid edge type: %s", edgeType))
		}
		v.References = append(v.References, Reference{FromUniqueId: fromUniqueId, EdgeType: edgeType}) //v.References{
	}
}

func validateEdgeType(edgeType string) bool {
	list := quote.Line(`
    contains
    federates
    members

    aliasOf
    assignedTo
    attachedTo
    classifies
    configures
    deployedTo
    exposes
    has
    implements
    locatedAt
    manages
    monitors
    movedTo
    origin
    owns
    rates
    resolvesTo
    realizes
    segregates
    uses

    accessedVia
    bindsTo
    communicatesWith
    connectedTo
    downlinkTo
    reachableVia
    receives
    routes
    routesVia
    loadBalances
    resolved
    resolves
    sends
    traverses
    uplinkTo

    dependsOn
    runsOn
  `)

	return slices.Contains(list, edgeType)
}

func WithOperation(operation string) VertexOpts {
	return func(v *Vertex) {
		v.Operation = operation
	}
}

func WithProvider(provider string) VertexOpts {
	return func(v *Vertex) {
		v.Provider = provider
	}
}

func NewVertex(name string, opts ...VertexOpts) *Vertex {
	v := &Vertex{
		Operation:   "InsertReplace",
		Name:        name,
		UniqueId:    name,
		MatchTokens: []string{name},
		Properties:  map[string]any{},
	}

	for _, opt := range opts {
		opt(v)
	}
	return v
}
