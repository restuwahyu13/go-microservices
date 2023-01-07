import { Controller, Get, Version } from '@nestjs/common'
import { OrdersService } from '@services/orders/order.service'
import { GrpcResponse } from '@schemas/proto-files/js/orders'

@Controller('order')
export class OrdersController {
	constructor(private readonly ordersService: OrdersService) {}

	@Version('1')
	@Get()
	PingOrders(): Promise<GrpcResponse> {
		return this.ordersService.PingOrders()
	}
}
