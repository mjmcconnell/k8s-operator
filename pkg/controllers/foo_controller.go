/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	testv1alpha1 "github.com/mjmcconnell/k8s-operator/pkg/api/v1alpha1"
)

// FooReconciler reconciles a Foo object
type FooReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *FooReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// GetOrSet related resources into the state manager
	// If parent resource (CRD), is not in a ready state, loop through all related
	// resources to check their state.

	// If a child resource is not in a ready state, check its readiness state, and report on its status.
	// Ensure resource matches expected state, and if not update resource.
	// If child status is not ready, re-trigger reconcile loop with a backoff.

	// If parent resource is in a ready state, query all child resources, to ensure their state matches expected.
	// If child is found drifted, then set both child and parent state to not ready, and trigger a new reconcile.

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FooReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testv1alpha1.Foo{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Secret{}).
		Complete(r)
}
