import { SubmissionStatus } from "@/app/ts/Submission";
import MapInfoPage from "./page";
import Header from "@/app/_components/header";

import "./styles/layout.scss";

export default function MapPage() {
    const placeholder_Comment = [
        {
            comment: "This map has been accepted and is in the game.",
            date: "on Dec 8 '24 at 18:46",
            name: "BhopMaptest"
        },
        {
            comment: "This map is so mid...",
            date: "on Dec 8 '24 at 18:46",
            name: "vmsize"
        },
        {
            comment: "I prefer strafe client",
            date: "on Dec 8 '24 at 18:46",
            name: "Quaternions"
        }
    ]

	return (
		<>
			<Header/>
            <main className="map-page-main">
                <MapInfoPage assetid={14783300138} status={SubmissionStatus.Accepted} comments={placeholder_Comment}/>
            </main>
        </>
    )
}