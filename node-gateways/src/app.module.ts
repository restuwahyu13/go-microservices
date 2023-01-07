import { Module } from '@nestjs/common'

import { AppController } from 'app.controller'
import { AppService } from 'app.service'
import { OrdersModule } from '@services/orders/order.module'
import { ProductsModule } from '@services/products/product.module'

@Module({
	imports: [OrdersModule, ProductsModule],
	controllers: [AppController],
	providers: [AppService]
})
export class AppModule {}
