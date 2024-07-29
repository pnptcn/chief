package data

/*
Read the current (top) layer into p sequentially..
*/
func (artifact *Artifact) Read(p []byte) (n int, err error) {
	return
}

/*
Write the contents of p into a new Layer on the Artifact.
*/
func (artifact *Artifact) Write(p []byte) (n int, err error) {
	return
}

/*
Close the Artifact.
*/
func (artifact *Artifact) Close() error {
	return nil
}
