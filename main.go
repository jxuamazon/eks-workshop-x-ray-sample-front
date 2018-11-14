package main

import (
	"io"
	"net/http"
	"fmt"
	"io/ioutil"

	_ "github.com/aws/aws-xray-sdk-go/plugins/ecs"
	"github.com/aws/aws-xray-sdk-go/xray"
)

const appName = "eks-workshop-x-ray-sample-front"

func init() {
	xray.Configure(xray.Config{
		DaemonAddr:     "xray-service.default:2000",
		LogLevel:       "info",
		ServiceVersion: "1.2.3",
	})
}

func main() {
	http.Handle("/", xray.Handler(xray.NewFixedSegmentNamer(appName), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {


		resp, err := http.Get("http://x-ray-sample-back-k8s.default.svc.cluster.local")

		if err != nil {
			fmt.Println(err)
			return
		}

		defer resp.Body.Close()

		fmt.Println(resp.Status)

		if resp.StatusCode == http.StatusOK {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, string(body))
		}
	})))
	http.ListenAndServe(":8080", nil)
}

