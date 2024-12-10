import Link from "next/link"

import "./styles/header.scss"

interface HeaderButton {
	name: string,
	href: string
}
function HeaderButton(header: HeaderButton) {
	return (
		<Link href={header.href}>
			<button>{header.name}</button>
		</Link>
	)
}

export default function Header() {
	return (
		<header className="header-bar">
			<nav className="left">
				<HeaderButton name="Maps" href=""/>
			</nav>
			<nav className="right">
				<HeaderButton name="Home" href=""/>
				<HeaderButton name="Need" href=""/>
				<HeaderButton name="Menu" href=""/>
				<HeaderButton name="Items" href=""/>
			</nav>
		</header>
	)
}