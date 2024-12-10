"use client"

import { useState, useEffect } from "react"
import { SubmissionInfo } from "@/app/ts/Submission"
import { AssetImage } from "@/app/ts/Roblox"
import Image from "next/image"

interface AssetID {
    id: SubmissionInfo["AssetID"]
}
function MapImage(asset: AssetID) {
	const [assetImage, setAssetImage] = useState("");

	useEffect(() => {
		AssetImage(asset.id, "420x420").then(image => setAssetImage(image))
	}, []);
	if (!assetImage) {
        return <p>Fetching map image...</p>;
    }
    return <Image src={assetImage} alt="Map Image"/>
}

export {
	type AssetID,
	MapImage
}