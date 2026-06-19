import { useState, useCallback } from 'react'
import { api } from '@/services/api'

export function useAuth() {
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState(null)
    const [user, setUser] = useState(null)
    const [token, setToken] = useState(null)
    const [success, setSuccess] = useState(false)

    const resetState = useCallback(() => {
        setError(null)
        setSuccess(false)
    }, [])

    const register = useCallback(async (userData) => {
        setLoading(true);
        setError(null);
        setSuccess(false)

        try {
            if (!userData.name || userData.name.length < 3) {
                throw new Error('Nome deve ter pelo menos 3 caracteres')
            }

            if (!userData.email || !userData.email.includes('@')) {
                throw new Error('Email inválido')
            }

            if (!userData.password || userData.password.length < 6) {
                throw new Error('Senha deve ter pelo menos 6 caracteres')
            }

            const response = await api.createUser(userData)
            
            console.log('Usuário criado: ', response)

            if (response.user) { //! Possivel erro com a api
                setUser(response.user)
            }
            
            if (response.token) {
                setToken(response.token)
                localStorage.setItem('token', response.token)
                localStorage.setItem('user', JSON.stringify(response.user)) //! Possivel erro com a api
            }

            setSuccess(true)
            return { success: true, data: response }

        } catch (err) {
            console.error('Erro no registro: ', err)
            
            let errorMessage = err.message || 'Erro ao criar usuário'
            
            if (err.data?.error) {
                errorMessage = err.data.error
                if (err.data.error.includes('duplicate') || err.data.error.includes('already exists')) {
                    errorMessage = 'Este email já está cadastrado'
                }
            }

            setError(errorMessage)
            return { success: false, error: errorMessage }

        } finally {
            setLoading(false)
        }
    }, [])

    const login = useCallback(async (email, password) => {
        setLoading(true)
        setError(null)
        setSuccess(false)

        try {
            if (!email || !password) {
                throw new Error('Email e senha são obrigatórios')
            }

            const response = await api.login({ email, password })
            
            if (response.token) {
                setToken(response.token)
                setUser(response.user)
                localStorage.setItem('token', response.token)
                localStorage.setItem('user', JSON.stringify(response.user))
                setSuccess(true)
            }

            return { success: true, data: response }

        } catch (err) {
            setError(err.message || 'Erro ao fazer login')
            return { success: false, error: err.message }
        } finally {
            setLoading(false)
        }
    }, [])

    // Logout
    const logout = useCallback(() => {
        setUser(null)
        setToken(null)
        setSuccess(false)
        setError(null)
        localStorage.removeItem('token')
        localStorage.removeItem('user')
    }, [])

    // Verificar se usuário está autenticado
    const isAuthenticated = useCallback(() => {
        const storedToken = localStorage.getItem('token')
        return !!storedToken && !!token
    }, [token])

    return {
        register,
        login,
        logout,
        user,
        token,
        loading,
        error,
        success,
        resetState,
        isAuthenticated: isAuthenticated(),
    }
}