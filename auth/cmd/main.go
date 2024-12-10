package main

import (
	"auth/internal/config"
	"auth/internal/db"
	"auth/internal/pkg"
	"auth/internal/repository"
	"auth/internal/resolver"
	"log"
	"net/http"

	"github.com/graphql-go/handler"

	mygraphql "auth/internal/graphql"
)

var sandboxHTML = []byte(`
<!DOCTYPE html>
<html lang="en">
<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
<div id="sandbox" style="height:100vh; width:100vw;"></div>
<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
<script>
new window.EmbeddedSandbox({
  target: "#sandbox",
  initialEndpoint: "http://localhost:8080/graphql",
});
</script>
</body>

</html>`)

func main() {
	db.Init()
	config,err := config.NewConfig();
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	mongoClient, err := pkg.NewMongoClient(config.Mongo)
	if err != nil {
		panic(err)
	}

	repositories := repository.NewRepositoryLayer(mongoClient);
	resolvers := resolver.NewResolver(repositories);

    // GraphQL handler
	h := handler.New(&handler.Config{
		Schema:   mygraphql.NewSchema(resolvers),
		GraphiQL: true,
	})

	http.Handle("/graphql", h)

	http.HandleFunc("/sandbox", func(w http.ResponseWriter, r *http.Request) {
		// Set the content type to "text/html" and return the sandboxHTML content
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sandboxHTML))
	})

	http.ListenAndServe(":8080", nil)
}
