package committee

import (
	"github.com/kava-labs/cosmos-sdk/x/params"
)

type ParamKeeper interface {
	GetSubspace(string) (params.Subspace, bool)
}
