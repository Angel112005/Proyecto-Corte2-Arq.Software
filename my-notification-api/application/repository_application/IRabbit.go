// application/repository_application/IRabbit.go
package repositoryapplication

type IRabbit interface {
	Publish(queue string, message []byte) error
}
