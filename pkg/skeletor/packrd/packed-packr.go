// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "8aae90ad4052adebea0b2bc91b07861f"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"14c5f6c5f317dbae6025d7de9c86e4e9": "1f8b08000000000000ffcc543d4fc33010ddf32b4e2e634b6042caccc880c4883a98fa92ba248e395f870af5bfa3b86d6aa74e90502b31e6cef7debb8f97ef0c40dcb9d51a1b290a106b665be4f9c6b5667188deb754e58a64c98b87a7fc109b897957a7b0d446b36e8d1305745000421bc7b2aedf186d1f0410bcb3d811b41f1b5cb1aff7714bad45628d2e780d20dc27d6c82d7551e8c3e34053603ea7d0ad48db4eed45b29b0161d9c1cef2a0abdc31dae7a030aadacf6306231b4c419f143b266daa690c2b4936c848971d4cf7eef352292f5cd6afe393f85513c03e1bfb8af48ef195b276183d24fcda6a42250a784f0c6d7c084162990d05f562d2f8e71b3a8696fdfb29e159002db6b622a97078cde96bd91addd81a1b348cca97c45866c41b7f451ba4afe4b67f6ab65b1ea14ab35ee1e2a2950d1b4c2cece8c79ea6d1e6054dc56b51c0637692e20153cb492de5f4474ed14922b93bb369c6667019e95585879c85231af8e69a9ca115d39ce6069dc6a68d79fb4d4cfe4ef6d94f000000fffff35f899665070000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("schema", "./schema")
		b.SetResolver("skeletor.json", packr.Pointer{ForwardBox: gk, ForwardPath: "14c5f6c5f317dbae6025d7de9c86e4e9"})
	}()

	return nil
}()
