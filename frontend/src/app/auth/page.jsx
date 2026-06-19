"use client"
const style = require('./auth.module.css')
import { LogIn } from 'lucide-react'

export default function Login() {

    return (
        <>
            <form method='POST' className={style.create_count} >
                <h2>Cadastrese-se no Efficax Counter!</h2>
                <hr />
            </form>
        </>
    )
}