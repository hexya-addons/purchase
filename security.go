package purchase

import (
	"github.com/hexya-erp/pool/h"
)

//vars

var (
	//User
	GroupPurchaseUser *security.Group
	//Manager
	GroupPurchaseManager *security.Group
	//Analytic Accounting for Purchases
	GroupAnalyticAccounting *security.Group
	//Manage Vendor Price
	GroupManageVendorPrice *security.Group
	//A warning can be set on a product or a customer (Purchase)
	GroupWarningPurchase *security.Group
)


//rights
func init() {
	h.PurchaseOrder().Methods().AllowAllToGroup(GroupPurchaseUser)
	h.PurchaseOrder().Methods().AllowAllToGroup(GroupPurchaseManager)
	h.PurchaseOrderLine().Methods().AllowAllToGroup(GroupPurchaseUser)
	h.PurchaseOrderLine().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Stock.ModelStockLocation().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.Stock.ModelStockWarehouse().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.Stock.ModelStockPicking().Methods().AllowAllToGroup(GroupPurchaseUser)
	h.Stock.ModelStockMove().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.Stock.ModelStockMove().Methods().Write().AllowGroup(GroupPurchaseUser)
	h.Stock.ModelStockMove().Methods().Create().AllowGroup(GroupPurchaseUser)
	h.PurchaseOrder().Methods().Load().AllowGroup(GroupStockUser)
	h.PurchaseOrderLine().Methods().Load().AllowGroup(GroupStockUser)
	h.Account.ModelAccountTax().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.PurchaseReport().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.PurchaseReport().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.PurchaseOrderLine().Methods().AllowAllToGroup(GroupPurchaseManager)
	h.Stock.ModelStockLocation().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Stock.ModelStockWarehouse().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Stock.ModelStockPicking().Methods().AllowAllToGroup(GroupPurchaseManager)
	h.Stock.ModelStockMove().Methods().AllowAllToGroup(GroupPurchaseManager)
	h.Account.ModelAccountTax().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Product.ModelProductProduct().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.Product.ModelProductProduct().Methods().AllowAllToGroup(GroupPurchaseManager)
	h.Product.ModelProductTemplate().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.Product.ModelProductTemplate().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Account.ModelAccountInvoice().Methods().AllowAllToGroup(GroupPurchaseUser)
	h.Account.ModelAccountInvoiceLine().Methods().AllowAllToGroup(GroupPurchaseUser)
	h.Account.ModelAccountInvoice().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Account.ModelAccountInvoiceLine().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Account.ModelAccountInvoiceTax().Methods().AllowAllToGroup(GroupPurchaseUser)
	h.Account.ModelAccountFiscalPosition().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.Base.ModelResPartner().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.Account.ModelAccountJournal().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.Account.ModelAccountJournal().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Account.ModelAccountMove().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.Account.ModelAccountMoveLine().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.Account.ModelAccountAnalyticLine().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.Base.ModelResPartner().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Base.ModelResPartner().Methods().Write().AllowGroup(GroupPurchaseManager)
	h.Base.ModelResPartner().Methods().Create().AllowGroup(GroupPurchaseManager)
	h.Product.ModelProductUomCateg().Methods().AllowAllToGroup(GroupPurchaseManager)
	h.Product.ModelProductUom().Methods().AllowAllToGroup(GroupPurchaseManager)
	h.Product.ModelProductCategory().Methods().AllowAllToGroup(GroupPurchaseManager)
	h.Product.ModelProductTemplate().Methods().AllowAllToGroup(GroupPurchaseManager)
	h.Product.ModelProductPackaging().Methods().AllowAllToGroup(GroupPurchaseManager)
	h.Product.ModelProductSupplierinfo().Methods().AllowAllToGroup(GroupPurchaseManager)
	h.Base.ModelResPartner().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Base.ModelResPartner().Methods().Write().AllowGroup(GroupPurchaseManager)
	h.Base.ModelResPartner().Methods().Create().AllowGroup(GroupPurchaseManager)
	h.Product.ModelProductPricelistItem().Methods().AllowAllToGroup(GroupPurchaseManager)
	h.Account.ModelAccountAccount().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Account.ModelAccountJournal().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Stock.ModelStockLocation().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Stock.ModelStockWarehouseOrderpoint().Methods().Load().AllowGroup(GroupPurchaseManager)
	h.Stock.ModelStockWarehouseOrderpoint().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.Product.ModelProductPriceHistory().Methods().Load().AllowGroup(GroupPurchaseUser)
	h.Product.ModelProductPriceHistory().Methods().AllowAllToGroup(GroupPurchaseManager)
	h.PurchaseOrder().Methods().Load().AllowGroup(GroupAccountInvoice)
	h.PurchaseOrder().Methods().Write().AllowGroup(GroupAccountInvoice)
	h.PurchaseOrderLine().Methods().Load().AllowGroup(GroupAccountInvoice)
	h.PurchaseOrderLine().Methods().Write().AllowGroup(GroupAccountInvoice)
}
