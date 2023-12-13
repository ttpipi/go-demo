package main

import "context"

func main() {
	handlers := HandleWithMiddleware(NewHandler(), NewA(), NewB())
	handlers.Handle(context.Background())
}
