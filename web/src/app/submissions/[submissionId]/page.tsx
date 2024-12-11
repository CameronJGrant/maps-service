"use client"

import { SubmissionStatus, SubmissionStatusToString } from "@/app/ts/Submission";
import type { CreatorAndReviewStatus } from "./_comments";
import { MapImage, type AssetID } from "./_map";
import { useParams } from "next/navigation";
import ReviewButtons from "./_reviewButtons";
import { Rating } from "@mui/material";
import Comments from "./_comments";
import Webpage from "@/app/_components/webpage";
import Window from "./_window";
import Link from "next/link";

import "./(styles)/page.scss";

function Ratings() {
	return (
		<Window className="rating-window" title="Rating">
		    <section className="rating-type">
		        <aside className="rating-left">
		            <p>Quality</p>
		            <p>Difficulty</p>
		            <p>Fun</p>
		            <p>Length</p>
		        </aside>
		        <aside className="rating-right">
		            <Rating defaultValue={2.5} precision={0.5}/>
		            <Rating defaultValue={2.5} precision={0.5}/>
		            <Rating defaultValue={2.5} precision={0.5}/>
		            <Rating defaultValue={2.5} precision={0.5}/>
		        </aside>
		    </section>
        </Window>
	)
}

function RatingArea(asset: AssetID) {
    return (
        <aside className="review-area">
        	<section className="map-image-area">
         		<MapImage id={asset.id}/>
         	</section>
            <Ratings/>
            <ReviewButtons/>
        </aside>
    )
}

function TitleAndComments(stats: CreatorAndReviewStatus) {
    const Review = SubmissionStatusToString(stats.review)

	return (
        <main className="review-info">
            <div>
                <h1>{stats.name}</h1>
                <aside data-review-status={stats.review} className="review-status">
                    <p>{Review}</p>
                </aside>
            </div>
            <p className="by-creator">by <Link href="" target="_blank">{stats.creator}</Link></p>
            <span className="spacer"></span>
            <Comments comments_data={stats}/>
        </main>
    )
}

export default function SubmissionInfoPage() {
    const dynamicId = useParams<{submissionId: string}>()

    return (
        <Webpage>
            <main className="map-page-main">
                <section className="review-section">
                    <RatingArea id={432}/>
                    <TitleAndComments name={dynamicId.submissionId} creator="Quaternions" review={SubmissionStatus.Accepted} comments={[]}/>
                </section>
            </main>
        </Webpage>
	)
}