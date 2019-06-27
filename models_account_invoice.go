package purchase

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.AccountInvoice().DeclareModel()

	h.AccountInvoice().AddFields(map[string]models.FieldDefinition{
		"PurchaseId": models.Many2OneField{
			RelationModel: h.PurchaseOrder(),
			String:        "Add Purchase Order",
			Help: "Encoding help. When selected, the associated purchase order" +
				"lines are added to the vendor bill. Several PO can be selected.",
		},
	})
	h.AccountInvoice().Methods().OnchangeAllowedPurchaseIds().DeclareMethod(
		`
        The purpose of the method is to define a domain
for the available
        purchase orders.
        `,
		func(rs m.AccountInvoiceSet) {
			//        result = {}
			//        purchase_line_ids = self.invoice_line_ids.mapped('purchase_line_id')
			//        purchase_ids = self.invoice_line_ids.mapped('purchase_id').filtered(
			//            lambda r: r.order_line <= purchase_line_ids)
			//        domain = [('invoice_status', '=', 'to invoice')]
			//        if self.partner_id:
			//            domain += [('partner_id', 'child_of', self.partner_id.id)]
			//        if purchase_ids:
			//            domain += [('id', 'not in', purchase_ids.ids)]
			//        result['domain'] = {'purchase_id': domain}
			//        return result
		})
	h.AccountInvoice().Methods().PrepareInvoiceLineFromPoLine().DeclareMethod(
		`PrepareInvoiceLineFromPoLine`,
		func(rs m.AccountInvoiceSet, line interface{}) {
			//        if line.product_id.purchase_method == 'purchase':
			//            qty = line.product_qty - line.qty_invoiced
			//        else:
			//            qty = line.qty_received - line.qty_invoiced
			//        if float_compare(qty, 0.0, precision_rounding=line.product_uom.rounding) <= 0:
			//            qty = 0.0
			//        taxes = line.taxes_id
			//        invoice_line_tax_ids = line.order_id.fiscal_position_id.map_tax(taxes)
			//        invoice_line = self.env['account.invoice.line']
			//        data = {
			//            'purchase_line_id': line.id,
			//            'name': line.order_id.name+': '+line.name,
			//            'origin': line.order_id.origin,
			//            'uom_id': line.product_uom.id,
			//            'product_id': line.product_id.id,
			//            'account_id': invoice_line.with_context({'journal_id': self.journal_id.id, 'type': 'in_invoice'})._default_account(),
			//            'price_unit': line.order_id.currency_id.with_context(date=self.date_invoice).compute(line.price_unit, self.currency_id, round=False),
			//            'quantity': qty,
			//            'discount': 0.0,
			//            'account_analytic_id': line.account_analytic_id.id,
			//            'analytic_tag_ids': line.analytic_tag_ids.ids,
			//            'invoice_line_tax_ids': invoice_line_tax_ids.ids
			//        }
			//        account = invoice_line.get_invoice_line_account(
			//            'in_invoice', line.product_id, line.order_id.fiscal_position_id, self.env.user.company_id)
			//        if account:
			//            data['account_id'] = account.id
			//        return data
		})
	h.AccountInvoice().Methods().OnchangeProductId().DeclareMethod(
		`OnchangeProductId`,
		func(rs m.AccountInvoiceSet) {
			//        domain = super(AccountInvoice, self)._onchange_product_id()
			//        if self.purchase_id:
			//            # Use the purchase uom by default
			//            self.uom_id = self.product_id.uom_po_id
			//        return domain
		})
	h.AccountInvoice().Methods().PurchaseOrderChange().DeclareMethod(
		`PurchaseOrderChange`,
		func(rs m.AccountInvoiceSet) {
			//        if not self.purchase_id:
			//            return {}
			//        if not self.partner_id:
			//            self.partner_id = self.purchase_id.partner_id.id
			//        new_lines = self.env['account.invoice.line']
			//        for line in self.purchase_id.order_line - self.invoice_line_ids.mapped('purchase_line_id'):
			//            data = self._prepare_invoice_line_from_po_line(line)
			//            new_line = new_lines.new(data)
			//            new_line._set_additional_fields(self)
			//            new_lines += new_line
			//        self.invoice_line_ids += new_lines
			//        self.purchase_id = False
			//        return {}
		})
	h.AccountInvoice().Methods().OnchangeCurrencyId().DeclareMethod(
		`OnchangeCurrencyId`,
		func(rs m.AccountInvoiceSet) {
			//        if self.currency_id:
			//            for line in self.invoice_line_ids.filtered(lambda r: r.purchase_line_id):
			//                line.price_unit = line.purchase_id.currency_id.with_context(date=self.date_invoice).compute(
			//                    line.purchase_line_id.price_unit, self.currency_id, round=False)
		})
	h.AccountInvoice().Methods().OnchangeOrigin().DeclareMethod(
		`OnchangeOrigin`,
		func(rs m.AccountInvoiceSet) {
			//        purchase_ids = self.invoice_line_ids.mapped('purchase_id')
			//        if purchase_ids:
			//            self.origin = ', '.join(purchase_ids.mapped('name'))
		})
	h.AccountInvoice().Methods().OnchangePartnerId().DeclareMethod(
		`OnchangePartnerId`,
		func(rs m.AccountInvoiceSet) {
			//        res = super(AccountInvoice, self)._onchange_partner_id()
			//        if not self.env.context.get('default_journal_id') and self.partner_id and self.currency_id and\
			//                self.type in ['in_invoice', 'in_refund'] and\
			//                self.currency_id != self.partner_id.property_purchase_currency_id:
			//            journal_domain = [
			//                ('type', '=', 'purchase'),
			//                ('company_id', '=', self.company_id.id),
			//                ('currency_id', '=', self.partner_id.property_purchase_currency_id.id),
			//            ]
			//            default_journal_id = self.env['account.journal'].search(
			//                journal_domain, limit=1)
			//            if default_journal_id:
			//                self.journal_id = default_journal_id
			//        return res
		})
	h.AccountInvoice().Methods().InvoiceLineMoveLineGet().DeclareMethod(
		`InvoiceLineMoveLineGet`,
		func(rs m.AccountInvoiceSet) {
			//        res = super(AccountInvoice, self).invoice_line_move_line_get()
			//        if self.env.user.company_id.anglo_saxon_accounting:
			//            if self.type in ['in_invoice', 'in_refund']:
			//                for i_line in self.invoice_line_ids:
			//                    res.extend(
			//                        self._anglo_saxon_purchase_move_lines(i_line, res))
			//        return res
		})
	h.AccountInvoice().Methods().AngloSaxonPurchaseMoveLines().DeclareMethod(
		`Return the additional move lines for purchase invoices and refunds.

        i_line: An account.invoice.line object.
        res: The move line entries produced so far by the
parent move_line_get.
        `,
		func(rs m.AccountInvoiceSet, i_line interface{}, res interface{}) {
			//        inv = i_line.invoice_id
			//        company_currency = inv.company_id.currency_id
			//        if i_line.product_id and i_line.product_id.valuation == 'real_time' and i_line.product_id.type == 'product':
			//            # get the fiscal position
			//            fpos = i_line.invoice_id.fiscal_position_id
			//            # get the price difference account at the product
			//            acc = i_line.product_id.property_account_creditor_price_difference
			//            if not acc:
			//                # if not found on the product get the price difference account at the category
			//                acc = i_line.product_id.categ_id.property_account_creditor_price_difference_categ
			//            acc = fpos.map_account(acc).id
			//            # reference_account_id is the stock input account
			//            reference_account_id = i_line.product_id.product_tmpl_id.get_product_accounts(
			//                fiscal_pos=fpos)['stock_input'].id
			//            diff_res = []
			//            account_prec = inv.company_id.currency_id.decimal_places
			//            # calculate and write down the possible price difference between invoice price and product price
			//            for line in res:
			//                if line.get('invl_id', 0) == i_line.id and reference_account_id == line['account_id']:
			//                    valuation_price_unit = i_line.product_id.uom_id._compute_price(
			//                        i_line.product_id.standard_price, i_line.uom_id)
			//                    if i_line.product_id.cost_method != 'standard' and i_line.purchase_line_id:
			//                        # for average/fifo/lifo costing method, fetch real cost price from incomming moves
			//                        valuation_price_unit = i_line.purchase_line_id.product_uom._compute_price(
			//                            i_line.purchase_line_id.price_unit, i_line.uom_id)
			//                        stock_move_obj = self.env['stock.move']
			//                        valuation_stock_move = stock_move_obj.search(
			//                            [('purchase_line_id', '=', i_line.purchase_line_id.id), ('state', '=', 'done')])
			//                        if valuation_stock_move:
			//                            valuation_price_unit_total = 0
			//                            valuation_total_qty = 0
			//                            for val_stock_move in valuation_stock_move:
			//                                valuation_price_unit_total += val_stock_move.price_unit * val_stock_move.product_qty
			//                                valuation_total_qty += val_stock_move.product_qty
			//                            valuation_price_unit = valuation_price_unit_total / valuation_total_qty
			//                            valuation_price_unit = i_line.product_id.uom_id._compute_price(
			//                                valuation_price_unit, i_line.uom_id)
			//                    if inv.currency_id.id != company_currency.id:
			//                        valuation_price_unit = company_currency.with_context(
			//                            date=inv.date_invoice).compute(valuation_price_unit, inv.currency_id, round=False)
			//                    if valuation_price_unit != i_line.price_unit and line['price_unit'] == i_line.price_unit and acc:
			//                        # price with discount and without tax included
			//                        price_unit = i_line.price_unit * \
			//                            (1 - (i_line.discount or 0.0) / 100.0)
			//                        tax_ids = []
			//                        if line['tax_ids']:
			//                            # line['tax_ids'] is like [(4, tax_id, None), (4, tax_id2, None)...]
			//                            taxes = self.env['account.tax'].browse(
			//                                [x[1] for x in line['tax_ids']])
			//                            price_unit = taxes.compute_all(
			//                                price_unit, currency=inv.currency_id, quantity=1.0)['total_excluded']
			//                            for tax in taxes:
			//                                tax_ids.append((4, tax.id, None))
			//                                for child in tax.children_tax_ids:
			//                                    if child.type_tax_use != 'none':
			//                                        tax_ids.append((4, child.id, None))
			//                        price_before = line.get('price', 0.0)
			//                        line.update(
			//                            {'price': round(valuation_price_unit * line['quantity'], account_prec)})
			//                        diff_res.append({
			//                            'type': 'src',
			//                            'name': i_line.name[:64],
			//                            'price_unit': round(price_unit - valuation_price_unit, account_prec),
			//                            'quantity': line['quantity'],
			//                            'price': round(price_before - line.get('price', 0.0), account_prec),
			//                            'account_id': acc,
			//                            'product_id': line['product_id'],
			//                            'uom_id': line['uom_id'],
			//                            'account_analytic_id': line['account_analytic_id'],
			//                            'tax_ids': tax_ids,
			//                        })
			//            return diff_res
			//        return []
		})
	h.AccountInvoice().Methods().Create().Extend(
		`Create`,
		func(rs m.AccountInvoiceSet, vals models.RecordData) {
			//        invoice = super(AccountInvoice, self).create(vals)
			//        purchase = invoice.invoice_line_ids.mapped('purchase_line_id.order_id')
			//        if purchase and not invoice.refund_invoice_id:
			//            message = _("This vendor bill has been created from: %s") % (",".join(
			//                ["<a href=# data-oe-model=purchase.order data-oe-id="+str(order.id)+">"+order.name+"</a>" for order in purchase]))
			//            invoice.message_post(body=message)
			//        return invoice
		})
	h.AccountInvoice().Methods().Write().Extend(
		`Write`,
		func(rs m.AccountInvoiceSet, vals models.RecordData) {
			//        result = True
			//        for invoice in self:
			//            purchase_old = invoice.invoice_line_ids.mapped(
			//                'purchase_line_id.order_id')
			//            result = result and super(AccountInvoice, invoice).write(vals)
			//            purchase_new = invoice.invoice_line_ids.mapped(
			//                'purchase_line_id.order_id')
			//            # To get all po reference when updating invoice line or adding purchase order reference from vendor bill.
			//            purchase = (purchase_old | purchase_new) - \
			//                (purchase_old & purchase_new)
			//            if purchase:
			//                message = _("This vendor bill has been modified from: %s") % (",".join(
			//                    ["<a href=# data-oe-model=purchase.order data-oe-id="+str(order.id)+">"+order.name+"</a>" for order in purchase]))
			//                invoice.message_post(body=message)
			//        return result
		})
	h.AccountInvoiceLine().DeclareModel()

	h.AccountInvoiceLine().AddFields(map[string]models.FieldDefinition{
		"PurchaseLineId": models.Many2OneField{
			RelationModel: h.PurchaseOrderLine(),
			String:        "Purchase Order Line",
			OnDelete:      `set null`,
			Index:         true,
			ReadOnly:      true,
		},
		"PurchaseId": models.Many2OneField{
			RelationModel: h.PurchaseOrder(),
			Related:       `PurchaseLineId.OrderId`,
			String:        "Purchase Order",
			Stored:        false,
			ReadOnly:      true,
			//related_sudo=False
			Help: "Associated Purchase Order. Filled in automatically when" +
				"a PO is chosen on the vendor bill.",
		},
	})
}
