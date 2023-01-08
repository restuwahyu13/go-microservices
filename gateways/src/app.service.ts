import { Injectable, HttpStatus as status } from '@nestjs/common'

interface ApiResponse {
	stat_code: number
	stat_message: string
}

@Injectable()
export class AppService {
	PingGateways(): ApiResponse {
		return {
			stat_code: status.OK,
			stat_message: 'Gateway Pong'
		}
	}
}
