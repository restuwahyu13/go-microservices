import { Injectable, OnModuleInit, Inject } from '@nestjs/common'
import { ClientGrpc, GrpcMethod } from '@nestjs/microservices'

import { GrpcResponse, OrdersService as OrdersRpcService, ServiceName, serviceNameToJSON } from '@schemas/proto-files/orders'
import { Empty } from '@schemas/google/protobuf/empty'

@Injectable()
export class OrdersService implements OnModuleInit {
	private ordersService: OrdersRpcService

	constructor(@Inject(OrdersService.name) private client: ClientGrpc) {}

	onModuleInit() {
		this.ordersService = this.client.getService<OrdersRpcService>(OrdersService.name)
	}

	@GrpcMethod(OrdersService.name, serviceNameToJSON(ServiceName.PingOrders))
	PingOrders(): Promise<GrpcResponse> {
		return this.ordersService.PingOrders(Empty)
	}
}
