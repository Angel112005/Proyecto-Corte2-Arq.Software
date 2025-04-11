package repositoryapplication

type IRabbit interface {
	Publish(queue string, message []byte) error
}
