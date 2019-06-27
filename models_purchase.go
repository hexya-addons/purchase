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
	
//import odoo.addons.decimal_precision as dp
func init() {
h.PurchaseOrder().DeclareModel()




h.PurchaseOrder().Methods().AmountAll().DeclareMethod(
`AmountAll`,
func(rs m.PurchaseOrderSet)  {
//        for order in self:
//            amount_untaxed = amount_tax = 0.0
//            for line in order.order_line:
//                amount_untaxed += line.price_subtotal
//                # FORWARDPORT UP TO 10.0
//                if order.company_id.tax_calculation_rounding_method == 'round_globally':
//                    taxes = line.taxes_id.compute_all(
//                        line.price_unit, line.order_id.currency_id, line.product_qty, product=line.product_id, partner=line.order_id.partner_id)
//                    amount_tax += sum(t.get('amount', 0.0)
//                                      for t in taxes.get('taxes', []))
//                else:
//                    amount_tax += line.price_tax
//            order.update({
//                'amount_untaxed': order.currency_id.round(amount_untaxed),
//                'amount_tax': order.currency_id.round(amount_tax),
//                'amount_total': amount_untaxed + amount_tax,
//            })
})
h.PurchaseOrder().Methods().ComputeDatePlanned().DeclareMethod(
`ComputeDatePlanned`,
func(rs m.PurchaseOrderSet)  {
//        for order in self:
//            min_date = False
//            for line in order.order_line:
//                if not min_date or line.date_planned < min_date:
//                    min_date = line.date_planned
//            if min_date:
//                order.date_planned = min_date
})
h.PurchaseOrder().Methods().GetInvoiced().DeclareMethod(
`GetInvoiced`,
func(rs m.PurchaseOrderSet)  {
//        precision = self.env['decimal.precision'].precision_get(
//            'Product Unit of Measure')
//        for order in self:
//            if order.state not in ('purchase', 'done'):
//                order.invoice_status = 'no'
//                continue
//
//            if any(float_compare(line.qty_invoiced, line.product_qty if line.product_id.purchase_method == 'purchase' else line.qty_received, precision_digits=precision) == -1 for line in order.order_line):
//                order.invoice_status = 'to invoice'
//            elif all(float_compare(line.qty_invoiced, line.product_qty if line.product_id.purchase_method == 'purchase' else line.qty_received, precision_digits=precision) >= 0 for line in order.order_line) and order.invoice_ids:
//                order.invoice_status = 'invoiced'
//            else:
//                order.invoice_status = 'no'
})
h.PurchaseOrder().Methods().ComputeInvoice().DeclareMethod(
`ComputeInvoice`,
func(rs m.PurchaseOrderSet)  {
//        for order in self:
//            invoices = self.env['account.invoice']
//            for line in order.order_line:
//                invoices |= line.invoice_lines.mapped('invoice_id')
//            order.invoice_ids = invoices
//            order.invoice_count = len(invoices)
})
h.PurchaseOrder().Methods().DefaultPickingType().DeclareMethod(
`DefaultPickingType`,
func(rs m.PurchaseOrderSet)  {
//        type_obj = self.env['stock.picking.type']
//        company_id = self.env.context.get(
//            'company_id') or self.env.user.company_id.id
//        types = type_obj.search(
//            [('code', '=', 'incoming'), ('warehouse_id.company_id', '=', company_id)])
//        if not types:
//            types = type_obj.search(
//                [('code', '=', 'incoming'), ('warehouse_id', '=', False)])
//        return types[:1]
})
h.PurchaseOrder().Methods().ComputePicking().DeclareMethod(
`ComputePicking`,
func(rs m.PurchaseOrderSet)  {
//        for order in self:
//            pickings = self.env['stock.picking']
//            for line in order.order_line:
//                # We keep a limited scope on purpose. Ideally, we should also use move_orig_ids and
//                # do some recursive search, but that could be prohibitive if not done correctly.
//                moves = line.move_ids | line.move_ids.mapped(
//                    'returned_move_ids')
//                moves = moves.filtered(lambda r: r.state != 'cancel')
//                pickings |= moves.mapped('picking_id')
//            order.picking_ids = pickings
//            order.picking_count = len(pickings)
})
h.PurchaseOrder().Methods().ComputeIsShipped().DeclareMethod(
`ComputeIsShipped`,
func(rs m.PurchaseOrderSet)  {
//        for order in self:
//            if order.picking_ids and all([x.state == 'done' for x in order.picking_ids]):
//                order.is_shipped = True
})
//    READONLY_STATES = {
//        'purchase': [('readonly', True)],
//        'done': [('readonly', True)],
//        'cancel': [('readonly', True)],
//    }
h.PurchaseOrder().AddFields(map[string]models.FieldDefinition{
"Name": models.CharField{
String: "Order Reference",
Required: true,
Index: true,
NoCopy: true,
Default: models.DefaultValue("New"),
},
"Origin": models.CharField{
String: "Source Document",
NoCopy: true,
Help: "Reference of the document that generated this purchase" + 
"order request (e.g. a sale order or an internal procurement request)",
},
"PartnerRef": models.CharField{
String: "Vendor Reference",
NoCopy: true,
Help: "Reference of the sales order or bid sent by the vendor." + 
"It's used to do the matching when you receive the products" + 
"as this reference is usually written on the delivery order" + 
"sent by your vendor.",
},
"DateOrder": models.DateTimeField{
String: "Order Date",
Required: true,
//states=READONLY_STATES
Index: true,
NoCopy: true,
Default: func (env models.Environment) interface{} { return dates.Now() },
Help: "Depicts the date where the Quotation should be validated" + 
"and converted into a purchase order.",
},
"DateApprove": models.DateField{
String: "Approval Date",
ReadOnly: true,
Index: true,
NoCopy: true,
},
"PartnerId": models.Many2OneField{
RelationModel: h.Partner(),
String: "Vendor",
Required: true,
//states=READONLY_STATES
//change_default=True
//track_visibility='always'
},
"DestAddressId": models.Many2OneField{
RelationModel: h.Partner(),
String: "Drop Ship Address",
//states=READONLY_STATES
Help: "Put an address if you want to deliver directly from the" + 
"vendor to the customer. Otherwise, keep empty to deliver" + 
"to your own company.",
},
"CurrencyId": models.Many2OneField{
RelationModel: h.Currency(),
String: "Currency",
Required: true,
//states=READONLY_STATES
Default: func (env models.Environment) interface{} { return env.Uid().company_id.currency_id.id },
},
"State": models.SelectionField{
Selection: types.Selection{
"draft": "RFQ",
"sent": "RFQ Sent",
"to approve": "To Approve",
"purchase": "Purchase Order",
"done": "Locked",
"cancel": "Cancelled",
},
String: "Status",
ReadOnly: true,
Index: true,
NoCopy: true,
Default: models.DefaultValue("draft"),
//track_visibility='onchange'
},
"OrderLine": models.One2ManyField{
RelationModel: h.PurchaseOrderLine(),
ReverseFK: "",
String: "Order Lines",
//states={'cancel': [('readonly', True)], 'done': [('readonly', True)]}
NoCopy: false,
},
"Notes": models.TextField{
String: "Terms and Conditions",
},
"InvoiceCount": models.IntegerField{
Compute: h.PurchaseOrder().Methods().ComputeInvoice(),
String: "# of Bills",
NoCopy: true,
Default: models.DefaultValue(0),
},
"InvoiceIds": models.Many2ManyField{
RelationModel: h.AccountInvoice(),
Compute: h.PurchaseOrder().Methods().ComputeInvoice(),
String: "Bills",
NoCopy: true,
},
"InvoiceStatus": models.SelectionField{
Selection: types.Selection{
"no": "Nothing to Bill",
"to invoice": "Waiting Bills",
"invoiced": "Bills Received",
},
String: "Billing Status",
Compute: h.PurchaseOrder().Methods().GetInvoiced(),
Stored: true,
ReadOnly: true,
NoCopy: true,
Default: models.DefaultValue("no"),
},
"PickingCount": models.IntegerField{
Compute: h.PurchaseOrder().Methods().ComputePicking(),
String: "Receptions",
Default: models.DefaultValue(0),
},
"PickingIds": models.Many2ManyField{
RelationModel: h.StockPicking(),
Compute: h.PurchaseOrder().Methods().ComputePicking(),
String: "Receptions",
NoCopy: true,
},
"DatePlanned": models.DateTimeField{
String: "Scheduled Date",
Compute: h.PurchaseOrder().Methods().ComputeDatePlanned(),
Stored: true,
Index: true,
},
"AmountUntaxed": models.MonetaryField{
String: "Untaxed Amount",
Stored: true,
ReadOnly: true,
Compute: h.PurchaseOrder().Methods().AmountAll(),
//track_visibility='always'
},
"AmountTax": models.MonetaryField{
String: "Taxes",
Stored: true,
ReadOnly: true,
Compute: h.PurchaseOrder().Methods().AmountAll(),
},
"AmountTotal": models.MonetaryField{
String: "Total",
Stored: true,
ReadOnly: true,
Compute: h.PurchaseOrder().Methods().AmountAll(),
},
"FiscalPositionId": models.Many2OneField{
RelationModel: h.AccountFiscalPosition(),
String: "Fiscal Position",
//oldname='fiscal_position'
},
"PaymentTermId": models.Many2OneField{
RelationModel: h.AccountPaymentTerm(),
String: "Payment Terms",
},
"IncotermId": models.Many2OneField{
RelationModel: h.StockIncoterms(),
String: "Incoterm",
//states={'done': [('readonly', True)]}
Help: "International Commercial Terms are a series of predefined" + 
"commercial terms used in international transactions.",
},
"ProductId": models.Many2OneField{
RelationModel: h.ProductProduct(),
Related: `OrderLine.ProductId`,
String: "Product",
},
"CreateUid": models.Many2OneField{
RelationModel: h.User(),
String: "Responsible",
},
"CompanyId": models.Many2OneField{
RelationModel: h.Company(),
String: "Company",
Required: true,
Index: true,
//states=READONLY_STATES
Default: func (env models.Environment) interface{} { return env.Uid().company_id.id },
},
"PickingTypeId": models.Many2OneField{
RelationModel: h.StockPickingType(),
String: "Deliver To",
//states=READONLY_STATES
Required: true,
Default: models.DefaultValue(_default_picking_type),
Help: "This will determine picking type of incoming shipment",
},
"DefaultLocationDestIdUsage": models.SelectionField{
Related: `PickingTypeId.DefaultLocationDestId.Usage`,
String: "Destination Location Type",
Help: "Technical field used to display the Drop Ship Address",
ReadOnly: true,
},
"GroupId": models.Many2OneField{
RelationModel: h.ProcurementGroup(),
String: "Procurement Group",
NoCopy: true,
},
"IsShipped": models.BooleanField{
Compute: h.PurchaseOrder().Methods().ComputeIsShipped(),
},
})
h.PurchaseOrder().Methods().NameSearch().Extend(
`NameSearch`,
func(rs m.PurchaseOrderSet, name webdata.NameSearchParams, args interface{}, operator interface{}, limit interface{})  {
//        args = args or []
//        domain = []
//        if name:
//            domain = ['|', ('name', operator, name),
//                      ('partner_ref', operator, name)]
//        pos = self.search(domain + args, limit=limit)
//        return pos.name_get()
})
h.PurchaseOrder().Methods().NameGet().Extend(
`NameGet`,
func(rs m.PurchaseOrderSet)  {
//        result = []
//        for po in self:
//            name = po.name
//            if po.partner_ref:
//                name += ' ('+po.partner_ref+')'
//            if po.amount_total:
//                name += ': ' + \
//                    formatLang(self.env, po.amount_total,
//                               currency_obj=po.currency_id)
//            result.append((po.id, name))
//        return result
})
h.PurchaseOrder().Methods().Create().Extend(
`Create`,
func(rs m.PurchaseOrderSet, vals models.RecordData)  {
//        if vals.get('name', 'New') == 'New':
//            vals['name'] = self.env['ir.sequence'].next_by_code(
//                'purchase.order') or '/'
//        return super(PurchaseOrder, self).create(vals)
})
h.PurchaseOrder().Methods().Unlink().Extend(
`Unlink`,
func(rs m.PurchaseOrderSet)  {
//        for order in self:
//            if not order.state == 'cancel':
//                raise UserError(
//                    _('In order to delete a purchase order, you must cancel it first.'))
//        return super(PurchaseOrder, self).unlink()
})
h.PurchaseOrder().Methods().Copy().Extend(
`Copy`,
func(rs m.PurchaseOrderSet, defaultName models.RecordData)  {
//        new_po = super(PurchaseOrder, self).copy(default=default)
//        for line in new_po.order_line:
//            seller = line.product_id._select_seller(
//                partner_id=line.partner_id, quantity=line.product_qty,
//                date=line.order_id.date_order and line.order_id.date_order[:10], uom_id=line.product_uom)
//            line.date_planned = line._get_date_planned(seller)
//        return new_po
})
h.PurchaseOrder().Methods().TrackSubtype().DeclareMethod(
`TrackSubtype`,
func(rs m.PurchaseOrderSet, init_values interface{})  {
//        self.ensure_one()
//        if 'state' in init_values and self.state == 'purchase':
//            return 'purchase.mt_rfq_approved'
//        elif 'state' in init_values and self.state == 'to approve':
//            return 'purchase.mt_rfq_confirmed'
//        elif 'state' in init_values and self.state == 'done':
//            return 'purchase.mt_rfq_done'
//        return super(PurchaseOrder, self)._track_subtype(init_values)
})
h.PurchaseOrder().Methods().OnchangePartnerId().DeclareMethod(
`OnchangePartnerId`,
func(rs m.PurchaseOrderSet)  {
//        if not self.partner_id:
//            self.fiscal_position_id = False
//            self.payment_term_id = False
//            self.currency_id = False
//        else:
//            self.fiscal_position_id = self.env['account.fiscal.position'].with_context(
//                company_id=self.company_id.id).get_fiscal_position(self.partner_id.id)
//            self.payment_term_id = self.partner_id.property_supplier_payment_term_id.id
//            self.currency_id = self.partner_id.property_purchase_currency_id.id or self.env.user.company_id.currency_id.id
//        return {}
})
h.PurchaseOrder().Methods().ComputeTaxId().DeclareMethod(
`
        Trigger the recompute of the taxes if the fiscal
position is changed on the PO.
        `,
func(rs m.PurchaseOrderSet)  {
//        for order in self:
//            order.order_line._compute_tax_id()
})
h.PurchaseOrder().Methods().OnchangePartnerIdWarning().DeclareMethod(
`OnchangePartnerIdWarning`,
func(rs m.PurchaseOrderSet)  {
//        if not self.partner_id:
//            return
//        warning = {}
//        title = False
//        message = False
//        partner = self.partner_id
//        if partner.purchase_warn == 'no-message' and partner.parent_id:
//            partner = partner.parent_id
//        if partner.purchase_warn != 'no-message':
//            # Block if partner only has warning but parent company is blocked
//            if partner.purchase_warn != 'block' and partner.parent_id and partner.parent_id.purchase_warn == 'block':
//                partner = partner.parent_id
//            title = _("Warning for %s") % partner.name
//            message = partner.purchase_warn_msg
//            warning = {
//                'title': title,
//                'message': message
//            }
//            if partner.purchase_warn == 'block':
//                self.update({'partner_id': False})
//            return {'warning': warning}
//        return {}
})
h.PurchaseOrder().Methods().OnchangePickingTypeId().DeclareMethod(
`OnchangePickingTypeId`,
func(rs m.PurchaseOrderSet)  {
//        if self.picking_type_id.default_location_dest_id.usage != 'customer':
//            self.dest_address_id = False
})
h.PurchaseOrder().Methods().ActionRfqSend().DeclareMethod(
`
        This function opens a window to compose an email,
with the edi purchase template message loaded by default
        `,
func(rs m.PurchaseOrderSet)  {
//        self.ensure_one()
//        ir_model_data = self.env['ir.model.data']
//        try:
//            if self.env.context.get('send_rfq', False):
//                template_id = ir_model_data.get_object_reference(
//                    'purchase', 'email_template_edi_purchase')[1]
//            else:
//                template_id = ir_model_data.get_object_reference(
//                    'purchase', 'email_template_edi_purchase_done')[1]
//        except ValueError:
//            template_id = False
//        try:
//            compose_form_id = ir_model_data.get_object_reference(
//                'mail', 'email_compose_message_wizard_form')[1]
//        except ValueError:
//            compose_form_id = False
//        ctx = dict(self.env.context or {})
//        ctx.update({
//            'default_model': 'purchase.order',
//            'default_res_id': self.ids[0],
//            'default_use_template': bool(template_id),
//            'default_template_id': template_id,
//            'default_composition_mode': 'comment',
//            'purchase_mark_rfq_sent': True,
//        })
//        return {
//            'name': _('Compose Email'),
//            'type': 'ir.actions.act_window',
//            'view_type': 'form',
//            'view_mode': 'form',
//            'res_model': 'mail.compose.message',
//            'views': [(compose_form_id, 'form')],
//            'view_id': compose_form_id,
//            'target': 'new',
//            'context': ctx,
//        }
})
h.PurchaseOrder().Methods().PrintQuotation().DeclareMethod(
`PrintQuotation`,
func(rs m.PurchaseOrderSet)  {
//        self.write({'state': "sent"})
//        return self.env['report'].get_action(self, 'purchase.report_purchasequotation')
})
h.PurchaseOrder().Methods().ButtonApprove().DeclareMethod(
`ButtonApprove`,
func(rs m.PurchaseOrderSet, force interface{})  {
//        self.write(
//            {'state': 'purchase', 'date_approve': fields.Date.context_today(self)})
//        self._create_picking()
//        self.filtered(
//            lambda p: p.company_id.po_lock == 'lock').write({'state': 'done'})
//        return {}
})
h.PurchaseOrder().Methods().ButtonDraft().DeclareMethod(
`ButtonDraft`,
func(rs m.PurchaseOrderSet)  {
//        self.write({'state': 'draft'})
//        return {}
})
h.PurchaseOrder().Methods().ButtonConfirm().DeclareMethod(
`ButtonConfirm`,
func(rs m.PurchaseOrderSet)  {
//        for order in self:
//            if order.state not in ['draft', 'sent']:
//                continue
//            order._add_supplier_to_product()
//            # Deal with double validation process
//            if order.company_id.po_double_validation == 'one_step'\
//                    or (order.company_id.po_double_validation == 'two_step'
//                        and order.amount_total < self.env.user.company_id.currency_id.compute(order.company_id.po_double_validation_amount, order.currency_id))\
//                    or order.user_has_groups('purchase.group_purchase_manager'):
//                order.button_approve()
//            else:
//                order.write({'state': 'to approve'})
//        return True
})
h.PurchaseOrder().Methods().ButtonCancel().DeclareMethod(
`ButtonCancel`,
func(rs m.PurchaseOrderSet)  {
//        for order in self:
//            for pick in order.picking_ids:
//                if pick.state == 'done':
//                    raise UserError(
//                        _('Unable to cancel purchase order %s as some receptions have already been done.') % (order.name))
//            for inv in order.invoice_ids:
//                if inv and inv.state not in ('cancel', 'draft'):
//                    raise UserError(
//                        _("Unable to cancel this purchase order. You must first cancel related vendor bills."))
//
//            for pick in order.picking_ids.filtered(lambda r: r.state != 'cancel'):
//                pick.action_cancel()
//            # TDE FIXME: I don' think context key is necessary, as actions are not related / called from each other
//            if not self.env.context.get('cancel_procurement'):
//                procurements = order.order_line.mapped('procurement_ids')
//                procurements.filtered(lambda r: r.state not in (
//                    'cancel', 'exception') and r.rule_id.propagate).write({'state': 'cancel'})
//                procurements.filtered(lambda r: r.state not in (
//                    'cancel', 'exception') and not r.rule_id.propagate).write({'state': 'exception'})
//                moves = procurements.filtered(
//                    lambda r: r.rule_id.propagate).mapped('move_dest_id')
//                moves.filtered(lambda r: r.state != 'cancel').action_cancel()
//        self.write({'state': 'cancel'})
})
h.PurchaseOrder().Methods().ButtonUnlock().DeclareMethod(
`ButtonUnlock`,
func(rs m.PurchaseOrderSet)  {
//        self.write({'state': 'purchase'})
})
h.PurchaseOrder().Methods().ButtonDone().DeclareMethod(
`ButtonDone`,
func(rs m.PurchaseOrderSet)  {
//        self.write({'state': 'done'})
})
h.PurchaseOrder().Methods().GetDestinationLocation().DeclareMethod(
`GetDestinationLocation`,
func(rs m.PurchaseOrderSet)  {
//        self.ensure_one()
//        if self.dest_address_id:
//            return self.dest_address_id.property_stock_customer.id
//        return self.picking_type_id.default_location_dest_id.id
})
h.PurchaseOrder().Methods().PreparePicking().DeclareMethod(
`PreparePicking`,
func(rs m.PurchaseOrderSet)  {
//        if not self.group_id:
//            self.group_id = self.group_id.create({
//                'name': self.name,
//                'partner_id': self.partner_id.id
//            })
//        if not self.partner_id.property_stock_supplier.id:
//            raise UserError(
//                _("You must set a Vendor Location for this partner %s") % self.partner_id.name)
//        return {
//            'picking_type_id': self.picking_type_id.id,
//            'partner_id': self.partner_id.id,
//            'date': self.date_order,
//            'origin': self.name,
//            'location_dest_id': self._get_destination_location(),
//            'location_id': self.partner_id.property_stock_supplier.id,
//            'company_id': self.company_id.id,
//        }
})
h.PurchaseOrder().Methods().CreatePicking().DeclareMethod(
`CreatePicking`,
func(rs m.PurchaseOrderSet)  {
//        StockPicking = self.env['stock.picking']
//        for order in self:
//            if any([ptype in ['product', 'consu'] for ptype in order.order_line.mapped('product_id.type')]):
//                pickings = order.picking_ids.filtered(
//                    lambda x: x.state not in ('done', 'cancel'))
//                if not pickings:
//                    res = order._prepare_picking()
//                    picking = StockPicking.create(res)
//                else:
//                    picking = pickings[0]
//                moves = order.order_line._create_stock_moves(picking)
//                moves = moves.filtered(lambda x: x.state not in (
//                    'done', 'cancel')).action_confirm()
//                seq = 0
//                for move in sorted(moves, key=lambda move: move.date_expected):
//                    seq += 5
//                    move.sequence = seq
//                moves.force_assign()
//                picking.message_post_with_view('mail.message_origin_link',
//                                               values={'self': picking,
//                                                       'origin': order},
//                                               subtype_id=self.env.ref('mail.mt_note').id)
//        return True
})
h.PurchaseOrder().Methods().AddSupplierToProduct().DeclareMethod(
`AddSupplierToProduct`,
func(rs m.PurchaseOrderSet)  {
//        for line in self.order_line:
//            # Do not add a contact as a supplier
//            partner = self.partner_id if not self.partner_id.parent_id else self.partner_id.parent_id
//            if partner not in line.product_id.seller_ids.mapped('name') and len(line.product_id.seller_ids) <= 10:
//                currency = partner.property_purchase_currency_id or self.env.user.company_id.currency_id
//                supplierinfo = {
//                    'name': partner.id,
//                    'sequence': max(line.product_id.seller_ids.mapped('sequence')) + 1 if line.product_id.seller_ids else 1,
//                    'product_uom': line.product_uom.id,
//                    'min_qty': 0.0,
//                    'price': self.currency_id.compute(line.price_unit, currency),
//                    'currency_id': currency.id,
//                    'delay': 0,
//                }
//                # In case the order partner is a contact address, a new supplierinfo is created on
//                # the parent company. In this case, we keep the product name and code.
//                seller = line.product_id._select_seller(
//                    partner_id=line.partner_id,
//                    quantity=line.product_qty,
//                    date=line.order_id.date_order and line.order_id.date_order[:10],
//                    uom_id=line.product_uom)
//                if seller:
//                    supplierinfo['product_name'] = seller.product_name
//                    supplierinfo['product_code'] = seller.product_code
//                vals = {
//                    'seller_ids': [(0, 0, supplierinfo)],
//                }
//                try:
//                    line.product_id.write(vals)
//                except AccessError:  # no write access rights -> just ignore
//                    break
})
h.PurchaseOrder().Methods().ActionViewPicking().DeclareMethod(
`
        This function returns an action that display existing
picking orders of given purchase order ids.
        When only one found, show the picking immediately.
        `,
func(rs m.PurchaseOrderSet)  {
//        action = self.env.ref('stock.action_picking_tree')
//        result = action.read()[0]
//        result.pop('id', None)
//        result['context'] = {}
//        pick_ids = sum([order.picking_ids.ids for order in self], [])
//        if len(pick_ids) > 1:
//            result['domain'] = "[('id','in',[" + \
//                ','.join(map(str, pick_ids)) + "])]"
//        elif len(pick_ids) == 1:
//            res = self.env.ref('stock.view_picking_form', False)
//            result['views'] = [(res and res.id or False, 'form')]
//            result['res_id'] = pick_ids and pick_ids[0] or False
//        return result
})
h.PurchaseOrder().Methods().ActionViewInvoice().DeclareMethod(
`
        This function returns an action that display existing
vendor bills of given purchase order ids.
        When only one found, show the vendor bill immediately.
        `,
func(rs m.PurchaseOrderSet)  {
//        action = self.env.ref('account.action_invoice_tree2')
//        result = action.read()[0]
//        result['context'] = {'type': 'in_invoice',
//                             'default_purchase_id': self.id}
//        if not self.invoice_ids:
//            # Choose a default account journal in the same currency in case a new invoice is created
//            journal_domain = [
//                ('type', '=', 'purchase'),
//                ('company_id', '=', self.company_id.id),
//                ('currency_id', '=', self.currency_id.id),
//            ]
//            default_journal_id = self.env['account.journal'].search(
//                journal_domain, limit=1)
//            if default_journal_id:
//                result['context']['default_journal_id'] = default_journal_id.id
//        else:
//            # Use the same account journal than a previous invoice
//            result['context']['default_journal_id'] = self.invoice_ids[0].journal_id.id
//        if len(self.invoice_ids) != 1:
//            result['domain'] = "[('id', 'in', " + \
//                str(self.invoice_ids.ids) + ")]"
//        elif len(self.invoice_ids) == 1:
//            res = self.env.ref('account.invoice_supplier_form', False)
//            result['views'] = [(res and res.id or False, 'form')]
//            result['res_id'] = self.invoice_ids.id
//        return result
})
h.PurchaseOrder().Methods().ActionSetDatePlanned().DeclareMethod(
`ActionSetDatePlanned`,
func(rs m.PurchaseOrderSet)  {
//        for order in self:
//            order.order_line.update({'date_planned': order.date_planned})
})
h.PurchaseOrderLine().DeclareModel()



h.PurchaseOrderLine().Methods().ComputeAmount().DeclareMethod(
`ComputeAmount`,
func(rs m.PurchaseOrderLineSet)  {
//        for line in self:
//            taxes = line.taxes_id.compute_all(line.price_unit, line.order_id.currency_id,
//                                              line.product_qty, product=line.product_id, partner=line.order_id.partner_id)
//            line.update({
//                'price_tax': taxes['total_included'] - taxes['total_excluded'],
//                'price_total': taxes['total_included'],
//                'price_subtotal': taxes['total_excluded'],
//            })
})
h.PurchaseOrderLine().Methods().ComputeTaxId().DeclareMethod(
`ComputeTaxId`,
func(rs m.PurchaseOrderLineSet)  {
//        for line in self:
//            fpos = line.order_id.fiscal_position_id or line.order_id.partner_id.property_account_position_id
//            # If company_id is set, always filter taxes by the company
//            taxes = line.product_id.supplier_taxes_id.filtered(
//                lambda r: not line.company_id or r.company_id == line.company_id)
//            line.taxes_id = fpos.map_tax(
//                taxes, line.product_id, line.order_id.partner_id) if fpos else taxes
})
h.PurchaseOrderLine().Methods().ComputeQtyInvoiced().DeclareMethod(
`ComputeQtyInvoiced`,
func(rs m.PurchaseOrderLineSet)  {
//        for line in self:
//            qty = 0.0
//            for inv_line in line.invoice_lines:
//                if inv_line.invoice_id.state not in ['cancel']:
//                    if inv_line.invoice_id.type == 'in_invoice':
//                        qty += inv_line.uom_id._compute_quantity(
//                            inv_line.quantity, line.product_uom)
//                    elif inv_line.invoice_id.type == 'in_refund':
//                        qty -= inv_line.uom_id._compute_quantity(
//                            inv_line.quantity, line.product_uom)
//            line.qty_invoiced = qty
})
h.PurchaseOrderLine().Methods().ComputeQtyReceived().DeclareMethod(
`ComputeQtyReceived`,
func(rs m.PurchaseOrderLineSet)  {
//        for line in self:
//            if line.order_id.state not in ['purchase', 'done']:
//                line.qty_received = 0.0
//                continue
//            if line.product_id.type not in ['consu', 'product']:
//                line.qty_received = line.product_qty
//                continue
//            total = 0.0
//            for move in line.move_ids:
//                if move.state == 'done':
//                    if move.product_uom != line.product_uom:
//                        total += move.product_uom._compute_quantity(
//                            move.product_uom_qty, line.product_uom)
//                    else:
//                        total += move.product_uom_qty
//            line.qty_received = total
})
h.PurchaseOrderLine().Methods().Create().Extend(
`Create`,
func(rs m.PurchaseOrderLineSet, values models.RecordData)  {
//        line = super(PurchaseOrderLine, self).create(values)
//        if line.order_id.state == 'purchase':
//            line.order_id._create_picking()
//            msg = _("Extra line with %s ") % (line.product_id.display_name)
//            line.order_id.message_post(body=msg)
//        return line
})
h.PurchaseOrderLine().Methods().Write().Extend(
`Write`,
func(rs m.PurchaseOrderLineSet, values models.RecordData)  {
//        orders = False
//        if 'product_qty' in values:
//            changed_lines = self.filtered(
//                lambda x: x.order_id.state == 'purchase')
//            if changed_lines:
//                orders = changed_lines.mapped('order_id')
//                for order in orders:
//                    order_lines = changed_lines.filtered(
//                        lambda x: x.order_id == order)
//                    msg = ""
//                    if any([values['product_qty'] < x.product_qty for x in order_lines]):
//                        msg += "<b>" + \
//                            _('The ordered quantity has been decreased. Do not forget to take it into account on your bills and receipts.') + '</b><br/>'
//                    msg += "<ul>"
//                    for line in order_lines:
//                        msg += "<li> %s:" % (line.product_id.display_name)
//                        msg += "<br/>" + _("Ordered Quantity") + ": %s -> %s <br/>" % (
//                            line.product_qty, float(values['product_qty']))
//                        if line.product_id.type in ('product', 'consu'):
//                            msg += _("Received Quantity") + \
//                                ": %s <br/>" % (line.qty_received)
//                        msg += _("Billed Quantity") + \
//                            ": %s <br/></li>" % (line.qty_invoiced)
//                    msg += "</ul>"
//                    order.message_post(body=msg)
//        if 'date_planned' in values:
//            self.env['stock.move'].search([
//                ('purchase_line_id', 'in', self.ids), ('state', '!=', 'done')
//            ]).write({'date_expected': values['date_planned']})
//        result = super(PurchaseOrderLine, self).write(values)
//        if orders:
//            orders._create_picking()
//        return result
})
h.PurchaseOrderLine().AddFields(map[string]models.FieldDefinition{
"Name": models.TextField{
String: "Description",
Required: true,
},
"Sequence": models.IntegerField{
String: "Sequence",
Default: models.DefaultValue(10),
},
"ProductQty": models.FloatField{
String: "Quantity",
//digits=dp.get_precision('Product Unit of Measure')
Required: true,
},
"DatePlanned": models.DateTimeField{
String: "Scheduled Date",
Required: true,
Index: true,
},
"TaxesId": models.Many2ManyField{
RelationModel: h.AccountTax(),
String: "Taxes",
Filter: q.Active().Equals(False).Or().Active().Equals(True),
},
"ProductUom": models.Many2OneField{
RelationModel: h.ProductUom(),
String: "Product Unit of Measure",
Required: true,
},
"ProductId": models.Many2OneField{
RelationModel: h.ProductProduct(),
String: "Product",
Filter: q.PurchaseOk().Equals(True),
//change_default=True
Required: true,
},
"MoveIds": models.One2ManyField{
RelationModel: h.StockMove(),
ReverseFK: "",
String: "Reservation",
ReadOnly: true,
OnDelete: `set null`,
NoCopy: true,
},
"PriceUnit": models.FloatField{
String: "Unit Price",
Required: true,
//digits=dp.get_precision('Product Price')
},
"PriceSubtotal": models.MonetaryField{
Compute: h.PurchaseOrderLine().Methods().ComputeAmount(),
String: "Subtotal",
Stored: true,
},
"PriceTotal": models.MonetaryField{
Compute: h.PurchaseOrderLine().Methods().ComputeAmount(),
String: "Total",
Stored: true,
},
"PriceTax": models.MonetaryField{
Compute: h.PurchaseOrderLine().Methods().ComputeAmount(),
String: "Tax",
Stored: true,
},
"OrderId": models.Many2OneField{
RelationModel: h.PurchaseOrder(),
String: "Order Reference",
Index: true,
Required: true,
OnDelete: `cascade`,
},
"AccountAnalyticId": models.Many2OneField{
RelationModel: h.AccountAnalyticAccount(),
String: "Analytic Account",
},
"AnalyticTagIds": models.Many2ManyField{
RelationModel: h.AccountAnalyticTag(),
String: "Analytic Tags",
},
"CompanyId": models.Many2OneField{
RelationModel: h.Company(),
Related: `OrderId.CompanyId`,
String: "Company",
Stored: true,
ReadOnly: true,
},
"State": models.SelectionField{
Related: `OrderId.State`,
Stored: true,
},
"InvoiceLines": models.One2ManyField{
RelationModel: h.AccountInvoiceLine(),
ReverseFK: "",
String: "Bill Lines",
ReadOnly: true,
NoCopy: true,
},
"QtyInvoiced": models.FloatField{
Compute: h.PurchaseOrderLine().Methods().ComputeQtyInvoiced(),
String: "Billed Qty",
//digits=dp.get_precision('Product Unit of Measure')
Stored: true,
},
"QtyReceived": models.FloatField{
Compute: h.PurchaseOrderLine().Methods().ComputeQtyReceived(),
String: "Received Qty",
//digits=dp.get_precision('Product Unit of Measure')
Stored: true,
},
"PartnerId": models.Many2OneField{
RelationModel: h.Partner(),
Related: `OrderId.PartnerId`,
String: "Partner",
ReadOnly: true,
Stored: true,
},
"CurrencyId": models.Many2OneField{
Related: `OrderId.CurrencyId`,
Stored: true,
String: "Currency",
ReadOnly: true,
},
"DateOrder": models.DateTimeField{
Related: `OrderId.DateOrder`,
String: "Order Date",
ReadOnly: true,
},
"ProcurementIds": models.One2ManyField{
RelationModel: h.ProcurementOrder(),
ReverseFK: "",
String: "Associated Procurements",
NoCopy: true,
},
})
h.PurchaseOrderLine().Methods().GetStockMovePriceUnit().DeclareMethod(
`GetStockMovePriceUnit`,
func(rs m.PurchaseOrderLineSet)  {
//        self.ensure_one()
//        line = self[0]
//        order = line.order_id
//        price_unit = line.price_unit
//        if line.taxes_id:
//            price_unit = line.taxes_id.with_context(round=False).compute_all(
//                price_unit, currency=line.order_id.currency_id, quantity=1.0, product=line.product_id, partner=line.order_id.partner_id
//            )['total_excluded']
//        if line.product_uom.id != line.product_id.uom_id.id:
//            price_unit *= line.product_uom.factor / line.product_id.uom_id.factor
//        if order.currency_id != order.company_id.currency_id:
//            price_unit = order.currency_id.compute(
//                price_unit, order.company_id.currency_id, round=False)
//        return price_unit
})
h.PurchaseOrderLine().Methods().PrepareStockMoves().DeclareMethod(
` Prepare the stock moves data for one order line. This
function returns a list of
        dictionary ready to be used in stock.move's create()
        `,
func(rs m.PurchaseOrderLineSet, picking interface{})  {
//        self.ensure_one()
//        res = []
//        if self.product_id.type not in ['product', 'consu']:
//            return res
//        qty = 0.0
//        price_unit = self._get_stock_move_price_unit()
//        for move in self.move_ids.filtered(lambda x: x.state != 'cancel'):
//            qty += move.product_qty
//        template = {
//            'name': self.name or '',
//            'product_id': self.product_id.id,
//            'product_uom': self.product_uom.id,
//            'date': self.order_id.date_order,
//            'date_expected': self.date_planned,
//            'location_id': self.order_id.partner_id.property_stock_supplier.id,
//            'location_dest_id': self.order_id._get_destination_location(),
//            'picking_id': picking.id,
//            'partner_id': self.order_id.dest_address_id.id,
//            'move_dest_id': False,
//            'state': 'draft',
//            'purchase_line_id': self.id,
//            'company_id': self.order_id.company_id.id,
//            'price_unit': price_unit,
//            'picking_type_id': self.order_id.picking_type_id.id,
//            'group_id': self.order_id.group_id.id,
//            'procurement_id': False,
//            'origin': self.order_id.name,
//            'route_ids': self.order_id.picking_type_id.warehouse_id and [(6, 0, [x.id for x in self.order_id.picking_type_id.warehouse_id.route_ids])] or [],
//            'warehouse_id': self.order_id.picking_type_id.warehouse_id.id,
//        }
//        diff_quantity = self.product_qty - qty
//        for procurement in self.procurement_ids.filtered(lambda p: p.state != 'cancel'):
//            # If the procurement has some moves already, we should deduct their quantity
//            sum_existing_moves = sum(
//                x.product_qty for x in procurement.move_ids if x.state != 'cancel')
//            existing_proc_qty = procurement.product_id.uom_id._compute_quantity(
//                sum_existing_moves, procurement.product_uom)
//            procurement_qty = procurement.product_uom._compute_quantity(
//                procurement.product_qty, self.product_uom) - existing_proc_qty
//            if float_compare(procurement_qty, 0.0, precision_rounding=procurement.product_uom.rounding) > 0 and float_compare(diff_quantity, 0.0, precision_rounding=self.product_uom.rounding) > 0:
//                tmp = template.copy()
//                tmp.update({
//                    'product_uom_qty': min(procurement_qty, diff_quantity),
//                    # move destination is same as procurement destination
//                    'move_dest_id': procurement.move_dest_id.id,
//                    'procurement_id': procurement.id,
//                    'propagate': procurement.rule_id.propagate,
//                })
//                res.append(tmp)
//                diff_quantity -= min(procurement_qty, diff_quantity)
//        if float_compare(diff_quantity, 0.0,  precision_rounding=self.product_uom.rounding) > 0:
//            template['product_uom_qty'] = diff_quantity
//            res.append(template)
//        return res
})
h.PurchaseOrderLine().Methods().CreateStockMoves().DeclareMethod(
`CreateStockMoves`,
func(rs m.PurchaseOrderLineSet, picking interface{})  {
//        moves = self.env['stock.move']
//        done = self.env['stock.move'].browse()
//        for line in self:
//            for val in line._prepare_stock_moves(picking):
//                done += moves.create(val)
//        return done
})
h.PurchaseOrderLine().Methods().Unlink().Extend(
`Unlink`,
func(rs m.PurchaseOrderLineSet)  {
//        for line in self:
//            if line.order_id.state in ['purchase', 'done']:
//                raise UserError(
//                    _('Cannot delete a purchase order line which is in state \'%s\'.') % (line.state))
//            for proc in line.procurement_ids:
//                proc.message_post(body=_('Purchase order line deleted.'))
//            line.procurement_ids.filtered(
//                lambda r: r.state != 'cancel').write({'state': 'exception'})
//        return super(PurchaseOrderLine, self).unlink()
})
h.PurchaseOrderLine().Methods().GetDatePlanned().DeclareMethod(
`Return the datetime value to use as Schedule Date (``date_planned``) for
           PO Lines that correspond to the given product.seller_ids,
           when ordered at `date_order_str`.

           :param browse_record | False product: product.product,
used to
               determine delivery delay thanks to the selected
seller field (if False, default delay = 0)
           :param browse_record | False po: purchase.order,
necessary only if
               the PO line is not yet attached to a PO.
           :rtype: datetime
           :return: desired Schedule Date for the PO line
        `,
func(rs m.PurchaseOrderLineSet, seller interface{}, po interface{})  {
//        date_order = po.date_order if po else self.order_id.date_order
//        if date_order:
//            return datetime.strptime(date_order, DEFAULT_SERVER_DATETIME_FORMAT) + relativedelta(days=seller.delay if seller else 0)
//        else:
//            return datetime.today() + relativedelta(days=seller.delay if seller else 0)
})
h.PurchaseOrderLine().Methods().OnchangeProductId().DeclareMethod(
`OnchangeProductId`,
func(rs m.PurchaseOrderLineSet)  {
//        result = {}
//        if not self.product_id:
//            return result
//        self.date_planned = datetime.today().strftime(DEFAULT_SERVER_DATETIME_FORMAT)
//        self.price_unit = self.product_qty = 0.0
//        self.product_uom = self.product_id.uom_po_id or self.product_id.uom_id
//        result['domain'] = {'product_uom': [
//            ('category_id', '=', self.product_id.uom_id.category_id.id)]}
//        product_lang = self.product_id.with_context(
//            lang=self.partner_id.lang,
//            partner_id=self.partner_id.id)
//        self.name = product_lang.display_name
//        if product_lang.description_purchase:
//            self.name += '\n' + product_lang.description_purchase
//        fpos = self.order_id.fiscal_position_id
//        if self.env.uid == SUPERUSER_ID:
//            company_id = self.env.user.company_id.id
//            self.taxes_id = fpos.map_tax(self.product_id.supplier_taxes_id.filtered(
//                lambda r: r.company_id.id == company_id))
//        else:
//            self.taxes_id = fpos.map_tax(self.product_id.supplier_taxes_id)
//        self._suggest_quantity()
//        self._onchange_quantity()
//        return result
})
h.PurchaseOrderLine().Methods().OnchangeProductIdWarning().DeclareMethod(
`OnchangeProductIdWarning`,
func(rs m.PurchaseOrderLineSet)  {
//        if not self.product_id:
//            return
//        warning = {}
//        title = False
//        message = False
//        product_info = self.product_id
//        if product_info.purchase_line_warn != 'no-message':
//            title = _("Warning for %s") % product_info.name
//            message = product_info.purchase_line_warn_msg
//            warning['title'] = title
//            warning['message'] = message
//            if product_info.purchase_line_warn == 'block':
//                self.product_id = False
//            return {'warning': warning}
//        return {}
})
h.PurchaseOrderLine().Methods().OnchangeQuantity().DeclareMethod(
`OnchangeQuantity`,
func(rs m.PurchaseOrderLineSet)  {
//        if not self.product_id:
//            return
//        seller = self.product_id._select_seller(
//            partner_id=self.partner_id,
//            quantity=self.product_qty,
//            date=self.order_id.date_order and self.order_id.date_order[:10],
//            uom_id=self.product_uom)
//        if seller or not self.date_planned:
//            self.date_planned = self._get_date_planned(
//                seller).strftime(DEFAULT_SERVER_DATETIME_FORMAT)
//        if not seller:
//            return
//        price_unit = self.env['account.tax']._fix_tax_included_price_company(
//            seller.price, self.product_id.supplier_taxes_id, self.taxes_id, self.company_id) if seller else 0.0
//        if price_unit and seller and self.order_id.currency_id and seller.currency_id != self.order_id.currency_id:
//            price_unit = seller.currency_id.compute(
//                price_unit, self.order_id.currency_id)
//        if seller and self.product_uom and seller.product_uom != self.product_uom:
//            price_unit = seller.product_uom._compute_price(
//                price_unit, self.product_uom)
//        self.price_unit = price_unit
})
h.PurchaseOrderLine().Methods().OnchangeProductQty().DeclareMethod(
`OnchangeProductQty`,
func(rs m.PurchaseOrderLineSet)  {
//        if (self.state == 'purchase' or self.state == 'to approve') and self.product_id.type in ['product', 'consu'] and self.product_qty < self._origin.product_qty:
//            warning_mess = {
//                'title': _('Ordered quantity decreased!'),
//                'message': _('You are decreasing the ordered quantity!\nYou must update the quantities on the reception and/or bills.'),
//            }
//            return {'warning': warning_mess}
})
h.PurchaseOrderLine().Methods().SuggestQuantity().DeclareMethod(
`
        Suggest a minimal quantity based on the seller
        `,
func(rs m.PurchaseOrderLineSet)  {
//        if not self.product_id:
//            return
//        seller_min_qty = self.product_id.seller_ids\
//            .filtered(lambda r: r.name == self.order_id.partner_id)\
//            .sorted(key=lambda r: r.min_qty)
//        if seller_min_qty:
//            self.product_qty = seller_min_qty[0].min_qty or 1.0
//            self.product_uom = seller_min_qty[0].product_uom
//        else:
//            self.product_qty = 1.0
})
h.ProcurementRule().DeclareModel()

h.ProcurementRule().Methods().GetAction().DeclareMethod(
`GetAction`,
func(rs m.ProcurementRuleSet)  {
//        return [('buy', _('Buy'))] + super(ProcurementRule, self)._get_action()
})
h.ProcurementOrder().DeclareModel()

h.ProcurementOrder().AddFields(map[string]models.FieldDefinition{
"PurchaseLineId": models.Many2OneField{
RelationModel: h.PurchaseOrderLine(),
String: "Purchase Order Line",
NoCopy: true,
},
"PurchaseId": models.Many2OneField{
Related: `PurchaseLineId.OrderId`,
String: "Purchase Order",
},
})
h.ProcurementOrder().Methods().PropagateCancels().DeclareMethod(
`PropagateCancels`,
func(rs m.ProcurementOrderSet)  {
//        result = super(ProcurementOrder, self).propagate_cancels()
//        for procurement in self:
//            if procurement.rule_id.action == 'buy' and procurement.purchase_line_id:
//                if procurement.purchase_line_id.order_id.state not in ('draft', 'cancel', 'sent', 'to validate'):
//                    raise UserError(
//                        _('Can not cancel a procurement related to a purchase order. Please cancel the purchase order first.'))
//            if procurement.purchase_line_id:
//                price_unit = 0.0
//                product_qty = 0.0
//                others_procs = procurement.purchase_line_id.procurement_ids.filtered(
//                    lambda r: r != procurement)
//                for other_proc in others_procs:
//                    if other_proc.state not in ['cancel', 'draft']:
//                        product_qty += other_proc.product_uom._compute_quantity(
//                            other_proc.product_qty, procurement.purchase_line_id.product_uom)
//
//                precision = self.env['decimal.precision'].precision_get(
//                    'Product Unit of Measure')
//                if not float_is_zero(product_qty, precision_digits=precision):
//                    seller = procurement.product_id._select_seller(
//                        partner_id=procurement.purchase_line_id.partner_id,
//                        quantity=product_qty,
//                        date=procurement.purchase_line_id.order_id.date_order and procurement.purchase_line_id.order_id.date_order[
//                            :10],
//                        uom_id=procurement.purchase_line_id.product_uom)
//
//                    price_unit = self.env['account.tax']._fix_tax_included_price_company(
//                        seller.price, procurement.purchase_line_id.product_id.supplier_taxes_id, procurement.purchase_line_id.taxes_id, procurement.company_id) if seller else 0.0
//                    if price_unit and seller and procurement.purchase_line_id.order_id.currency_id and seller.currency_id != procurement.purchase_line_id.order_id.currency_id:
//                        price_unit = seller.currency_id.compute(
//                            price_unit, procurement.purchase_line_id.order_id.currency_id)
//
//                    if seller and seller.product_uom != procurement.purchase_line_id.product_uom:
//                        price_unit = seller.product_uom._compute_price(
//                            price_unit, procurement.purchase_line_id.product_uom)
//
//                    procurement.purchase_line_id.product_qty = product_qty
//                    procurement.purchase_line_id.price_unit = price_unit
//                else:
//                    procurement.purchase_line_id.unlink()
//        return result
})
h.ProcurementOrder().Methods().Run().DeclareMethod(
`Run`,
func(rs m.ProcurementOrderSet)  {
//        if self.rule_id and self.rule_id.action == 'buy':
//            return self.make_po()
//        return super(ProcurementOrder, self)._run()
})
h.ProcurementOrder().Methods().Check().DeclareMethod(
`Check`,
func(rs m.ProcurementOrderSet)  {
//        if self.rule_id.action == 'buy':
//            # In case Phantom BoM splits only into procurements
//            if not self.move_ids:
//                if self.purchase_line_id and self.purchase_line_id.order_id.state not in ('purchase', 'done', 'cancel'):
//                    return False
//                else:
//                    return True
//            move_all_done_or_cancel = all(
//                move.state in ['done', 'cancel'] for move in self.move_ids)
//            move_all_cancel = all(
//                move.state == 'cancel' for move in self.move_ids)
//            if not move_all_done_or_cancel:
//                return False
//            elif move_all_done_or_cancel and not move_all_cancel:
//                return True
//            else:
//                self.message_post(
//                    body=_('All stock moves have been cancelled for this procurement.'))
//                self.write({'state': 'cancel'})
//                return False
//        return super(ProcurementOrder, self)._check()
})
h.ProcurementOrder().Methods().GetPurchaseScheduleDate().DeclareMethod(
`Return the datetime value to use as Schedule Date (``date_planned``)
for the
           Purchase Order Lines created to satisfy the
given procurement. `,
func(rs m.ProcurementOrderSet)  {
//        procurement_date_planned = datetime.strptime(
//            self.date_planned, DEFAULT_SERVER_DATETIME_FORMAT)
//        schedule_date = (procurement_date_planned -
//                         relativedelta(days=self.company_id.po_lead))
//        return schedule_date
})
h.ProcurementOrder().Methods().GetPurchaseOrderDate().DeclareMethod(
`Return the datetime value to use as Order Date (``date_order``) for the
           Purchase Order created to satisfy the given procurement. `,
func(rs m.ProcurementOrderSet, schedule_date interface{})  {
//        self.ensure_one()
//        seller_delay = int(
//            self.product_id.with_context(force_company=self.company_id.id)._select_seller(
//                quantity=self.product_qty, uom_id=self.product_uom
//            ).delay
//        )
//        return schedule_date - relativedelta(days=seller_delay)
})
h.ProcurementOrder().Methods().PreparePurchaseOrderLine().DeclareMethod(
`PreparePurchaseOrderLine`,
func(rs m.ProcurementOrderSet, po interface{}, supplier interface{})  {
//        self.ensure_one()
//        procurement_uom_po_qty = self.product_uom._compute_quantity(
//            self.product_qty, self.product_id.uom_po_id)
//        seller = self.product_id.with_context(force_company=self.company_id.id)._select_seller(
//            partner_id=supplier.name,
//            quantity=procurement_uom_po_qty,
//            date=po.date_order and po.date_order[:10],
//            uom_id=self.product_id.uom_po_id)
//        taxes = self.product_id.supplier_taxes_id
//        fpos = po.fiscal_position_id
//        taxes_id = fpos.map_tax(taxes) if fpos else taxes
//        if taxes_id:
//            taxes_id = taxes_id.filtered(
//                lambda x: x.company_id.id == self.company_id.id)
//        price_unit = self.env['account.tax']._fix_tax_included_price_company(
//            seller.price, self.product_id.supplier_taxes_id, taxes_id, self.company_id) if seller else 0.0
//        if price_unit and seller and po.currency_id and seller.currency_id != po.currency_id:
//            price_unit = seller.currency_id.compute(price_unit, po.currency_id)
//        product_lang = self.product_id.with_context({
//            'lang': supplier.name.lang,
//            'partner_id': supplier.name.id,
//        })
//        name = product_lang.display_name
//        if product_lang.description_purchase:
//            name += '\n' + product_lang.description_purchase
//        date_planned = self.env['purchase.order.line']._get_date_planned(
//            seller, po=po).strftime(DEFAULT_SERVER_DATETIME_FORMAT)
//        return {
//            'name': name,
//            'product_qty': procurement_uom_po_qty,
//            'product_id': self.product_id.id,
//            'product_uom': self.product_id.uom_po_id.id,
//            'price_unit': price_unit,
//            'date_planned': date_planned,
//            'taxes_id': [(6, 0, taxes_id.ids)],
//            'procurement_ids': [(4, self.id)],
//            'order_id': po.id,
//        }
})
h.ProcurementOrder().Methods().PreparePurchaseOrder().DeclareMethod(
`PreparePurchaseOrder`,
func(rs m.ProcurementOrderSet, partner interface{})  {
//        self.ensure_one()
//        schedule_date = self._get_purchase_schedule_date()
//        purchase_date = self._get_purchase_order_date(schedule_date)
//        fpos = self.env['account.fiscal.position'].with_context(
//            force_company=self.company_id.id).get_fiscal_position(partner.id)
//        gpo = self.rule_id.group_propagation_option
//        group = (gpo == 'fixed' and self.rule_id.group_id.id) or \
//                (gpo == 'propagate' and self.group_id.id) or False
//        return {
//            'partner_id': partner.id,
//            'picking_type_id': self.rule_id.picking_type_id.id,
//            'company_id': self.company_id.id,
//            'currency_id': partner.property_purchase_currency_id.id or self.env.user.company_id.currency_id.id,
//            'dest_address_id': self.partner_dest_id.id,
//            'origin': self.origin,
//            'payment_term_id': partner.property_supplier_payment_term_id.id,
//            'date_order': purchase_date.strftime(DEFAULT_SERVER_DATETIME_FORMAT),
//            'fiscal_position_id': fpos,
//            'group_id': group
//        }
})
h.ProcurementOrder().Methods().MakePoSelectSupplier().DeclareMethod(
` Method intended to be overridden by customized modules
to implement any logic in the
            selection of supplier.
        `,
func(rs m.ProcurementOrderSet, suppliers interface{})  {
//        return suppliers[0]
})
h.ProcurementOrder().Methods().MakePoGetDomain().DeclareMethod(
`MakePoGetDomain`,
func(rs m.ProcurementOrderSet, partner interface{})  {
//        gpo = self.rule_id.group_propagation_option
//        group = (gpo == 'fixed' and self.rule_id.group_id) or \
//                (gpo == 'propagate' and self.group_id) or False
//        domain = (
//            ('partner_id', '=', partner.id),
//            ('state', '=', 'draft'),
//            ('picking_type_id', '=', self.rule_id.picking_type_id.id),
//            ('company_id', '=', self.company_id.id),
//            ('dest_address_id', '=', self.partner_dest_id.id))
//        if group:
//            domain += (('group_id', '=', group.id))
//        return domain
})
h.ProcurementOrder().Methods().MakePo().DeclareMethod(
`MakePo`,
func(rs m.ProcurementOrderSet)  {
//        cache = {}
//        res = []
//        for procurement in self:
//            suppliers = procurement.product_id.seller_ids\
//                .filtered(lambda r: (not r.company_id or r.company_id == procurement.company_id) and (not r.product_id or r.product_id == procurement.product_id))
//            if not suppliers:
//                procurement.message_post(body=_(
//                    'No vendor associated to product %s. Please set one to fix this procurement.') % (procurement.product_id.name))
//                continue
//            supplier = procurement._make_po_select_supplier(suppliers)
//            partner = supplier.name
//
//            domain = procurement._make_po_get_domain(partner)
//
//            if domain in cache:
//                po = cache[domain]
//            else:
//                po = self.env['purchase.order'].search([dom for dom in domain])
//                po = po[0] if po else False
//                cache[domain] = po
//            if not po:
//                vals = procurement._prepare_purchase_order(partner)
//                po = self.env['purchase.order'].create(vals)
//                name = (procurement.group_id and (procurement.group_id.name + ":")
//                        or "") + (procurement.name != "/" and procurement.name or "")
//                message = _(
//                    "This purchase order has been created from: <a href=# data-oe-model=procurement.order data-oe-id=%d>%s</a>") % (procurement.id, name)
//                po.message_post(body=message)
//                cache[domain] = po
//            elif not po.origin or procurement.origin not in po.origin.split(', '):
//                # Keep track of all procurements
//                if po.origin:
//                    if procurement.origin:
//                        po.write({'origin': po.origin +
//                                  ', ' + procurement.origin})
//                    else:
//                        po.write({'origin': po.origin})
//                else:
//                    po.write({'origin': procurement.origin})
//                name = (self.group_id and (self.group_id.name + ":")
//                        or "") + (self.name != "/" and self.name or "")
//                message = _(
//                    "This purchase order has been modified from: <a href=# data-oe-model=procurement.order data-oe-id=%d>%s</a>") % (procurement.id, name)
//                po.message_post(body=message)
//            if po:
//                res += [procurement.id]
//
//            # Create Line
//            po_line = False
//            for line in po.order_line:
//                if line.product_id == procurement.product_id and line.product_uom == procurement.product_id.uom_po_id:
//                    procurement_uom_po_qty = procurement.product_uom._compute_quantity(
//                        procurement.product_qty, procurement.product_id.uom_po_id)
//                    seller = procurement.product_id._select_seller(
//                        partner_id=partner,
//                        quantity=line.product_qty + procurement_uom_po_qty,
//                        date=po.date_order and po.date_order[:10],
//                        uom_id=procurement.product_id.uom_po_id)
//
//                    price_unit = self.env['account.tax']._fix_tax_included_price_company(
//                        seller.price, line.product_id.supplier_taxes_id, line.taxes_id, self.company_id) if seller else 0.0
//                    if price_unit and seller and po.currency_id and seller.currency_id != po.currency_id:
//                        price_unit = seller.currency_id.compute(
//                            price_unit, po.currency_id)
//
//                    po_line = line.write({
//                        'product_qty': line.product_qty + procurement_uom_po_qty,
//                        'price_unit': price_unit,
//                        'procurement_ids': [(4, procurement.id)]
//                    })
//                    break
//            if not po_line:
//                vals = procurement._prepare_purchase_order_line(po, supplier)
//                self.env['purchase.order.line'].create(vals)
//        return res
})
h.ProcurementOrder().Methods().OpenPurchaseOrder().DeclareMethod(
`OpenPurchaseOrder`,
func(rs m.ProcurementOrderSet)  {
//        action = self.env.ref('purchase.purchase_order_action_generic')
//        action_dict = action.read()[0]
//        action_dict['res_id'] = self.purchase_id.id
//        action_dict['target'] = 'current'
//        return action_dict
})
h.ProductTemplate().DeclareModel()


h.ProductTemplate().Methods().GetBuyRoute().DeclareMethod(
`GetBuyRoute`,
func(rs m.ProductTemplateSet)  {
//        buy_route = self.env.ref(
//            'purchase.route_warehouse0_buy', raise_if_not_found=False)
//        if buy_route:
//            return buy_route.ids
//        return []
})
h.ProductTemplate().Methods().PurchaseCount().DeclareMethod(
`PurchaseCount`,
func(rs m.ProductTemplateSet)  {
//        for template in self:
//            template.purchase_count = sum(
//                [p.purchase_count for p in template.product_variant_ids])
//        return True
})
h.ProductTemplate().AddFields(map[string]models.FieldDefinition{
"PropertyAccountCreditorPriceDifference": models.Many2OneField{
RelationModel: h.AccountAccount(),
String: "Price Difference Account",
//company_dependent=True
Help: "This account will be used to value price difference between" + 
"purchase price and cost price.",
},
"PurchaseCount": models.IntegerField{
Compute: h.ProductTemplate().Methods().PurchaseCount(),
String: "# Purchases",
},
"PurchaseMethod": models.SelectionField{
Selection: types.Selection{
"purchase": "On ordered quantities",
"receive": "On received quantities",
},
String: "Control Purchase Bills",
Help: "On ordered quantities: control bills based on ordered quantities." + 
"On received quantities: control bills based on received quantity.",
Default: models.DefaultValue("receive"),
},
"RouteIds": models.Many2ManyField{
Default: func (env models.Environment) interface{} { return self._get_buy_route() },
},
"PurchaseLineWarn": models.SelectionField{
Selection: WARNING_MESSAGE,
String: "Purchase Order Line",
{"_fields":["id","ctx"],"ctx":null,"id":"WARNING_HELP","loc":{"end":{"column":65,"line":1286},"start":{"column":53,"line":1286}},"type":"Name"}
Required: true,
Default: models.DefaultValue("no-message"),
},
"PurchaseLineWarnMsg": models.TextField{
String: "Message for Purchase Order Line",
},
})
h.ProductProduct().DeclareModel()


h.ProductProduct().Methods().PurchaseCount().DeclareMethod(
`PurchaseCount`,
func(rs h.ProductProductSet) h.ProductProductData {
//        domain = [
//            ('state', 'in', ['purchase', 'done']),
//            ('product_id', 'in', self.mapped('id')),
//        ]
//        PurchaseOrderLines = self.env['purchase.order.line'].search(domain)
//        for product in self:
//            product.purchase_count = len(PurchaseOrderLines.filtered(
//                lambda r: r.product_id == product).mapped('order_id'))
})
h.ProductProduct().AddFields(map[string]models.FieldDefinition{
"PurchaseCount": models.IntegerField{
Compute: h.ProductProduct().Methods().PurchaseCount(),
String: "# Purchases",
},
})
h.ProductCategory().DeclareModel()

h.ProductCategory().AddFields(map[string]models.FieldDefinition{
"PropertyAccountCreditorPriceDifferenceCateg": models.Many2OneField{
RelationModel: h.AccountAccount(),
String: "Price Difference Account",
//company_dependent=True
Help: "This account will be used to value price difference between" + 
"purchase price and accounting cost.",
},
})
h.MailComposeMessage().DeclareModel()

h.MailComposeMessage().Methods().MailPurchaseOrderOnSend().DeclareMethod(
`MailPurchaseOrderOnSend`,
func(rs m.MailComposeMessageSet)  {
//        if self._context.get('purchase_mark_rfq_sent'):
//            order = self.env['purchase.order'].browse(
//                self._context['default_res_id'])
//            if order.state == 'draft':
//                order.state = 'sent'
})
h.MailComposeMessage().Methods().SendMail().DeclareMethod(
`SendMail`,
func(rs m.MailComposeMessageSet, auto_commit interface{})  {
//        if self._context.get('default_model') == 'purchase.order' and self._context.get('default_res_id'):
//            self = self.with_context(mail_post_autofollow=True)
//            self.mail_purchase_order_on_send()
//        return super(MailComposeMessage, self).send_mail(auto_commit=auto_commit)
})
}