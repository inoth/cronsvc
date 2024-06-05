package registry

import (
	"context"
	"fmt"
	"sort"
)

type Register interface {
	Register(ctx context.Context, svc ServiceInstance) error
	Deregister(ctx context.Context, svc ServiceInstance) error
}

type ServiceInstance struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Version   string            `json:"version"`
	Metadata  map[string]string `json:"metadata"`
	Endpoints []string          `json:"endpoints"`
}

func (svc *ServiceInstance) String() string {
	return fmt.Sprintf("%s-%s-%s", svc.Name, svc.ID, svc.Version)
}

func (svc *ServiceInstance) Equal(o any) bool {
	if svc == nil && o == nil {
		return true
	}

	if svc == nil || o == nil {
		return false
	}

	t, ok := o.(*ServiceInstance)
	if !ok {
		return false
	}

	if len(svc.Endpoints) != len(t.Endpoints) {
		return false
	}

	sort.Strings(svc.Endpoints)
	sort.Strings(t.Endpoints)
	for j := 0; j < len(svc.Endpoints); j++ {
		if svc.Endpoints[j] != t.Endpoints[j] {
			return false
		}
	}

	if len(svc.Metadata) != len(t.Metadata) {
		return false
	}

	for k, v := range svc.Metadata {
		if v != t.Metadata[k] {
			return false
		}
	}

	return svc.ID == t.ID && svc.Name == t.Name && svc.Version == t.Version
}
