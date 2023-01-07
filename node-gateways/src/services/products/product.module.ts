import { Module } from '@nestjs/common'
import { ClientsModule, Transport } from '@nestjs/microservices'
import path from 'path'

import { ProductsService } from '@services/products/product.service'
import { ProductsController } from '@services/products/product.controller'
import { protobufPackage } from '@schemas/proto-files/js/products'

@Module({
	controllers: [ProductsController],
	imports: [
		ClientsModule.register([
			{
				name: ProductsService.name,
				transport: Transport.GRPC,
				options: {
					package: protobufPackage,
					url: 'localhost:4002',
					protoPath: path.join(process.cwd(), '../proto-files/js/products.proto')
				}
			}
		])
	],
	providers: [ProductsService]
})
export class ProductsModule {}
