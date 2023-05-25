package reader

import (
	"fmt"

	"github.com/observiq/bolt-explorer/router"
	bolt "go.etcd.io/bbolt"
)

var ErrBucketNil = fmt.Errorf("bucket is nil")

// ParseDB takes in a bbolt database filepath and returns an array of all
// routes in the database.
func ParseDB(db *bolt.DB) ([]*router.Route, error) {
	routes := []*router.Route{}

	// Get the top level buckets
	err := db.View(func(tx *bolt.Tx) error {
		tx.ForEach(func(bucketKey []byte, bucket *bolt.Bucket) error {
			path := makePath(make([]byte, 0), bucketKey)
			route := router.NewRoute(path, "", nil)
			bucketRoutes := getRoutesFromBucket(bucket, path)
			routes = append(routes, route)
			routes = append(routes, bucketRoutes...)
			return nil
		})

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("ParseDB: %w", err)
	}
	return routes, err
}

func getRoutesFromBucket(bucket *bolt.Bucket, parentKey string) []*router.Route {
	routes := []*router.Route{}

	bucket.ForEach(func(childKey, v []byte) error {
		route := router.NewRoute(string(childKey), parentKey, v)

		childBucket := bucket.Bucket(childKey)
		if childBucket != nil {
			childRoutes := getRoutesFromBucket(childBucket, route.Path)

			for _, childRoute := range childRoutes {
				if childRoute.Parent == route.Path {
					route.Children = append(route.Children, childRoute.Path)
				}

			}
			routes = append(routes, childRoutes...)
		} else {
			route.Children = append(route.Children, string(childKey))
		}

		routes = append(routes, route)
		return nil
	})

	return routes
}

func makePath(parent, child []byte) string {
	if len(parent) == 0 {
		return string(child)
	}

	return fmt.Sprintf("%s.%s", parent, child)
}
