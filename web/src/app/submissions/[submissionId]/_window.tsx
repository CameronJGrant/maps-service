interface WindowStruct {
    className: string,
    title: string,
    children: React.ReactNode
}

export default function Window(window: WindowStruct) {
    return (
        <section className={window.className}>
            <header>
                <p>{window.title}</p>
            </header>
            <main>{window.children}</main>
        </section>
    )
}

export {
	type WindowStruct
}