class ApiService {
    constructor() {
        this.baseURL = 'http://localhost:8000'
        this.timeout = 30000
    }

    async request(endpoint, options = {}) {
        const url = `${this.baseURL}${endpoint}`
        
        const config = {
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'Key': process.env.NEXT_PUBLIC_TOKENAPI,
                ...options.headers,
            },
            ...options,
        };

        const controller = new AbortController()
        const timeoutId = setTimeout(() => controller.abort(), this.timeout)

        try {
            console.log(`Enviando requisição para: ${url}`)
            console.log('Config:', config)

            const response = await fetch(url, {
                ...config,
                signal: controller.signal,
            })

            clearTimeout(timeoutId);

            console.log(`Status: ${response.status}`)

            let data
            const contentType = response.headers.get('content-type')
            
            if (contentType && contentType.includes('application/json')) {
                data = await response.json()
            } else {
                const text = await response.text()
                try {
                    data = JSON.parse(text)
                } catch {
                    data = { message: text }
                }
            }

            console.log('Dados:', data)


            if (!response.ok) {
                const error = new Error(data.error || data.message || `Erro ${response.status}`)
                error.status = response.status
                error.data = data
                throw error
            }

            return data

        } catch (error) {
            clearTimeout(timeoutId)

            // Tratamento de erros específicos
            if (error.name === 'AbortError') {
                throw new Error('A requisição demorou muito tempo. Verifique sua conexão.')
            }

            if (error.message === 'Failed to fetch') {
                throw new Error('Não foi possível conectar ao servidor. Verifique se a API está rodando.')
            }

            throw error
        }
    }

    async createUser(userData) {
        return this.request('/user', {
            method: 'POST',
            body: JSON.stringify(userData),
        })
    }

    //* Pega dados de usuario por id
    async getUser(id, token) { //! Falta implementacao na api
        return this.request(`/user/${id}`, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`,
            },
        })
    }
}

export const api = new ApiService()