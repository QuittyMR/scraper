package scraper

import (
	"fmt"
	"quitty.tech/scraper/utils"
)

func (scraper Scraper) parsingError(err error) {
	utils.BaseError(err, fmt.Sprintf("failed parsing the node hierarchy: %v", scraper.target))
}

func (scraper Scraper) renderingError(err error) {
	utils.BaseError(err, fmt.Sprintf("failed rendering the node hierarchy to text: %v", scraper.target))
}

func (scraper Scraper) unknownTargetType() {
	utils.BaseError(nil, fmt.Sprintf("unknown target type to scrape: %v (%T)", scraper.target, scraper.target))
}
