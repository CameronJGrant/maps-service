import { Button, ButtonOwnProps } from "@mui/material";

type Review = "Completed" | "Submit" | "Reject" | "Revoke" | "Validate" | "Publish"
interface ReviewButton {
	name: Review,
	color: ButtonOwnProps["color"]
}

function ReviewButtonClicked(type: Review) {
	//magical api requesting goes here, i hope it wont blow up
	console.log(type)
}

function ReviewButton(props: ReviewButton) {
	return <Button color={props.color} variant="contained" onClick={() => { ReviewButtonClicked(props.name) }}>{props.name}</Button>
}

export default function ReviewButtons() {
	return (
		<section className="review-set">
			<ReviewButton color="error" name="Reject"/>
			<ReviewButton color="info"  name="Revoke"/>
			<ReviewButton color="info"  name="Publish"/>
			<ReviewButton color="info"  name="Completed"/>
			<ReviewButton color="info"  name="Submit"/>
			<ReviewButton color="info"  name="Validate"/>
		</section>
	)
}