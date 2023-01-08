require('dotenv/config')
const path = require('path')

const migrationsDir = path.join(__dirname, 'migrations')
const seedsDir = path.join(__dirname, 'seeds')

module.exports = [
	{
		default: 'users',
		type: 'postgres',
		host: process.env.PG_HOST,
		port: parseInt(process.env.PG_PORT),
		username: process.env.PG_USER,
		password: process.env.PG_PASSWORD,
		database: process.env.PG_USERS_DB_NAME,
		migrations: [migrationsDir],
		cli: {
			migrationsDir: migrationsDir,
			seedsDir: seedsDir
		}
	},
	{
		default: 'orders',
		type: 'postgres',
		host: process.env.PG_HOST,
		port: parseInt(process.env.PG_PORT),
		username: process.env.PG_USER,
		password: process.env.PG_PASSWORD,
		database: process.env.PG_ORDERS_DB_NAME,
		migrations: [migrationsDir],
		cli: {
			migrationsDir: migrationsDir,
			seedsDir: seedsDir
		}
	},
	{
		default: 'products',
		type: 'postgres',
		host: process.env.PG_HOST,
		port: parseInt(process.env.PG_PORT),
		username: process.env.PG_USER,
		password: process.env.PG_PASSWORD,
		database: process.env.PG_PRODUCTS_DB_NAME,
		migrations: [migrationsDir],
		cli: {
			migrationsDir: migrationsDir,
			seedsDir: seedsDir
		}
	},
	{
		default: 'shippings',
		type: 'postgres',
		host: process.env.PG_HOST,
		port: parseInt(process.env.PG_PORT),
		username: process.env.PG_USER,
		password: process.env.PG_PASSWORD,
		database: process.env.PG_SHIPPINGS_DB_NAME,
		migrations: [migrationsDir],
		cli: {
			migrationsDir: migrationsDir,
			seedsDir: seedsDir
		}
	}
]
