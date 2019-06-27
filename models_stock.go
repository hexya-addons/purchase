package purchase

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.StockPicking().DeclareModel()

	h.StockPicking().AddFields(map[string]models.FieldDefinition{
		"PurchaseId": models.Many2OneField{
			RelationModel: h.PurchaseOrder(),
			Related:       `MoveLines.PurchaseLineId.OrderId`,
			String:        "Purchase Orders",
			ReadOnly:      true,
		},
	})
	h.StockPicking().Methods().PrepareValuesExtraMove().DeclareMethod(
		`PrepareValuesExtraMove`,
		func(rs m.StockPickingSet, op interface{}, product interface{}, remaining_qty interface{}) {
			//        res = super(StockPicking, self)._prepare_values_extra_move(
			//            op, product, remaining_qty)
			//        for m in op.linked_move_operation_ids:
			//            if m.move_id.purchase_line_id and m.move_id.product_id == product:
			//                res['purchase_line_id'] = m.move_id.purchase_line_id.id
			//                break
			//        return res
		})
	h.StockPicking().Methods().CreateBackorder().DeclareMethod(
		`CreateBackorder`,
		func(rs m.StockPickingSet, backorder_moves interface{}) {
			//        res = super(StockPicking, self)._create_backorder(backorder_moves)
			//        for picking in self:
			//            if picking.picking_type_id.code == 'incoming':
			//                for backorder in self.search([('backorder_id', '=', picking.id)]):
			//                    backorder.message_post_with_view('mail.message_origin_link',
			//                                                     values={
			//                                                         'self': backorder, 'origin': backorder.purchase_id},
			//                                                     subtype_id=self.env.ref('mail.mt_note').id)
			//        return res
		})
	h.StockMove().DeclareModel()

	h.StockMove().AddFields(map[string]models.FieldDefinition{
		"PurchaseLineId": models.Many2OneField{
			RelationModel: h.PurchaseOrderLine(),
			String:        "Purchase Order Line",
			OnDelete:      `set null`,
			Index:         true,
			ReadOnly:      true,
		},
	})
	h.StockMove().Methods().GetPriceUnit().DeclareMethod(
		` Returns the unit price to store on the quant `,
		func(rs m.StockMoveSet) {
			//        if self.purchase_line_id:
			//            order = self.purchase_line_id.order_id
			//            # if the currency of the PO is different than the company one, the price_unit on the move must be reevaluated
			//            #(was created at the rate of the PO confirmation, but must be valuated at the rate of stock move execution)
			//            if order.currency_id != self.company_id.currency_id:
			//                # we don't pass the move.date in the compute() for the currency rate on purpose because
			//                # 1) get_price_unit() is supposed to be called only through move.action_done(),
			//                # 2) the move hasn't yet the correct date (currently it is the expected date, after
			//                #    completion of action_done() it will be now() )
			//                price_unit = self.purchase_line_id._get_stock_move_price_unit()
			//                self.write({'price_unit': price_unit})
			//                return price_unit
			//            return self.price_unit
			//        return super(StockMove, self).get_price_unit()
		})
	h.StockMove().Methods().Copy().Extend(
		`Copy`,
		func(rs m.StockMoveSet, defaultName models.RecordData) {
			//        self.ensure_one()
			//        default = default or {}
			//        if not default.get('split_from'):
			//            # we don't want to propagate the link to the purchase order line except in case of move split
			//            default['purchase_line_id'] = False
			//        return super(StockMove, self).copy(default)
		})
	h.StockWarehouse().DeclareModel()

	h.StockWarehouse().AddFields(map[string]models.FieldDefinition{
		"BuyToResupply": models.BooleanField{
			String:  "Purchase to resupply this warehouse",
			Default: models.DefaultValue(true),
			Help:    "When products are bought, they can be delivered to this warehouse",
		},
		"BuyPullId": models.Many2OneField{
			RelationModel: h.ProcurementRule(),
			String:        "Buy rule",
		},
	})
	h.StockWarehouse().Methods().GetBuyPullRule().DeclareMethod(
		`GetBuyPullRule`,
		func(rs m.StockWarehouseSet) {
			//        try:
			//            buy_route_id = self.env['ir.model.data'].get_object_reference(
			//                'purchase', 'route_warehouse0_buy')[1]
			//        except:
			//            buy_route_id = self.env['stock.location.route'].search(
			//                [('name', 'like', _('Buy'))])
			//            buy_route_id = buy_route_id[0].id if buy_route_id else False
			//        if not buy_route_id:
			//            raise UserError(_("Can't find any generic Buy route."))
			//        return {
			//            'name': self._format_routename(_(' Buy')),
			//            'location_id': self.in_type_id.default_location_dest_id.id,
			//            'route_id': buy_route_id,
			//            'action': 'buy',
			//            'picking_type_id': self.in_type_id.id,
			//            'warehouse_id': self.id,
			//            'group_propagation_option': 'none',
			//        }
		})
	h.StockWarehouse().Methods().CreateRoutes().DeclareMethod(
		`CreateRoutes`,
		func(rs m.StockWarehouseSet) {
			//        res = super(StockWarehouse, self).create_routes()
			//        if self.buy_to_resupply:
			//            buy_pull_vals = self._get_buy_pull_rule()
			//            buy_pull = self.env['procurement.rule'].create(buy_pull_vals)
			//            res['buy_pull_id'] = buy_pull.id
			//        return res
		})
	h.StockWarehouse().Methods().Write().Extend(
		`Write`,
		func(rs m.StockWarehouseSet, vals models.RecordData) {
			//        if 'buy_to_resupply' in vals:
			//            if vals.get("buy_to_resupply"):
			//                for warehouse in self:
			//                    if not warehouse.buy_pull_id:
			//                        buy_pull_vals = self._get_buy_pull_rule()
			//                        buy_pull = self.env['procurement.rule'].create(
			//                            buy_pull_vals)
			//                        vals['buy_pull_id'] = buy_pull.id
			//            else:
			//                for warehouse in self:
			//                    if warehouse.buy_pull_id:
			//                        warehouse.buy_pull_id.unlink()
			//        return super(StockWarehouse, self).write(vals)
		})
	h.StockWarehouse().Methods().GetAllRoutes().DeclareMethod(
		`GetAllRoutes`,
		func(rs m.StockWarehouseSet) {
			//        routes = super(StockWarehouse, self).get_all_routes_for_wh()
			//        routes |= self.filtered(lambda self: self.buy_to_resupply and self.buy_pull_id and self.buy_pull_id.route_id).mapped(
			//            'buy_pull_id').mapped('route_id')
			//        return routes
		})
	h.StockWarehouse().Methods().UpdateNameAndCode().DeclareMethod(
		`UpdateNameAndCode`,
		func(rs m.StockWarehouseSet, name interface{}, code interface{}) {
			//        res = super(StockWarehouse, self)._update_name_and_code(name, code)
			//        warehouse = self[0]
			//        if warehouse.buy_pull_id and name:
			//            warehouse.buy_pull_id.write(
			//                {'name': warehouse.buy_pull_id.name.replace(warehouse.name, name, 1)})
			//        return res
		})
	h.StockWarehouse().Methods().UpdateRoutes().DeclareMethod(
		`UpdateRoutes`,
		func(rs m.StockWarehouseSet) {
			//        res = super(StockWarehouse, self)._update_routes()
			//        for warehouse in self:
			//            if warehouse.in_type_id.default_location_dest_id != warehouse.buy_pull_id.location_id:
			//                warehouse.buy_pull_id.write(
			//                    {'location_id': warehouse.in_type_id.default_location_dest_id.id})
			//        return res
		})
}
