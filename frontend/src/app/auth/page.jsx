"use client"
const style = require('./auth.module.css')
import LoginOption from '@/components/layout/LoginOption/LoginOption'
import { LogIn } from 'lucide-react'
import { useState } from 'react'

export default function Login() {
    const [formData, setFormData] = useState({
        name: '',
        email: '',
        password: ''
    })

    const [loading, setLoading] = useState(false)
    const [resposta, setResposta] = useState(null)

    const handleChange = (e) => {
        const {name, value} = e.target
        setFormData(prev => ({
            ...prev,
            [name]: value
        }))
    }

    const handleSubmit = async (e) => {
        e.preventDefault()
        setLoading(true)

        try {
            const response = await fetch('http://localhost:8000/user/', {
                method: "POST",
                headers: {
                    'Key': process.env.TOKENAPI
                },
                body: JSON.stringify(formData),
            })

            const data = await response.json()

            if (response.ok) {
                setResposta({
                    sucesso: true,
                    data: data
                })
                setFormData({
                    name: '',
                    email: '',
                    password: ''
                })
            } else {
                setResposta({
                    sucesso: false,
                    erro: data.error || 'Erro ao enviar dados',
                    code: 400
                })
            }

        } catch (error) {
            setResposta({
                sucesso: false,
                error: error.message
            })
        } finally {
            setLoading(false)
        }
    }

    return (
        <>
            <form  onSubmit={handleSubmit} method="POST" className={style.create_count} >
                <h2>Cadastrese-se no Efficax Counter!</h2>
                <hr />

                <LoginOption name='name' 
                    type='text' 
                    label='Nome' 
                    maxText={50} 
                    placeH='Chikita123' 
                    formValue={formData.name} 
                />

                <LoginOption name='email' 
                    type='email' 
                    label='Email' 
                    maxText={50} 
                    placeH='example@counter.com'
                    formValue={formData.email}
                />
                <LoginOption name='password' 
                    type='password' 
                    label='Senha' 
                    maxText={120} 
                    placeH='********'
                    formValue={formData.password}
                />

                <button type='submit'>
                    <LogIn />
                    Cadastre-se
                </button>

                {resposta && 
                    <span>
                        {resposta}
                    </span>
                }
            </form>
        </>
    )
}