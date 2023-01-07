import { HttpStatus as status, Injectable } from '@nestjs/common'

@Injectable()
export class AppService {
	getHello(): Record<string, any> {
		return {
			stat_code: status.OK,
			stat_message: 'Microservices Gateway'
		}
	}
}
