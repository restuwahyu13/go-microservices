import { Controller, Get, Version } from '@nestjs/common'
import { AppService } from 'app.service'

@Controller()
export class AppController {
	constructor(private readonly appService: AppService) {}

	@Version('1')
	@Get()
	PingOrders(): Record<string, any> {
		return this.appService.PingGateways()
	}
}
