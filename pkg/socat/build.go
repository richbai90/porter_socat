package socat

import (
	"fmt"
	"strings"

	"github.com/getporter/socat/pkg/socat/socatfs"
)

// BuildInput represents stdin passed to the mixin for the build command.
type BuildInput struct {
	Config MixinConfig
}

// MixinConfig represents configuration that can be set on the socat mixin in porter.yaml
// mixins:
// - socat:
//	  clientVersion: "v0.0.0"

type MixinConfig struct {
}

// This is an example. Replace the following with whatever steps are needed to
// install required components into
// const dockerfileLines = `RUN apt-get update && \
// apt-get install gnupg apt-transport-https lsb-release software-properties-common -y && \
// echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ stretch main" | \
//    tee /etc/apt/sources.list.d/azure-cli.list && \
// apt-key --keyring /etc/apt/trusted.gpg.d/Microsoft.gpg adv \
// 	--keyserver packages.microsoft.com \
// 	--recv-keys BC528686B50D79E339D3721CEB3E94ADBE1229CF && \
// apt-get update && apt-get install azure-cli
// `

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {
	fmt.Fprintf(m.Out, `SHELL ["/bin/bash", "-c"]
	`)
	fmt.Fprintf(m.Out, "RUN apt-get update && apt-get install socat -y\n")
	fmt.Fprintf(m.Out, `RUN echo $'%s' > /usr/bin/porter_socat && `, socatfs.PorterSocat)
	fmt.Fprint(m.Out, `chmod +x /usr/bin/porter_socat
	SHELL ["/bin/sh", "-c"]
	`)
	// Example of pulling and defining a client version for your mixin
	// fmt.Fprintf(m.Out, "\nRUN curl https://get.helm.sh/helm-%s-linux-amd64.tar.gz --output helm3.tar.gz", m.ClientVersion)

	return nil
}
