package purchase

	import (
		"net/http"

		"github.com/hexya-erp/hexya/src/controllers"
		"github.com/hexya-erp/hexya/src/models"
		"github.com/hexya-erp/hexya/src/models/types"
		"github.com/hexya-erp/hexya/src/models/types/dates"
		"github.com/hexya-erp/pool/h"
		"github.com/hexya-erp/pool/q"
	)
	
func init() {


h.Partner().Methods().PurchaseInvoiceCount().DeclareMethod(
`PurchaseInvoiceCount`,
func(rs m.PartnerSet)  {
//        all_partners = self.search([('id', 'child_of', self.ids)])
//        all_partners.read(['parent_id'])
//        purchase_order_groups = self.env['purchase.order'].read_group(
//            domain=[('partner_id', 'in', all_partners.ids)],
//            fields=['partner_id'], groupby=['partner_id']
//        )
//        for group in purchase_order_groups:
//            partner = self.browse(group['partner_id'][0])
//            while partner:
//                if partner in self:
//                    partner.purchase_order_count += group['partner_id_count']
//                partner = partner.parent_id
//        supplier_invoice_groups = self.env['account.invoice'].read_group(
//            domain=[('partner_id', 'in', all_partners.ids),
//                    ('type', '=', 'in_invoice')],
//            fields=['partner_id'], groupby=['partner_id']
//        )
//        for group in supplier_invoice_groups:
//            partner = self.browse(group['partner_id'][0])
//            while partner:
//                if partner in self:
//                    partner.supplier_invoice_count += group['partner_id_count']
//                partner = partner.parent_id
})
h.Partner().Methods().CommercialFields().DeclareMethod(
`CommercialFields`,
func(rs m.PartnerSet)  {
//        return super(res_partner, self)._commercial_fields()
})
h.Partner().AddFields(map[string]models.FieldDefinition{
"PropertyPurchaseCurrencyId": models.Many2OneField{
RelationModel: h.Currency(),
String: "Supplier Currency",
//company_dependent=True
Help: "This currency will be used, instead of the default one," + 
"for purchases from the current partner",
},
"PurchaseOrderCount": models.IntegerField{
Compute: h.Partner().Methods().PurchaseInvoiceCount(),
String: "# of Purchase Order",
},
"SupplierInvoiceCount": models.IntegerField{
Compute: h.Partner().Methods().PurchaseInvoiceCount(),
String: "# Vendor Bills",
},
"PurchaseWarn": models.SelectionField{
Selection: WARNING_MESSAGE,
String: "Purchase Order",
{"_fields":["id","ctx"],"ctx":null,"id":"WARNING_HELP","loc":{"end":{"column":60,"line":53},"start":{"column":48,"line":53}},"type":"Name"}
Required: true,
Default: models.DefaultValue("no-message"),
},
"PurchaseWarnMsg": models.TextField{
String: "Message for Purchase Order",
},
})
}