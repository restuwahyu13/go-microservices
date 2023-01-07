module.exports = {
	apps: [
		{
			name: 'microservices-gateways',
			namespace: 'gateways',
			script: 'dist/main.js',
			watch: false,
			autorestart: true,
			env: {
				NODE_ENV: 'production',
				NODE_OPTIONS: '--max_old_space_size=16384',
				PORT: 3000
			},
			exec_mode: 'cluster',
			instances: 'max',
			max_memory_restart: '512M',
			listen_timeout: 3000,
			kill_timeout: 6000,
			combine_logs: true
		}
	]
}
