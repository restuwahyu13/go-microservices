import { Injectable, OnModuleInit, Inject } from '@nestjs/common'
import { ClientGrpc, GrpcMethod } from '@nestjs/microservices'

import { GrpcResponse, ProductsService as ProductsRpcService, ServiceName, serviceNameToJSON } from '@schemas/proto-files/js/products'
import { Empty } from '@schemas/google/protobuf/empty'

@Injectable()
export class ProductsService implements OnModuleInit {
	private productsService: ProductsRpcService

	constructor(@Inject(ProductsService.name) private client: ClientGrpc) {}

	onModuleInit() {
		this.productsService = this.client.getService<ProductsRpcService>(ProductsService.name)
	}

	@GrpcMethod(ProductsService.name, serviceNameToJSON(ServiceName.PingProducts))
	async PingProducts(): Promise<GrpcResponse> {
		return await this.productsService.PingProducts(Empty)
	}
}
