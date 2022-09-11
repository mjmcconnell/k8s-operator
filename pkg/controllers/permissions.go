/*
Copyright 2021 white.space.

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

/*
This file is used to specify all required permissions for the  Operator

Each kubebuilder definition result in a matching rule being defined within config/rbac/role.yaml

Any updates to this file, sould be sync, with `make sync.bundle`, to update the config files.
*/

package controllers

// Foo
//+kubebuilder:rbac:groups=test.mjmcconnell.com,resources=foos,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=test.mjmcconnell.com,resources=foos/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=test.mjmcconnell.com,resources=foos/finalizers,verbs=update

// Native resources generated as part of the above CRD's
//+kubebuilder:rbac:groups="",namespace=mjm,resources=configmaps;secrets;serviceaccounts;services;endpoints,verbs=*
//+kubebuilder:rbac:groups=apps,namespace=mjm,resources=deployments;statefulsets,verbs=*
//+kubebuilder:rbac:groups=batch,namespace=mjm,resources=cronjobs;jobs,verbs=*
//+kubebuilder:rbac:groups=networking.k8s.io,namespace=mjm,resources=networkpolicies;ingresses,verbs=*
//+kubebuilder:rbac:groups=rbac.authorization.k8s.io,namespace=mjm,resources=roles;rolebindings,verbs=*
//+kubebuilder:rbac:groups=autoscaling,namespace=mjm,resources=horizontalpodautoscalers,verbs=*

// Prometheus
//+kubebuilder:rbac:groups=monitoring.coreos.com,namespace=mjm,resources=servicemonitors;prometheusrules,verbs=*

// Openshift
//+kubebuilder:rbac:groups=route.openshift.io,namespace=mjm,resources=routes;routes/custom-host,verbs=*
//+kubebuilder:rbac:groups=security.openshift.io,namespace=mjm,resources=securitycontextconstraints,resourceNames=restricted,verbs=get

// GCP
//+kubebuilder:rbac:groups=monitoring.googleapis.com,namespace=mjm,resources=podmonitorings,verbs=*
