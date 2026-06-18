const style = require(`./LoginButtom.module.css`)
import Link from "next/link"

export default function LoginButtom() {
    return (
        <Link href='/auth' className={style.login}>
            <span>Entrar</span>
        </Link>
    )
}