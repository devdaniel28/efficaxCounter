const style = require('./Header.module.css')
import { Binary } from 'lucide-react'
import NavItem from '../NavItem/Navitem'
import Link from 'next/link'
import LoginButtom from '@/components/layout/LoginButtom/LoginButtom'

export default function Header() {
    return (
        <header className={style.header}>
                <Link href='/'>
                  <Binary/>
                </Link>
                  <h1>Efficax Counter</h1>
                    <nav>
                        <NavItem link='' name='Home' />
                        <NavItem link='config' name='Config' />
                        <LoginButtom />
                    </nav>
        </header>
    )
}