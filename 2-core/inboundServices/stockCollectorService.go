package inboundServices

import (
	"stockCollector/2-core/outboundProviders"
)

func CollectStock(stockProvider outboundProviders.StockProviderInterface) error {
	response, err := stockProvider.GetDailyStockPriceHistory20y("MSFT")

	_ = response
	return err
}
