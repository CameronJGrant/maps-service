const FALLBACK_IMAGE = ""

type thumbsizes = "420" | "720"
type thumbsize<S extends thumbsizes> = `${S}x${S}`
type ParsedJson<A> = {
	errors: A,
	data: {
		[0]: {
			state: string,
			imageUrl: string,
		}
	}
}

function Parse<A>(json: ParsedJson<A>): string {
	if (json.errors) {
		console.warn(json.errors)
		return FALLBACK_IMAGE
	}
	if (json.data) {
		const data = json.data[0]
		if (!data) { //For whatever reason roblox will sometimes return an empty array instead of an error message
			console.warn("Roblox gave us no data,", data)
			return FALLBACK_IMAGE
		}
		if (data.state === "Completed") {
			return data.imageUrl
		}
		console.warn(data)
		return FALLBACK_IMAGE
	}
	return FALLBACK_IMAGE
}

async function AvatarHeadshot<S extends thumbsizes>(userid: number, size: thumbsize<S>): Promise<string> {
	const avatarthumb_api = fetch(`https://thumbnails.roproxy.com/v1/users/avatar-headshot?userIds=${userid}&size=${size}&format=Png&isCircular=false`)
	return avatarthumb_api.then(res => res.json()).then(json => Parse(json))
}

async function AssetImage<S extends thumbsizes>(assetid: number, size: thumbsize<S>): Promise<string> {
	const avatarthumb_api = fetch(`https://thumbnails.roblox.com/v1/assets?assetIds=${assetid}&returnPolicy=PlaceHolder&size=${size}&format=Png&isCircular=false`)
	return avatarthumb_api.then(res => res.json()).then(json => Parse(json))
}

export {
	AvatarHeadshot,
	AssetImage
}