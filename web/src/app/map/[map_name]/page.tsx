import { SubmissionStatus, type SubmissionInfo } from "@/app/ts/Submission";
import { useState, type ReactNode } from "react";
import { Rating, Button } from "@mui/material";
import { AssetImage } from "@/app/ts/Roblox";
import SendIcon from '@mui/icons-material/Send';

import "./styles/page.scss";

interface Window {
    className: string,
    title: string,
    children: ReactNode
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

interface AssetID {
    id: SubmissionInfo["AssetID"]
}
async function ImageAndRatings(asset: AssetID) {
	const [assetImage, setAssetImage] = useState(null);
    const asset_image = await AssetImage(asset.id, "420x420")

    return (
        <aside className="review-area">
        	<section className="rating">
         		<img src={asset_image} alt="Map Image"/>
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

interface Comment {
    picture?: string, //TEMP
    comment: string,
    date: string,
    name: string
}
function Comment(comment: Comment) {
    let placeHolder;
    if (comment.name === "BhopMaptest") {
        placeHolder = "https://tr.rbxcdn.com/30DAY-AvatarHeadshot-FB29ADF0A483B2745DB2571DC4785202-Png/150/150/AvatarHeadshot/Webp/noFilter"
    } else if (comment.name === "vmsize") {
        placeHolder = "https://tr.rbxcdn.com/30DAY-AvatarHeadshot-ACEB71FADC70B458ECB9D6AA9AAE5913-Png/150/150/AvatarHeadshot/Webp/noFilter"
    } else {
        placeHolder = "https://tr.rbxcdn.com/30DAY-AvatarHeadshot-1ED6D3ED61793733397BB596F0ADD369-Png/150/150/AvatarHeadshot/Webp/noFilter"
    }

    //Highlighted comment
    const IsBhopMaptest = comment.name == "BhopMaptest"

    return (
        <div className="commenter" data-highlighted={IsBhopMaptest}>
            <img src={placeHolder} alt={`${comment.name}'s comment`}/>
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

interface CreatorAndReviewStatus {
    creator: SubmissionInfo["DisplayName"],
    review: SubmissionInfo["StatusID"],
    comments: Comment[]
}
function TitleAndComments(stats: CreatorAndReviewStatus) {
    //TODO: switch case this for matching the enums
    return (
        <main className="review-info">
            <div>
                <h1>bhop_quaternions</h1>
                <aside data-review-status={stats.review} className="review-status">
                    <p>ACCEPTED</p>
                </aside>
            </div>
            <p className="by-creator">by <a href="" target="_blank">{stats.creator}</a></p>
            <span className="spacer"></span>
            <main className="comments">
                {stats.comments.map(comment => (
                    <Comment name={comment.name} date={comment.date} comment={comment.comment}/>
                ))}
            </main>
            <Window title="Leave a Comment:" className="leave-comment-window">
                <textarea name="comment-box" id="comment-text-field"></textarea>
                <Button variant="contained" endIcon={<SendIcon/>}>Submit</Button>
            </Window>
        </main>
    )
}

interface MapInfo {
    assetid: SubmissionInfo["AssetID"],
    status: SubmissionStatus,
    comments: Comment[]
}
export default function MapInfoPage(info: MapInfo) {
	return (
        <section className="review-section">
            <ImageAndRatings id={info.assetid}/>
            <TitleAndComments creator="Quaternions" review={info.status} comments={info.comments}/>
        </section>
	)
}