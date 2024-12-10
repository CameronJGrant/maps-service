const FALLBACK_IMAGE = ""

type thumbsizes = "420" | "720"
type thumbsize<S extends thumbsizes> = `${S}x${S}`

async function RoproxyParse(api: Response): Promise<string> {
	if (api.ok) {
		const json = await api.json()
		if (json.errors) {
			console.warn(json.errors)
			return FALLBACK_IMAGE
		}
		if (json.data) {
			const data = json.data[0]
			if (!data) { //For whatever reason roblox will sometimes return an empty array instead of an error message
				console.warn("Roblox gave us no data,", json)
				return FALLBACK_IMAGE
			}
			if (data.state === "Completed") {
				return data.imageUrl
			}
			console.warn(data)
			return FALLBACK_IMAGE
		}
	}

	console.warn(api)
	return FALLBACK_IMAGE
}

export async function AvatarHeadshot<S extends thumbsizes>(userid: number, size: thumbsize<S>): Promise<string> {
	const avatarthumb_api = await fetch(`https://thumbnails.roproxy.com/v1/users/avatar-headshot?userIds=${userid}&size=${size}&format=Png&isCircular=false`)
	return RoproxyParse(avatarthumb_api)
}

export async function AssetImage<S extends thumbsizes>(assetid: number, size: thumbsize<S>): Promise<string> {
	const avatarthumb_api = await fetch(`https://thumbnails.roblox.com/v1/assets?assetIds=${assetid}&returnPolicy=PlaceHolder&size=${size}&format=Png&isCircular=false`)
	return RoproxyParse(avatarthumb_api)
}