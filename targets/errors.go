package targets

import (
	"fmt"
	"quitty.tech/scraper/utils"
)

func (target responseTarget) UriError(err error) {
	utils.BaseError(err, fmt.Sprintf("Failed reaching URI: %v", target.Name()))
}

func (target responseTarget) ContentMissingError() {
	utils.BaseError(nil, fmt.Sprintf("Target URI has no content: %v", target.Name()))
}

func (target responseTarget) marshallingError(err error) {
	utils.BaseError(err, fmt.Sprintf("Failed deserializing page content at: %v", target.Name()))
}

func (target responseTarget) bodyHandleError(err error) {
	utils.BaseError(err, fmt.Sprintf("Failed closing the body handle: %v", target.Name()))
}

func (target nodeTarget) renderingError(err error) {
	utils.BaseError(err, fmt.Sprintf("failed rendering the target hierarchy to text: %v", target.Name()))
}

func (target nodeTarget) ambiguousTargetError() {
	utils.BaseError(nil, fmt.Sprintf("cannot locate root for given target: %v", target.Name()))
}
