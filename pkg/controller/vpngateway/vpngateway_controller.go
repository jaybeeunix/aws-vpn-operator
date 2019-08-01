package vpngateway

import (
	"context"
	"fmt"

	awsvpnv1alpha1 "github.com/jaybeeunix/aws-vpn-operator/pkg/apis/awsvpn/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	// "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"sigs.k8s.io/yaml"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	configv1 "github.com/openshift/api/config/v1"
)

var (
	log = logf.Log.WithName("controller_vpngateway")
)

const (
	awsCredsSecretName      = "aws-vpn-operator-credentials"
	awsCredsSecretIDKey     = "aws_access_key_id"
	awsCredsSecretAccessKey = "aws_secret_access_key"
)

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new VpnGateway Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileVpnGateway{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("vpngateway-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource VpnGateway
	err = c.Watch(&source.Kind{Type: &awsvpnv1alpha1.VpnGateway{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner VpnGateway
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &awsvpnv1alpha1.VpnGateway{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileVpnGateway implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileVpnGateway{}

// ReconcileVpnGateway reconciles a VpnGateway object
type ReconcileVpnGateway struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a VpnGateway object and makes changes based on the state read
// and what is in the VpnGateway.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileVpnGateway) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling VpnGateway")

	// Fetch the VpnGateway instance
	instance := &awsvpnv1alpha1.VpnGateway{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func newEc2Client(kubeClient client.Client, namespace, region string) (*ec2.EC2, error) {
	awsConfig := &aws.Config{Region: aws.String(region)}

	secret := &corev1.Secret{}
	err := kubeClient.Get(context.TODO(),
		types.NamespacedName{
			Name:      awsCredsSecretName,
			Namespace: namespace,
		},
		secret)
	if err != nil {
		return nil, err
	}
	accessKeyID, ok := secret.Data[awsCredsSecretIDKey]
	if !ok {
		return nil, fmt.Errorf("AWS credentials secret %v did not contain key %v",
			awsCredsSecretName, awsCredsSecretIDKey)
	}
	secretAccessKey, ok := secret.Data[awsCredsSecretAccessKey]
	if !ok {
		return nil, fmt.Errorf("AWS credentials secret %v did not contain key %v",
			awsCredsSecretName, awsCredsSecretAccessKey)
	}

	awsConfig.Credentials = credentials.NewStaticCredentials(
		string(accessKeyID), string(secretAccessKey), "")

	// Otherwise default to relying on the IAM role of the masters where the actuator is running:
	s, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}

	return ec2.New(s), nil
}

// getPlatformStatus provides a backwards-compatible way to look up platform status. AWS is the
// special case. 4.1 clusters on AWS expose the region config only through install-config. New AWS clusters
// and all other 4.2+ platforms are configured via platform status.
func getPlatformStatus(client client.Client) (*configv1.PlatformStatus, error) {
	var err error

	// Retrieve the cluster infrastructure config.
	infra := &configv1.Infrastructure{}
	err = client.Get(context.TODO(), types.NamespacedName{Name: "cluster"}, infra)
	if err != nil {
		return nil, err
	}

	if status := infra.Status.PlatformStatus; status != nil {
		// Only AWS needs backwards compatibility with install-config
		if status.Type != configv1.AWSPlatformType {
			return status, nil
		}

		// Check whether the cluster config is already migrated
		if status.AWS != nil && len(status.AWS.Region) > 0 {
			return status, nil
		}
	}

	// Otherwise build a platform status from the deprecated install-config
	type installConfig struct {
		Platform struct {
			AWS struct {
				Region string `json:"region"`
			} `json:"aws"`
		} `json:"platform"`
	}
	clusterConfigName := types.NamespacedName{Namespace: "kube-system", Name: "cluster-config-v1"}
	clusterConfig := &corev1.ConfigMap{}
	if err := client.Get(context.TODO(), clusterConfigName, clusterConfig); err != nil {
		return nil, fmt.Errorf("failed to get configmap %s: %v", clusterConfigName, err)
	}
	data, ok := clusterConfig.Data["install-config"]
	if !ok {
		return nil, fmt.Errorf("missing install-config in configmap")
	}
	var ic installConfig
	if err := yaml.Unmarshal([]byte(data), &ic); err != nil {
		return nil, fmt.Errorf("invalid install-config: %v\njson:\n%s", err, data)
	}
	return &configv1.PlatformStatus{
		Type: infra.Status.Platform,
		AWS: &configv1.AWSPlatformStatus{
			Region: ic.Platform.AWS.Region,
		},
	}, nil
}
