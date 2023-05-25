package reader

import (
	"fmt"

	"github.com/observiq/bolt-explorer/router"
	"go.etcd.io/bbolt"
	bolt "go.etcd.io/bbolt"
)

var ErrBucketNil = fmt.Errorf("bucket is nil")

// ParseDB takes in a bbolt database filepath and returns an array of all
// routes in the database.
func ParseDB(db *bbolt.DB) ([]*router.Route, error) {
	routes := []*router.Route{}

	// Get the top level buckets
	err := db.View(func(tx *bolt.Tx) error {
		err := tx.Cursor().Bucket().ForEach(func(k, v []byte) error {
			path := makePath(make([]byte, 0), k)
			route := router.NewRoute(path, "", nil)

			// Add the children
			tx.Bucket(k).ForEach(func(k, v []byte) error {
				// Add children
				subPath := makePath([]byte{}, k)
				route.Children = append(route.Children, subPath)
				return nil
			})

			routes = append(routes, route)

			foundKeys, err := getRoutes(tx, k)
			if err != nil {
				return err
			}
			routes = append(routes, foundKeys...)
			return nil
		})

		return err
	})

	if err != nil {
		return nil, fmt.Errorf("ParseDB: %w", err)
	}

	return routes, err
}

func getRoutes(tx *bolt.Tx, bucketName []byte) ([]*router.Route, error) {
	routes := []*router.Route{}
	bucket := tx.Bucket(bucketName)
	if bucket == nil {
		return routes, nil
	}

	err := bucket.ForEach(func(key, value []byte) error {
		// Add this key to the keys array
		path := makePath(bucketName, key)
		route := router.NewRoute(path, string(bucketName), value)

		bucket.Cursor().Bucket().ForEachBucket(func(k []byte) error {
			// Add children
			subPath := makePath(key, k)
			route.Children = append(route.Children, subPath)
			return nil
		})

		routes = append(routes, route)

		// Add all of its childrens keys to the keys array
		childRoutes, err := getRoutes(tx, key)
		if err != nil {
			return err
		}
		routes = append(routes, childRoutes...)
		return nil
	})

	return routes, err
}

func makePath(parent, child []byte) string {
	if len(parent) == 0 {
		return string(child)
	}

	return fmt.Sprintf("%s.%s", parent, child)
}
