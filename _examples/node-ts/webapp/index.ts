import {Example, VersionFromHeader, WebrpcError, WebrpcHeader} from './client.gen'

const fetchWithWebrpcHeaderParsing = async (input: RequestInfo | URL, init?: RequestInit): Promise<Response> => {
	if (init?.headers) {
		console.log('client headers', VersionFromHeader(new Headers(init.headers)))
	}
	const res = await fetch(input, init)
	console.log('server headers', res.headers.get(WebrpcHeader))
	return res
}

const example = new Example('http://localhost:3000', fetchWithWebrpcHeaderParsing)

async function onPingClick(pingText: HTMLElement) {
	try {
		await example.ping({})
		pingText.textContent = 'PONG'
	} catch (error) {
		if (error instanceof WebrpcError) {
			console.error(error)
			pingText.textContent = `error: ${error.message}, cause: ${error.cause}`
		}
	}
}

async function onGetUserClick(usernameText: HTMLElement) {
	try {
		const { user } = await example.getUser({ userId: 1 })
		console.log('getUser() responded with:', { user })

		if (typeof user.balance !== 'bigint') {
			throw new Error('user.balance is not a bigint')
		}
		console.log('user.balance as bigint:', user.balance)
		if (user.balance !== BigInt(31337)) {
			throw new Error(`user.balance has unexpected value: ${user.balance.toString()}`)
		}
		if (typeof user.extra?.amount !== 'bigint') {
			throw new Error('user.extra.amount is not a bigint')
		}
		if (user.extra.amount !== BigInt(5678)) {
			throw new Error(`user.extra.amount has unexpected value: ${user.extra.amount.toString()}`)
		}
		console.log('user.extra.amount as bigint:', user.extra.amount)

		usernameText.textContent = user.USERNAME
	} catch (error) {
		if (error instanceof WebrpcError) {
			console.error(error)
			usernameText.textContent = `webrpc error: ${error.message}, cause: ${error.cause}`
		} else {
			console.error('unexpected error:', error)
		}
	}
}

async function onGetUserWithErrorClick(usernameWithErrorText: HTMLElement) {
	try {
		const { user } = await example.getUser({ userId: 911 })
		console.log('getUser() responded with:', { user })
		usernameWithErrorText.textContent = user.USERNAME
	} catch (error) {
		if (error instanceof WebrpcError) {
			console.error(error)
			usernameWithErrorText.textContent = `webrpc error: ${error.message}, cause: ${error.cause}`
		} else {
			console.error('unexpected error:', error)
		}
	}
}

async function onGetArticleClick(articleText: HTMLElement) {
	try {
		const article = await example.getArticle({ articleId: 1, byBN: BigInt(444555) })
		console.log('getArticle() responded with:', { article })

		if (typeof article.largeNum !== 'bigint') {
			throw new Error('article.largeNum is not a bigint')
		}
		console.log('article.largeNum as bigint:', article.largeNum)
		if (article.largeNum !== BigInt(444555) * BigInt(2)) {
			throw new Error(`article.largeNum has unexpected value: ${article.largeNum.toString()}`)
		}

		articleText.textContent = `Title: ${article.title}\n\nContent: ${article.content}`
	} catch (error) {
		if (error instanceof WebrpcError) {
			console.error(error)
			articleText.textContent = `webrpc error: ${error.message}, cause: ${error.cause}`
		} else {
			console.error('unexpected error:', error)
		}
	}
}

document.addEventListener('DOMContentLoaded', () => {
	const pingButton = document.getElementById('js-ping-btn')
	const pingText = document.getElementById('js-ping-text')
	const getUserButton = document.getElementById('js-get-user-btn')
	const usernameText = document.getElementById('js-username-text')
	const getUserWithErrorButton = document.getElementById('js-get-user-error-btn')
	const usernameWithErrorText = document.getElementById('js-username-error-text')
	const getArticleButton = document.getElementById('js-get-article-btn')
	const articleText = document.getElementById('js-article-text')

	if (!pingButton || !pingText) {
		console.log('error getting ping HTML elements')
	} else {
		pingButton.addEventListener('click', () => onPingClick(pingText))
	}

	if (!getUserButton || !usernameText) {
		console.log('error getting username HTML elements')
	} else {
		getUserButton.addEventListener('click', () => onGetUserClick(usernameText))
	}

	if (!getUserWithErrorButton || !usernameWithErrorText) {
		console.log('error getting username HTML elements')
	} else {
		getUserWithErrorButton.addEventListener('click', () => onGetUserWithErrorClick(usernameWithErrorText))
	}

	if (!getArticleButton || !articleText) {
		console.log('error getting article HTML elements')
	} else {
		getArticleButton.addEventListener('click', () => { void onGetArticleClick(articleText) })
	}
})
