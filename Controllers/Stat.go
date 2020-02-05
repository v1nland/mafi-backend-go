package Controllers

import (
    "mafi-backend-go/ApiHelpers"
	"mafi-backend-go/Config"
    "github.com/gin-gonic/gin"
)

func GetCurrency(c *gin.Context) {
    type CurrencyResult struct{
        Income      int
        Outcome     int
        Currency    int
    }

    var result CurrencyResult
    rows := Config.DB.Raw("SELECT income, outcome, income-outcome AS currency FROM (SELECT SUM(i.purchase_price * p.item_qty) as outcome FROM purchases AS p, items AS i WHERE p.item_id = i.id) AS T1, (SELECT SUM(ROUND(i.sell_price*(1-o.discount/100),0)) AS income FROM orders AS o, items AS i WHERE o.item_id = i.id AND o.finished = 1) AS T2").Scan(&result)
	if rows == nil {
		ApiHelpers.RespondJSON(c, 404, result)
	} else {
		ApiHelpers.RespondJSON(c, 200, result)
	}
}

func GetSources(c *gin.Context) {
    type SourceResult struct{
        Source      string
        Quantity    int
    }

    var result []SourceResult
    rows := Config.DB.Raw("SELECT source, COUNT(source) AS quantity FROM orders WHERE finished = 1 GROUP BY source").Scan(&result)
	if rows == nil {
		ApiHelpers.RespondJSON(c, 404, result)
	} else {
		ApiHelpers.RespondJSON(c, 200, result)
	}
}

func GetMostSold(c *gin.Context) {
    type MostSoldResult struct{
		Id			int
		Description	string
        Type		string
        Color		string
		Times		int
    }

    var result []MostSoldResult
    rows := Config.DB.Raw("SELECT item_id as id, description, type, color, timesSold as times FROM items, (SELECT item_id, count(item_id) as timesSold FROM orders WHERE finished = 1 GROUP BY item_id ORDER BY count(item_id) DESC LIMIT 1) as t1 WHERE items.id = t1.item_id").Scan(&result)
	if rows == nil {
		ApiHelpers.RespondJSON(c, 404, result)
	} else {
		ApiHelpers.RespondJSON(c, 200, result)
	}
}

func GetStock(c *gin.Context) {
    type StockResult struct{
		Id			int
		Description	string
        Bought		int
        Sold		int
        Type		string
        Color		string
		Sell		int
    }

    var result []StockResult
    rows := Config.DB.Raw("SELECT t1.id, t1.description, coalesce(t2.i_times_bought, 0) AS bought, coalesce(t2.i_times_sold, 0) AS sold, t1.type, t1.color, t1.sell_price AS sell FROM items AS t1 LEFT OUTER JOIN (SELECT t1.item_id, coalesce(i_times_bought,0) AS i_times_bought, coalesce(i_times_sold, 0) AS i_times_sold FROM (SELECT sum(item_qty) AS i_times_bought, item_id FROM purchases GROUP BY item_id) AS t1 LEFT OUTER JOIN (SELECT count(*) AS i_times_sold, item_id FROM orders GROUP BY item_id) AS t2 ON t1.item_id = t2.item_id) AS t2 ON t1.id = t2.item_id").Scan(&result)
	if rows == nil {
		ApiHelpers.RespondJSON(c, 404, result)
	} else {
		ApiHelpers.RespondJSON(c, 200, result)
	}
}
