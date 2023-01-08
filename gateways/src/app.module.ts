import { Module } from '@nestjs/common'
import { ConfigModule } from '@nestjs/config'

import { AppController } from 'app.controller'
import { AppService } from 'app.service'
import { OrdersModule } from '@services/orders/order.module'
import { ProductsModule } from '@services/products/product.module'

@Module({
	imports: [
		OrdersModule,
		ProductsModule,
		ConfigModule.forRoot({
			envFilePath: ['.env'],
			isGlobal: true
		})
	],
	controllers: [AppController],
	providers: [AppService]
})
export class AppModule {}
