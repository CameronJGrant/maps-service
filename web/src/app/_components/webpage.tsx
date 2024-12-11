import Header from "./header";

export default function Webpage({children}: Readonly<{children?: React.ReactNode}>) {
    return (<>
        <Header/>
        {children}
    </>)
}