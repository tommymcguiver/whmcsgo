package whmcsgo

type ClientGroup struct {
	Groups struct {
		Group []struct {
			ID               int    `json:"id"`
			GroupName        string `json:"groupname"`
			GroupColour      string `json:"groupcolour"`
			DiscountPercent  string `json:"discountpercent"`
			Susptermexempt   string `json:"susptermexempt"`
			SeparateInvoices string `json:"separateinvoices"`
		} `json:"group"`
	} `json:"groups"`
	Result       string `json:"result"`
	TotalResults int    `json:"totalresults"`
}
