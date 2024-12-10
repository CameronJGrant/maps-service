"use client"

import { SubmissionStatus, SubmissionStatusToString, type SubmissionInfo } from "@/app/ts/Submission";
import { MapImage, type AssetID } from "./map";
import { Rating, Button } from "@mui/material";
import type {ReactNode} from "react";
import SendIcon from '@mui/icons-material/Send';
import Link from "next/link";

import "./styles/page.scss";

interface Window {
    className: string,
    title: string,
    children: ReactNode
}
interface Comment {
    picture?: string, //TEMP
    comment: string,
    date: string,
    name: string
}
interface CreatorAndReviewStatus {
    creator: SubmissionInfo["DisplayName"],
    review: SubmissionInfo["StatusID"],
    comments: Comment[]
}
interface MapInfo {
    assetid: SubmissionInfo["AssetID"],
    status: SubmissionStatus,
    comments: Comment[]
}

function Window(window: Window) {
    return (
        <section className={window.className}>
            <header>
                <p>{window.title}</p>
            </header>
            <main>{window.children}</main>
        </section>
    )
}

function ImageAndRatings(asset: AssetID) {
    return (
        <aside className="review-area">
        	<section className="map-image-area">
         		<MapImage id={asset.id}/>
         	</section>
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
        </aside>
    )
}

function Comment(comment: Comment) {
    const IsBhopMaptest = comment.name == "BhopMaptest" //Highlighted commenter

    return (
        <div className="commenter" data-highlighted={IsBhopMaptest}>
            <img src={comment.picture} alt={`${comment.name}'s comment`}/>
            <div className="details">
                <header>
                    <p className="name">{comment.name}</p>
                    <p className="date">{comment.date}</p>
                </header>
                <p className="comment">{comment.comment}</p>
            </div>
        </div>
    );
}

function TitleAndComments(stats: CreatorAndReviewStatus) {
    const Review = SubmissionStatusToString(stats.review)

	return (
        <main className="review-info">
            <div>
                <h1>bhop_quaternions</h1>
                <aside data-review-status={stats.review} className="review-status">
                    <p>{Review}</p>
                </aside>
            </div>
            <p className="by-creator">by <Link href="" target="_blank">{stats.creator}</Link></p>
            <span className="spacer"></span>
            <section className="comments">
            	{stats.comments.length===0 && <p className="no-comments">There are no comments.</p> || stats.comments.map(comment => (
                	<Comment key={comment.name} name={comment.name} date={comment.date} comment={comment.comment}/>
             	))}
            </section>
            <Window title="Leave a Comment:" className="leave-comment-window">
                <textarea name="comment-box" id="comment-text-field"></textarea>
                <Button variant="contained" endIcon={<SendIcon/>}>Submit</Button>
            </Window>
        </main>
    )
}

export default function MapInfoPage(info: MapInfo) {
	return (
        <section className="review-section">
            <ImageAndRatings id={info.assetid}/>
            <TitleAndComments creator="Quaternions" review={info.status} comments={info.comments}/>
        </section>
	)
}