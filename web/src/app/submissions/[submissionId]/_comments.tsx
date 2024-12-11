import type { SubmissionInfo } from "@/app/ts/Submission";
import { Button } from "@mui/material"
import Window from "./_window";
import SendIcon from '@mui/icons-material/Send';
import Image from "next/image";

interface CommentersProps {
    comments_data: CreatorAndReviewStatus
}

interface CreatorAndReviewStatus {
    creator: SubmissionInfo["DisplayName"],
    review: SubmissionInfo["StatusID"],
    comments: Comment[],
    name: string
}

interface Comment {
    picture?: string, //TEMP
    comment: string,
    date: string,
    name: string
}

function AddComment(comment: Comment) {
    const IsBhopMaptest = comment.name == "BhopMaptest" //Highlighted commenter

    return (
        <div className="commenter" data-highlighted={IsBhopMaptest}>
            <Image src={comment.picture as string} alt={`${comment.name}'s comment`}/>
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

function LeaveAComment() {
	return (
		<Window title="Leave a Comment:" className="leave-comment-window">
            <textarea name="comment-box" id="comment-text-field"></textarea>
            <Button variant="outlined" endIcon={<SendIcon/>}>Submit</Button>
        </Window>
	)
}

export default function Comments(stats: CommentersProps) {
    return (<>
        <section className="comments">
            {stats.comments_data.comments.length===0
            && <p className="no-comments">There are no comments.</p>
            || stats.comments_data.comments.map(comment => (
                <AddComment key={comment.name} name={comment.name} date={comment.date} comment={comment.comment}/>
            ))}
        </section>
        <LeaveAComment/>
    </>)
}

export {
    type CreatorAndReviewStatus
}