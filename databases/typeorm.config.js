require('dotenv/config')
const path = require('path')

const migrationsDir = path.resolve(process.cwd(), 'database-migrations/migrations')
const seedsDir = path.resolve(process.cwd(), 'database-migrations/seeds')

module.exports = {
	type: 'postgres',
	host: process.env.DB_HOST,
	port: parseInt(process.env.DB_PORT),
	username: process.env.DB_USER,
	password: process.env.DB_PASSWORD,
	database: process.env.DB_NAME,
	entities: [entitiesDir],
	migrations: [migrationsDir],
	cli: {
		migrationsDir: migrationsDir,
		seedsDir: seedsDir
	}
}
