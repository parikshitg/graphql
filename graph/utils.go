package graph

func (r *mutationResolver) GetNextID() int {
	id := 0
	for k := range r.DB {
		if k > id {
			id = k
		}
	}
	return id + 1
}
