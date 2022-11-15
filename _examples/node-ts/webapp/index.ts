import { ExampleService } from './client.gen'

const exampleService = new ExampleService(
	'http://localhost:3000',
	(input, init) => fetch(input, init)
)

document.addEventListener('DOMContentLoaded', () => {
	const pingButton = document.getElementById('js-ping-btn')
	const pingText = document.getElementById('js-ping-text')

	if (!pingButton || !pingText) {
		console.log('error getting ping HTML elements')
		return
	}

	pingButton.addEventListener('click', () => {
		exampleService
			.ping({})
			.then(({}) => {
				console.log('ping() responded:', {})
				pingText.textContent = 'PONG'
			})
			.catch((e) => {
				console.log('ping() failed:', e)
				pingText.textContent = 'ping() failed: ' + e.message
			})
	})
})

document.addEventListener('DOMContentLoaded', () => {
	const getUserButton = document.getElementById('js-get-user-btn')
	const usernameText = document.getElementById('js-username-text')

	if (!getUserButton || !usernameText) {
		console.log('error getting username HTML elements')
		return
	}

	getUserButton.addEventListener('click', () => {
		exampleService
			.getUser({
				userID: 1,
			})
			.then(({ user }) => {
				console.log('getUser() responded with:', { user })
				usernameText.textContent = user.USERNAME
			})
			.catch((e) => {
				console.log('getUser() failed:', e)
				usernameText.textContent = 'getUser() failed: ' + e.message
			})
	})
})
