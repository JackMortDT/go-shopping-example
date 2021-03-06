package shopping

import (
  "shopping/database"
)

func PriceCheck(itemId int) (float64, bool) {
  item := database.LoadItem(itemId)
  if item == nil {
    return 0, false
  }
  return item.Price, true
}