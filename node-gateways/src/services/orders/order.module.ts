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
					url: 'localhost:4001',
					protoPath: path.join(process.cwd(), '../proto-files/orders.proto')
				}
			}
		])
	],
	providers: [OrdersService]
})
export class OrdersModule {}
