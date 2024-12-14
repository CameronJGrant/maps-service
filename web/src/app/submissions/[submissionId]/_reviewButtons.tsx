import { Button, ButtonOwnProps } from "@mui/material";

type Review = "Completed" | "Submit" | "Reject" | "Revoke" | "Accept" | "Publish"
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
		}
	})
}

function ReviewButton(props: ReviewButton) {
	return <Button color={props.color} variant="contained" onClick={() => { ReviewButtonClicked(props.action) }}>{props.name}</Button>
}

export default function ReviewButtons() {
	return (
		<section className="review-set">
			<ReviewButton color="info"  name="Submit"    action="submit"/>
			<ReviewButton color="info"  name="Revoke"    action="revoke"/>
			<ReviewButton color="info"  name="Accept"    action="trigger-validate"/>
			<ReviewButton color="error" name="Reject"    action="reject"/>
			<ReviewButton color="info"  name="Publish"   action="trigger-publish"/>
			<ReviewButton color="info"  name="Completed" action="completed"/>
		</section>
	)
}
