import React, { SyntheticEvent, useState } from 'react'
import { useNavigate } from 'react-router-dom'

type Props = {
    onLoggedIn: () => Promise<void>
}

const Login: React.FC<Props> = ({onLoggedIn}) => {

    const [email, setEmail] = useState<string>('')
    const [password, setPassword] = useState<string>('')

    const navigate = useNavigate();

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault()

        const request = {
            email,
            password
        }

        const response = await fetch('/api/login', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(request)
        })

        if(response.ok){
            await onLoggedIn()
            navigate("/")
        }
    }

    return (
        <div>
            <h2>Login</h2>
            <form onSubmit={submit}>
                <input type="text" name="email" placeholder='Email' value={email} onChange={(e) => setEmail(e.target.value)} />
                <input type="password" name="password" placeholder='Password' value={password} onChange={(e) => setPassword(e.target.value)} />
                <button type="submit">Login</button>
            </form>            
        </div>
    )
}

export default Login
