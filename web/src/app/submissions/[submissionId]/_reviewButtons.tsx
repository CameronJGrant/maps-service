import { Button, ButtonOwnProps } from "@mui/material";

type Review = "Completed" | "Submit" | "Reject" | "Revoke" | "Validate" | "Publish"
type Action = "completed" | "submit" | "reject" | "revoke" | "trigger-validate" | "trigger-publish"
interface ReviewButton {
	name: Review,
	action: Action,
	color: ButtonOwnProps["color"]
}

function ReviewButtonClicked(action: Action) {
	const post = fetch(`http://localhost:3000/v1/submissions/1/status/${action}`, {
		method: "POST",
		headers: {
			"Content-type": "application/json",
			"Cookie": "session_id=c5191ddc-eee1-4010-900c-6b2c7b6780ab"
		}
	})
}

function ReviewButton(props: ReviewButton) {
	return <Button color={props.color} variant="contained" onClick={() => { ReviewButtonClicked(props.action) }}>{props.name}</Button>
}

export default function ReviewButtons() {
	return (
		<section className="review-set">
			<ReviewButton color="error" name="Reject"    action="reject"/>
			<ReviewButton color="info"  name="Revoke"    action="revoke"/>
			<ReviewButton color="info"  name="Publish"   action="trigger-publish"/>
			<ReviewButton color="info"  name="Completed" action="completed"/>
			<ReviewButton color="info"  name="Submit"    action="submit"/>
			<ReviewButton color="info"  name="Validate"  action="trigger-validate"/>
		</section>
	)
}
