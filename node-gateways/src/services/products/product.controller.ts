import { Controller, Get, Version } from '@nestjs/common'
import { ProductsService } from '@services/products/product.service'
import { GrpcResponse } from '@schemas/proto-files/js/products'

@Controller('product')
export class ProductsController {
	constructor(private readonly productsService: ProductsService) {}

	@Version('1')
	@Get()
	PingProducts(): Promise<GrpcResponse> {
		return this.productsService.PingProducts()
	}
}
