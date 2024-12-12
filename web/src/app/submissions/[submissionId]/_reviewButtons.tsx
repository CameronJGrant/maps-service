import { Button, ButtonOwnProps } from "@mui/material";

type Review = "Completed" | "Submit" | "Reject" | "Revoke" | "Validate" | "Publish"
type Action = "completed" | "submit" | "reject" | "revoke" | "trigger-validate" | "trigger-publish"
interface ReviewButton {
	name: Review,
	action: Action,
	color: ButtonOwnProps["color"]
}

// CREATE:
// curl -H 'Content-Type: application/json' \
// --cookie "session_id=123" \
// -d '{"ID":0,"DisplayName":"Example Map","Creator":"Quaternions","GameID":1,"Date":"2024-12-09T20:43:49.957660142-08:00","Submitter":52250025,"AssetID":255299419,"AssetVersion":7,"Completed":true,"TargetAssetID":5692134283,"StatusID":0}' \
// -X POST \
// localhost:8081/v1/submissions

// GET:
// curl -H 'Content-Type: application/json' \
// --cookie "session_id=c5191ddc-eee1-4010-900c-6b2c7b6780ab" \
// -X GET \
// localhost:8081/v1/submissions/1

// SUBMIT:
// curl -H 'Content-Type: application/json' \
// --cookie "session_id=c5191ddc-eee1-4010-900c-6b2c7b6780ab" \
// -X POST \
// localhost:8081/v1/submissions/1/status/submit

function ReviewButtonClicked(action: Action) {
	const post = fetch(`http://localhost:8081/v1/submissions/1/status/${action}`, {
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
