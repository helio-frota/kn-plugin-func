package knative

import (
	"context"
	"fmt"
	"time"

	"knative.dev/kn-plugin-func/k8s"
)

const RemoveTimeout = 120 * time.Second

func NewRemover(namespaceOverride string) *Remover {
	return &Remover{
		Namespace: namespaceOverride,
	}
}

type Remover struct {
	Namespace string
	Verbose   bool
}

func (remover *Remover) Remove(ctx context.Context, name string) (err error) {
	if remover.Namespace == "" {
		remover.Namespace, err = k8s.GetNamespace(remover.Namespace)
		if err != nil {
			return err
		}
	}

	client, err := NewServingClient(remover.Namespace)
	if err != nil {
		return
	}

	err = client.DeleteService(ctx, name, RemoveTimeout)
	if err != nil {
		err = fmt.Errorf("knative remover failed to delete the service: %v", err)
	}

	return
}
