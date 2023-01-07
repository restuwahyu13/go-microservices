import 'express-async-errors'
import 'dotenv/config'
import 'reflect-metadata'
import reusify from 'reusify'
import { NestExpressApplication } from '@nestjs/platform-express'
import { Logger, VERSION_NEUTRAL, VersioningType } from '@nestjs/common'
import { NestFactory } from '@nestjs/core'
import compression from 'compression'
import helmet from 'helmet'
import hpp from 'hpp'
import nocache from 'nocache'
import zlib from 'zlib'
import bodyParser from 'body-parser'
import { AppModule } from 'app.module'

class Application {
	private nest: typeof NestFactory
	private app: NestExpressApplication
	private logger: Logger
	private env: string
	private port: number

	constructor() {
		this.nest = NestFactory
		this.env = process.env.NODE_ENV || 'development'
		this.port = +process.env.PORT || 3000
		this.logger = new Logger('HttpServer')
	}

	private async setupApplication(): Promise<void> {
		this.app = await this.nest.create<NestExpressApplication>(AppModule, {
			bodyParser: true,
			cors: true,
			bufferLogs: this.env == 'production' ? false : true,
			logger: this.env == 'production' ? false : ['debug', 'error', 'log', 'verbose', 'warn']
		})
	}

	private globalConfig(): void {
		this.app.setGlobalPrefix('api', { exclude: [] })
		this.app.enableVersioning({ type: VersioningType.URI, defaultVersion: VERSION_NEUTRAL })
		this.app.enableCors()
		this.app.disable('x-powered-by')
	}

	private globalMiddleware(): void {
		this.app.use(hpp())
		this.app.use(nocache())
		this.app.use(helmet())
		this.app.use(bodyParser.json({ limit: '3mb' }))
		this.app.use(bodyParser.raw({ limit: '3mb', type: 'application/json' }))
		this.app.use(bodyParser.urlencoded({ extended: true }))
		this.app.use(
			compression({
				level: zlib.constants.Z_BEST_COMPRESSION,
				memLevel: zlib.constants.Z_BEST_COMPRESSION,
				strategy: zlib.constants.Z_RLE
			})
		)
	}

	private serverListening(): void {
		this.app.listen(this.port, () => this.logger.log(`server is running on port ${this.port}`))
	}

	public async bootstraping(): Promise<void> {
		await this.setupApplication()
		this.globalConfig()
		this.globalMiddleware()
		this.serverListening()
	}
}

// boostraping nestjs application
;(() => {
	const app: InstanceType<typeof Application> = reusify(Application).get()
	app.bootstraping()
})()
