package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/utils/pointer"
	"net/http"
	"os/exec"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"strings"
)

func main() {
	router := gin.Default()
	SetupRouter(router)
}

func DeployK8s(ctx *gin.Context) {

	//kubeconfigPath := ""
	config, err := config.GetConfig()
	if err != nil {
		fmt.Printf("Error building config: %s\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating cleintset: %s\n", err.Error())
		return
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "my-deploy",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "example-app"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"app": "example-app"}},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  "example-container",
						Image: "nginx:latest",
					}},
				},
			},
		},
	}

	_, err = clientset.AppsV1().Deployments("default").Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Failed to create Deployment:", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "OK")

}

func GetPod(ctx *gin.Context) {
	kubeconfigPath := "/Users/huannguyen/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pods, err := clientset.CoreV1().Pods("default").List(ctx.Request.Context(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Failed to list Pods:", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	var podsName []string
	for _, pod := range pods.Items {
		podsName = append(podsName, pod.String())
	}

	ctx.JSON(http.StatusOK, podsName)

}

func UploadFile(ctx *gin.Context) {
	file, fh, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	tempFile := "/Users/huannguyen/Documents/GO/Test/" + fh.Filename
	if err := ctx.SaveUploadedFile(fh, tempFile); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error0": err.Error()})
		return
	}
	deploymentName := "nginx-deployment" // Tên của Deployment

	cmd := exec.Command("kubectl", "get", "pod", "--selector=app="+deploymentName, "-o=jsonpath={.items..metadata.name}")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting Pod name:", err)
		return
	}
	podNames := strings.Split(string(output), " ")
	for _, podName := range podNames {

		// Sao chép tệp zip vào container
		containerName := deploymentName // Tên của Container trong Pod
		targetPath := "/www/onefarm"    // Đường dẫn đích trong container
		fmt.Println("kubectl", "cp", tempFile, podName+":"+targetPath, "-c", containerName)
		cmd = exec.Command("kubectl", "cp", tempFile, podName+":"+targetPath, "-c", containerName)
		err = cmd.Run()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
			return
		}

		// Giải nén tệp zip trong container
		execCommand := "unzip -o " + targetPath + "/" + fh.Filename + " -d " + targetPath + "/"
		cmd = exec.Command("kubectl", "exec", podName, "-c", containerName, "--", "sh", "-c", execCommand)
		err = cmd.Run()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "File upload pushed to PVC"})
}

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("", DeployK8s)
		v1.GET("", GetPod)
		v1.POST("/upload-file", UploadFile)
	}
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Opening HTTP server: %v", err)
		panic(err)
	}
}
