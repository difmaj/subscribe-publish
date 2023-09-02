package repository

// Subscribe add a new subscription channel to a given queue.
func (r *Repository) Subscribe(queue string, callback func(message any)) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	ch := make(chan any)
	r.channels[queue] = append(r.channels[queue], ch)

	go func() {
		for msg := range ch {
			callback(msg)
		}
	}()
	return nil
}
