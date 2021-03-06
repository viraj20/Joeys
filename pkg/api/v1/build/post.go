package build

import (
	"fmt"
	"net/http"
	"time"

	v1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	//"k8s.io/client-go/kubernetes/typed/batch/v1/"
)

//Post To Create Kubernetes Job To Build Images
func Post(responseWriter http.ResponseWriter, request *http.Request) {

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	jobsClient := clientset.BatchV1().Jobs(corev1.NamespaceDefault)

	job := &v1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "dockerimagebuildjob",
			Namespace: "default",
		},
		Spec: v1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "dockerimagebuildjob",
							Image:   "docker",
							Command: []string{"/bin/sh", "-c"},
							Args:    []string{"docker login -u viraj24 -p Pass@123; docker build -f /clone-volume/Dockerfile -t viraj24/viraj:hellotestdockerimage /clone-volume/; docker push viraj24/viraj:hellotestdockerimage"},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "clone-volume",
									MountPath: "/clone-volume",
								},
								{
									Name:      "docker-pv-storage",
									MountPath: "/var/run/docker.sock",
								},
							},
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
					InitContainers: []corev1.Container{
						{
							Name:    "git-clone",
							Image:   "alpine/git",
							Command: []string{"/bin/sh", "-c", "git clone https://github.com/viraj20/docker.git /clone-volume"},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "clone-volume",
									MountPath: "/clone-volume",
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "clone-volume",
							VolumeSource: corev1.VolumeSource{
								EmptyDir: &corev1.EmptyDirVolumeSource{},
							},
						},
						{
							Name: "docker-pv-storage",
							VolumeSource: corev1.VolumeSource{
								PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
									ClaimName: "docker-pv-claim",
								},
							},
						},
					},
				},
			},
		},
	}
	_, err = jobsClient.Create(job)
	if err != nil {
		panic(err)
	}

	sharedInformerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)
	podInformer := sharedInformerFactory.Core().V1().Pods().Informer()
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Printf("service changed: %s \n", newObj)
		},
	})
	stop := make(chan struct{})
	go podInformer.Run(stop)
	if !cache.WaitForCacheSync(stop, podInformer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}
	<-stop

}
