package purchase

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.PurchaseConfigSettings().DeclareModel()

	h.PurchaseConfigSettings().AddFields(map[string]models.FieldDefinition{
		"CompanyId": models.Many2OneField{
			RelationModel: h.Company(),
			String:        "Company",
			Required:      true,
			Default:       func(env models.Environment) interface{} { return env.Uid().company_id },
		},
		"PoLead": models.FloatField{
			Related: `CompanyId.PoLead`,
			String:  "Purchase Lead Time *",
		},
		"PoLock": models.SelectionField{
			Related: `CompanyId.PoLock`,
			String:  "Purchase Order Modification *",
		},
		"PoDoubleValidation": models.SelectionField{
			Related: `CompanyId.PoDoubleValidation`,
			String:  "Levels of Approvals *",
		},
		"PoDoubleValidationAmount": models.MonetaryField{
			Related: `CompanyId.PoDoubleValidationAmount`,
			String:  "Double validation amount *",
			//currency_field='company_currency_id'
		},
		"CompanyCurrencyId": models.Many2OneField{
			RelationModel: h.Currency(),
			Related:       `CompanyId.CurrencyId`,
			ReadOnly:      true,
			Help:          "Utility field to express amount currency",
		},
		"GroupProductVariant": models.SelectionField{
			Selection: types.Selection{
				"": "No variants on products",
				"": "Products can have several attributes, defining variants (Example: size, color,...)",
			},
			String: "Product Variants",
			Help: "Work with product variant allows you to define some variant" +
				"of the same products, an ease the product management in" +
				"the ecommerce for example",
			//implied_group='product.group_product_variant'
		},
		"GroupUom": models.SelectionField{
			Selection: types.Selection{
				"": "Products have only one unit of measure (easier)",
				"": "Some products may be sold/puchased in different units of measure (advanced)",
			},
			String: "Units of Measure",
			//implied_group='product.group_uom'
			Help: "Allows you to select and maintain different units of measure" +
				"for products.",
		},
		"GroupCostingMethod": models.SelectionField{
			Selection: types.Selection{
				"": "Set a fixed cost price on each product",
				"": "Use a 'Fixed', 'Real' or 'Average' price costing method",
			},
			String: "Costing Methods",
			//implied_group='stock_account.group_inventory_valuation'
			Help: "Allows you to compute product cost price based on average cost.",
		},
		"ModulePurchaseRequisition": models.SelectionField{
			Selection: types.Selection{
				"": "Purchase propositions trigger draft purchase orders to a single supplier",
				"": "Allow using call for tenders to get quotes from multiple suppliers (advanced)",
			},
			String: "Calls for Tenders",
			Help: "Calls for tenders are used when you want to generate requests" +
				"for quotations to several vendors for a given set of products." +
				"You can configure per product if you directly do a Request" +
				"for Quotation to one vendor or if you want a Call for Tenders" +
				"to compare offers from several vendors.",
		},
		"GroupWarningPurchase": models.SelectionField{
			Selection: types.Selection{
				"": "All the products and the customers can be used in purchase orders",
				"": "An informative or blocking warning can be set on a product or a customer",
			},
			String: "Warning",
			//implied_group='purchase.group_warning_purchase'
		},
		"ModuleStockDropshipping": models.SelectionField{
			Selection: types.Selection{
				"": "Suppliers always deliver to your warehouse(s)",
				"": "Allow suppliers to deliver directly to your customers",
			},
			String: "Dropshipping",
			Help: "" +
				"Creates the dropship Route and add more complex tests" +
				"-This installs the module stock_dropshipping.",
		},
		"GroupManageVendorPrice": models.SelectionField{
			Selection: types.Selection{
				"": "Manage vendor price on the product form",
				"": "Allow using and importing vendor pricelists",
			},
			String: "Vendor Price",
			//implied_group="purchase.group_manage_vendor_price"
		},
	})
	h.AccountConfigSettings().DeclareModel()

	h.AccountConfigSettings().AddFields(map[string]models.FieldDefinition{
		"GroupAnalyticAccountForPurchases": models.BooleanField{
			String: "Analytic accounting for purchases",
			//implied_group='purchase.group_analytic_accounting'
			Help: "Allows you to specify an analytic account on purchase order lines.",
		},
	})
}
