import { Module } from '@nestjs/common'
import { ClientsModule, Transport } from '@nestjs/microservices'
import path from 'path'

import { OrdersService } from '@services/orders/order.service'
import { OrdersController } from '@services/orders/order.controller'
import { protobufPackage } from '@schemas/proto-files/orders'

@Module({
	controllers: [OrdersController],
	imports: [
		ClientsModule.register([
			{
				name: OrdersService.name,
				transport: Transport.GRPC,
				options: {
					package: protobufPackage,
					url: process.env.MS_ORDERS_HOST,
					protoPath: path.join(process.cwd(), '../protofiles/orders.proto')
				}
			}
		])
	],
	providers: [OrdersService]
})
export class OrdersModule {}
