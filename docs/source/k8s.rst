Kubernetes work
===============

Workceptor has the ability to launch Kubernetes pods to perform work.

foo.yml

.. code-block:: yaml

    ---
    - node:
        id: foo

    - log-level:
        level: Debug

    - tcp-listener:
        port: 2222

    - control-service:
        service: control
        filename: /tmp/foo.sock

    - work-kubernetes:
        worktype: kubeit
        authmethod: kubeconfig
        allowruntimeauth: true
        allowruntimepod: true
        allowruntimeparams: true

kubeitpod.yml

.. code-block:: yaml

    apiVersion: v1
    kind: Pod
    metadata:
        generateName: myapp-pod-
        labels:
            app: myapp
    spec:
        containers:
        - name: worker
          image: busybox
          command: ['sh', '-c', 'echo The Pod is running && sleep 6 && exit 0']
        restartPolicy: Never

Note: at least one of the containers in the pod spec must be named "worker". This is the container that stdin is passed into, and that stdout is retrieved from.

.. code-block:: yaml

    $ receptorctl --socket /tmp/foo.sock work submit kubeit --param secret_kube_config=@$HOME/.kube/config --param secret_kube_pod=@kubeitpod.yml --no-payload
    Result:  Job Started
    Unit ID: FfpQ4zk2

``secret_kube_config`` The contents of kubeconfig file. The "@" tells receptorctl to read in a file name and pass the contents on.

``secret_kube_pod`` The contents of a pod definition. The "@" tells receptorctl to read in a file name and pass the contents on.

Runtime Params
^^^^^^^^^^^^^^

Additional parameters can be passed in when issuing a "work submit" command, using "--param" in receptorctl. These params must have the correct ``allowruntime*`` fields specified in the ``work-kubernetes`` definition.

.. list-table::
    :widths: 25 50 25
    :header-rows: 1

    * - param
      - description
      - permission
    * - kube_image
      - container image to use
      - allowruntimecommand
    * - kube_command
      - command container should run
      - allowruntimecommand
    * - kube_params
      - parameters to pass into kube_command
      - allowruntimeparams
    * - kube_namespace
      - kubernetes namespace to use
      - allowruntimeauth
    * - secret_kube_config
      - kubeconfig to authenticate with
      - allowruntimeauth
    * - secret_kube_pod
      - pod definition
      - allowruntimepod
    * - pod_pending_timeout
      - allowed duration for pod to be Pending
      - allowruntimeparams

``pod_pending_timeout`` is provided as a string, for example 1h20m30s or 30m10s.
