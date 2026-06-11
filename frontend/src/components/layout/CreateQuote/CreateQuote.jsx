const style = require('./CreateQuote.module.css')

import { Plus } from 'lucide-react'
import Link from 'next/link'

//todo Estilizar o componente
export default function CreateQuote() {
    return (
        <Link href='/config'>
            <Plus/>
                <h5>Create Quote</h5>
        </Link>
    )
}