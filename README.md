# KubeSphere client-go Project

The KubeSphere client-go Project is a rest-client of go libraries for communicating with the KubeSphere API Server.

# How to use it

## 1. Get client-go packages

```shell
go get kubesphere.io/client-go@59e48ec
```

```golang
import (
	"kubesphere.io/client-go/rest"
)
```
## 2. Create a generic client instance

### Option 1. Authenticate by KubeSphere User
```golang
var client client.Client
config := &rest.Config{
    Host:     "ks-apiserver.kubesphere-system:9090",
    Username: "admin",
    Password: "P@88w0rd",
}
client, err = rest.RESTClientFor(config)
if err != nil {
    return err
}

```

### Option 2. Authenticate by KubeSphere ServiceAccount

At first, You need to create a KubeSphere ServiceAccount.

```yaml
apiVersion: kubesphere.io/v1alpha1
kind: ServiceAccount
metadata:
  name: test-sa
  namespace: test
```

> You should bind roles to KubeSphere ServiceAccount before using it.

Mounting KubeSphere ServiceAccount to your workloads.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-workload
  namespace: test
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: test-workload
  template:
    metadata:
      labels:
        app: test-workload
      annotations:
        ## using annotation to mount KubeSphere ServiceAccount test-sa
        kubesphere.io/serviceaccount-name: test-sa
    spec:
      containers:
      - name: test-workload
        image: test-workload:latest
        ports:
        - containerPort: 80
```

Once the KubeSphere ServiceAccount is mounted successfully, you can use it in the pod belongs to the `test-workload`.

Create RESTClient with KubeSphere ServiceAccount. 
```golang
import (
	"kubesphere.io/client-go/rest"
)
```

```golang
config, err := rest.InClusterConfig()
if err != nil {
    return err
}
client, err := rest.RESTClientFor(config)
if err != nil {
    return err
}
```

> rest.RESTClientFor returns a *rest.RESTClient that reads and writes from/to an KubeSphere API server. 

> It's only compatible with Kubernetes-like API objects.

## 3. Using RESTClient

```golang
resp := client.Get().
    Group("tenant.kubesphere.io").
    Version("v1beta1").
    Resource("workspaces").
    Do(context.TODO())

list := &v1beta1.WorkspaceList{}
if resp.Error() != nil {
    return resp.Error()
}

err := resp.Into(list)
if err != nil {
    return err
}
```

The KubeSphere API Architecture can be found at https://kubesphere.io/docs/reference/api-docs/

# Where does it come from?

client-go is synced from https://github.com/kubesphere/kubesphere/blob/master/staging/src/kubesphere.io/client-go. Code changes are made in that location, merged into `kubesphere.io/client-go` and later synced here.