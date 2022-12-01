package oneplatform

import (
	"github.com/benzdeus/oneplatform/configs"
	"github.com/benzdeus/oneplatform/env"
)

func Init(options configs.Options) {
	env.Options = options
}
