package webhook

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch;create;update;patch
// +kubebuilder:webhook:path=/validate-v1-namespace,mutating=false,failurePolicy=ignore,groups="",resources=namespaces,verbs=create;update;delete,versions=v1,name=namespacelogger.dana.io,admissionReviewVersions=v1,sideEffects=None

type NamespaceLoggerValidator struct{}

// Handle logs requests for namespaces.
func (r *NamespaceLoggerValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	logger := log.FromContext(ctx)
	logger.Info("Details:", "Namespace", req.Name, "Operation", req.Operation, "User", req.UserInfo.Username)

	return admission.Allowed("Namespace event logged")
}
