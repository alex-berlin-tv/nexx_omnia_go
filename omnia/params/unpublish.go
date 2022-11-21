package params

import (
	"github.com/alex-berlin-tv/nexx_omnia_go/omnia/enums"
	"github.com/pasztorpisti/qs"
)

// Parameters for the unpublish ManagementAPI call. The documentation can be found [here].
//
// [here]: https://api.docs.nexx.cloud/management-api/endpoints/management-endpoint#unpublish
type Unpublish struct {
	// If set to 1, any call to the /publish method will fail, unless the
	// /unblock Endpoint has been called.
	BlockFuturePublishing enums.Bool `qs:"blockFuturePublishing,omitempty"`
}

func (u Unpublish) UrlEncode(extra map[string]interface{}) (string, error) {
	return qs.Marshal(&u)
}