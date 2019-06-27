package purchase

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {

	h.Company().AddFields(map[string]models.FieldDefinition{
		"PoLead": models.FloatField{
			String:   "Purchase Lead Time",
			Required: true,
			Help: "Margin of error for vendor lead times. When the system" +
				"generates Purchase Orders for procuring products, they" +
				"will be scheduled that many days earlier to cope with unexpected" +
				"vendor delays.",
			Default: models.DefaultValue(0),
		},
		"PoLock": models.SelectionField{
			Selection: types.Selection{
				"edit": "Allow to edit purchase orders",
				"lock": "Confirmed purchase orders are not editable",
			},
			String:  "Purchase Order Modification",
			Default: models.DefaultValue("edit"),
			Help: "Purchase Order Modification used when you want to purchase" +
				"order editable after confirm",
		},
		"PoDoubleValidation": models.SelectionField{
			Selection: types.Selection{
				"one_step": "Confirm purchase orders in one step",
				"two_step": "Get 2 levels of approvals to confirm a purchase order",
			},
			String:  "Levels of Approvals",
			Default: models.DefaultValue("one_step"),
			Help:    "Provide a double validation mechanism for purchases",
		},
		"PoDoubleValidationAmount": models.MonetaryField{
			String:  "Double validation amount",
			Default: models.DefaultValue(5000),
			Help:    "Minimum amount for which a double validation is required",
		},
	})
}
