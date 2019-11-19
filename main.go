package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/yuin/goldmark"
)

func main() {
	content := `## Deploying Federation and configuring an external policy engine

The Federation control plane can be deployed using kubefed init.

After deploying the Federation control plane, you must configure an Admission
Controller in the Federation API server that enforces placement decisions
received from the external policy engine.

    kubectl apply -f scheduling-policy-admission.yaml

Shown below is an example ConfigMap for the Admission Controller:

















<table class="includecode" id="federation-scheduling-policy-admission-yaml">
    <thead>
        <tr>
            <th>
                <a href="https://raw.githubusercontent.com/kubernetes/website/master/content/en/examples/federation/scheduling-policy-admission.yaml" download="federation/scheduling-policy-admission.yaml">
                    <code>federation/scheduling-policy-admission.yaml</code>
                </a>
                <img src="/images/copycode.svg" style="max-height:24px" onclick="copyCode('federation-scheduling-policy-admission-yaml')" title="Copy federation/scheduling-policy-admission.yaml to clipboard">
            </th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td><div class="highlight"><pre style="background-color:#f8f8f8;-moz-tab-size:4;-o-tab-size:4;tab-size:4"><code class="language-yaml" data-lang="yaml">apiVersion:<span style="color:#bbb"> </span>v1<span style="color:#bbb">
</span><span style="color:#bbb"></span>kind:<span style="color:#bbb"> </span>ConfigMap<span style="color:#bbb">
</span><span style="color:#bbb"></span>metadata:<span style="color:#bbb">
</span><span style="color:#bbb">  </span>name:<span style="color:#bbb"> </span>admission<span style="color:#bbb">
</span><span style="color:#bbb">  </span>namespace:<span style="color:#bbb"> </span>federation-system<span style="color:#bbb">
</span><span style="color:#bbb"></span>data:<span style="color:#bbb">
</span><span style="color:#bbb">  </span>config.yml:<span style="color:#bbb"> </span><span style="color:#b44;font-style:italic">|
</span><span style="color:#b44;font-style:italic">   </span><span style="color:#b44;font-style:italic"> </span><span style="color:#b44;font-style:italic">apiVersion: apiserver.k8s.io/v1alpha1</span><span style="color:#bbb">
</span><span style="color:#bbb">    </span>kind:<span style="color:#bbb"> </span>AdmissionConfiguration<span style="color:#bbb">
</span><span style="color:#bbb">    </span>plugins:<span style="color:#bbb">
</span><span style="color:#bbb">    </span>-<span style="color:#bbb"> </span>name:<span style="color:#bbb"> </span>SchedulingPolicy<span style="color:#bbb">
</span><span style="color:#bbb">      </span>path:<span style="color:#bbb"> </span>/etc/kubernetes/admission/scheduling-policy-config.yml<span style="color:#bbb">
</span><span style="color:#bbb">  </span>scheduling-policy-config.yml:<span style="color:#bbb"> </span><span style="color:#b44;font-style:italic">|
</span><span style="color:#b44;font-style:italic">   </span><span style="color:#b44;font-style:italic"> </span><span style="color:#b44;font-style:italic">kubeconfig: /etc/kubernetes/admission/opa-kubeconfig</span><span style="color:#bbb">
</span><span style="color:#bbb">  </span>opa-kubeconfig:<span style="color:#bbb"> </span><span style="color:#b44;font-style:italic">|
</span><span style="color:#b44;font-style:italic">   </span><span style="color:#b44;font-style:italic"> </span><span style="color:#b44;font-style:italic">clusters:</span><span style="color:#bbb">
</span><span style="color:#bbb">      </span>-<span style="color:#bbb"> </span>name:<span style="color:#bbb"> </span>opa-api<span style="color:#bbb">
</span><span style="color:#bbb">        </span>cluster:<span style="color:#bbb">
</span><span style="color:#bbb">          </span>server:<span style="color:#bbb"> </span>http://opa.federation-system.svc.cluster.local:<span style="color:#666">8181</span>/v0/data/kubernetes/placement<span style="color:#bbb">
</span><span style="color:#bbb">    </span>users:<span style="color:#bbb">
</span><span style="color:#bbb">      </span>-<span style="color:#bbb"> </span>name:<span style="color:#bbb"> </span>scheduling-policy<span style="color:#bbb">
</span><span style="color:#bbb">        </span>user:<span style="color:#bbb">
</span><span style="color:#bbb">          </span>token:<span style="color:#bbb"> </span>deadbeefsecret<span style="color:#bbb">
</span><span style="color:#bbb">    </span>contexts:<span style="color:#bbb">
</span><span style="color:#bbb">      </span>-<span style="color:#bbb"> </span>name:<span style="color:#bbb"> </span>default<span style="color:#bbb">
</span><span style="color:#bbb">        </span>context:<span style="color:#bbb">
</span><span style="color:#bbb">          </span>cluster:<span style="color:#bbb"> </span>opa-api<span style="color:#bbb">
</span><span style="color:#bbb">          </span>user:<span style="color:#bbb"> </span>scheduling-policy<span style="color:#bbb">
</span><span style="color:#bbb">    </span>current-context:<span style="color:#bbb"> </span>default<span style="color:#bbb">
</span></code></pre></div>  </td>
        </tr>
    </tbody>
</table>

`

	markdown := goldmark.New()

	var buf bytes.Buffer
	err := markdown.Convert([]byte(content), &buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())
}
