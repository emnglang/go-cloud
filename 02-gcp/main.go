
// Sample storage-quickstart 
// creates a Google Cloud Storage bucket

package main

import (
        "fmt"
        "log"

        "cloud.google.com/go/storage"
        "golang.org/x/net/context"
)

func main() {
        ctx := context.Background()

        projectID := "YOUR_PROJECT_ID"

        client, err := storage.NewClient(ctx)
        if err != nil {
                log.Fatalf("Failed to create client: %v", err)
        }

        bucketName := "my-new-bucket"

        bucket := client.Bucket(bucketName)

        if err := bucket.Create(ctx, projectID, nil); err != nil {
                log.Fatalf("Failed to create bucket: %v", err)
        }

        fmt.Printf("Bucket %v created.\n", bucketName)
}