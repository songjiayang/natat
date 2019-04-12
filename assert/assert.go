package assert

import (
	"github.com/gortc/stun"
)

func IsSymmetric(xorAddr, xorAddr2 *stun.XORMappedAddress) bool {
	return xorAddr.String() != xorAddr2.String()
}
