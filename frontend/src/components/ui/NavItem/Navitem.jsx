const style = require('./NavItem.module.css')

import Link from 'next/link'

export default function NavItem({link, name}) {
    const linkStr = `/${link}`

    return (
        <Link href={linkStr} className={style.navitem}>
            <p>{name}</p>
        </Link>
    )
}