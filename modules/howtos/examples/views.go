package main

// #tag::connect[]
import (
	"fmt"

	gocb "github.com/couchbase/gocb/v2"
)

func main() {
	opts := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			"Administrator",
			"password",
		},
	}
	cluster, err := gocb.Connect("localhost", opts)
	if err != nil {
		panic(err)
	}
	// #end::connect[]

	// #tag::bucket[]
	// get a bucket reference
	bucket := cluster.Bucket("bucket-name", &gocb.BucketOptions{})
	// #end::bucket[]

	// #tag::beerview[]
	viewResult, err := bucket.ViewQuery("beer", "by_name", &gocb.ViewOptions{
		StartKey: "A",
		Limit:    10,
	})
	if err != nil {
		panic(err)
	}
	// #end::beerview[]
	fmt.Println(viewResult)

	// #tag::landmarksview[]
	landmarksResult, err := bucket.ViewQuery("landmarks", "by_name", &gocb.ViewOptions{
		Key:       "<landmark-name>",
		Namespace: gocb.DevelopmentDesignDocumentNamespace,
	})
	if err != nil {
		panic(err)
	}
	// #end::landmarksview[]

	// #tag::results[]
	var landmarkRow gocb.ViewRow
	for landmarksResult.Next(&landmarkRow) {
		fmt.Printf("Document ID: %s\n", landmarkRow.ID)
		var key string
		err = landmarkRow.Key(&key)
		if err != nil {
			panic(err)
		}

		var landmark interface{}
		err = landmarkRow.Value(&landmark)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Landmark named %s has value %v\n", key, landmark)
	}

	err = landmarksResult.Close()
	if err != nil {
		panic(err)
	}
	// #end::results[]

}