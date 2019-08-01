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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
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
