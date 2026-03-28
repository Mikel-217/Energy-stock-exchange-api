package database

type ReaderClientBuilder[T any] struct {
	client ReadDataClient[T]
}

func CreateNewBuilder[T any]() *ReaderClientBuilder[T] {
	return &ReaderClientBuilder[T]{
		client: ReadDataClient[T]{
			Query: "Default",
		},
	}
}

// Adds the given query
func (r *ReaderClientBuilder[T]) AddQuery(query string) *ReaderClientBuilder[T] {
	r.client.Query = query
	return r
}

// Adds the given parameter
func (r *ReaderClientBuilder[T]) AddQueryParams(queryParams []any) *ReaderClientBuilder[T] {
	r.client.QueryParams = queryParams
	return r
}

// Builds our client
func (r *ReaderClientBuilder[T]) Build() *ReadDataClient[T] {
	return &r.client
}
