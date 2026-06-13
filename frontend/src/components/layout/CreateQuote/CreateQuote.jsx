const style = require('./CreateQuote.module.css')

import { Plus } from 'lucide-react'
import Link from 'next/link'

export default function CreateQuote() {
    return (
        <Link href='/config' className={style.createquote}>
            <Plus/>
                <h5>Create Quote</h5>
        </Link>
    )
}