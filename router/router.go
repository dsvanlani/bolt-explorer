package router

import (
	"sort"

	"go.etcd.io/bbolt"
)

type PathValueFunc func() string

func getPathValueFunc(db *bbolt.DB, path string) PathValueFunc {
	return func() string {
		return path
	}
}

type Route struct {
	Path     string
	Parent   string
	Children []string
	Value    []byte
}

func NewRoute(path string, parent string, value []byte) *Route {
	return &Route{
		Path:     path,
		Parent:   parent,
		Children: make([]string, 0),
		Value:    value,
	}
}

type Router struct {
	Paths    map[string]*Route
	Location string
	DB       *bbolt.DB
}

func NewRouter(routes []*Route, db *bbolt.DB) *Router {
	paths := make(map[string]*Route)
	for _, route := range routes {
		paths[route.Path] = route
	}

	return &Router{
		Paths:    paths,
		Location: "",
		DB:       db,
	}
}

func (r *Router) GetLocation() string {
	return r.Location
}

func (r *Router) SetLocation(location string) {

	r.Location = location
}

// GetPathsForLocation returns the paths for the current location.
// If the location is not a bucket, it returns the keys
// for the parent bucket.
func (r *Router) GetPathsForLocation() []string {
	keys := []string{}

	for path := range r.Paths {
		if r.Paths[path].Parent == r.Location {
			keys = append(keys, path)
		}
	}

	// alphabetize the keys
	sort.Strings(keys)

	return keys
}

func (r *Router) GoUpOneLevel() {
	if r.Location == "" {
		return
	}

	route := r.Paths[r.Location]

	r.SetLocation(route.Parent)
}

func (r *Route) IsBucket() bool {
	return r.Value == nil
}
